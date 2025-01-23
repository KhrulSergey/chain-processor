package service

import (
	storage2 "analytic/storage"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	AuthScopeOpenId  = "open_id"
	AuthScopeProfile = "profile"
	AuthScopeEmail   = "email"
	AuthScopePhone   = "phone"

	IDTokenIDPrefix     = "IT."
	AccessTokenIDPrefix = "AT."

	TokenType = "Bearer"
)

var (
	ErrInvalidValue       = errors.New("invalid value")
	ErrInvalidCredentials = errors.New("invalid credentials")

	ErrSessionNotFound = errors.New("session not found")
	ErrSessionInvalid  = errors.New("session is invalid")
	ErrSessionExpired  = errors.New("session expired")
)

var (
	knownScopes = map[string]struct{}{
		AuthScopeOpenId:  {},
		AuthScopeProfile: {},
		AuthScopeEmail:   {},
		AuthScopePhone:   {},
	}

	mandatoryScopes = map[string]struct{}{
		AuthScopeOpenId: {},
	}
)

type AuthService interface {
	GenerateAuthenticationCode(model.AuthenticationRequest) (string, error)
	GenerateTokenWithCodeAuth(model.CodeAuthCreds, model.SenderInfo) (*model.Token, error)
	GetSessionForAccessToken(accessToken string) (*model.AuthenticatedSession, error)
}

var (
	_ AuthService = &authService{}
)

type AuthServiceConfig struct {
	AuthCodeLifetime    time.Duration
	AccessTokenLifetime time.Duration
	IDTokenLifetime     time.Duration
	TokenIssuer         string
	SecretKey           []byte
}

type authService struct {
	authStorage storage2.AuthStorage

	userService UserService

	config AuthServiceConfig

	log *zap.SugaredLogger
}

func NewAuthService(authStorage storage2.AuthStorage, userService UserService, config AuthServiceConfig, log *zap.SugaredLogger) *authService {
	return &authService{
		authStorage: authStorage,
		userService: userService,
		log:         log,
		config:      config,
	}
}

func (s *authService) GenerateAuthenticationCode(authRequest model.AuthenticationRequest) (string, error) {
	if err := s.checkOpenIdClient(authRequest.ClientID); err != nil {
		return "", err
	}
	filteredScope, err := s.filterAuthScope(authRequest.Scope)
	if err != nil {
		return "", err
	}
	authRequest.Scope = filteredScope
	user, err := s.userService.GetAuthenticatedUser(authRequest.Username, authRequest.Password)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInvalidCredentials, err)
	}
	openIdAuthRequest, err := s.createAuthorizationRequest(authRequest, user)
	if err != nil {
		return "", err
	}
	return openIdAuthRequest.Code, nil
}

func (s *authService) checkOpenIdClient(cliendID uuid.UUID) error {
	_, err := s.authStorage.FindOpenIdClientByID(cliendID)
	if err == storage2.ErrNotFound {
		return fmt.Errorf("%w: unknown client", ErrInvalidCredentials)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) filterAuthScope(scopes []string) ([]string, error) {
	filteredScopes := make([]string, 0, len(scopes))
	visitedScopes := make(map[string]struct{})
	for _, scope := range scopes {
		if _, ok := knownScopes[scope]; !ok {
			return nil, fmt.Errorf("%w: unknown scope '%v'", ErrInvalidValue, scope)
		}
		if _, ok := visitedScopes[scope]; ok {
			continue
		}
		visitedScopes[scope] = struct{}{}
		filteredScopes = append(filteredScopes, scope)
	}
	for scope := range mandatoryScopes {
		if _, ok := visitedScopes[scope]; !ok {
			return nil, fmt.Errorf("%w: missing mandatory scope '%v'", ErrInvalidValue, scope)
		}
	}
	return filteredScopes, nil
}

func (s *authService) createAuthorizationRequest(authRequest model.AuthenticationRequest, user *model.User) (*model.OpenIdAuthenticationRequest, error) {
	openIdAuthRequest := &model.OpenIdAuthenticationRequest{
		ID:          uuid.New(),
		UserID:      user.ID,
		CliendID:    authRequest.ClientID,
		Scope:       strings.Join(authRequest.Scope, ","),
		RedirectURI: authRequest.RedirectURI,
		ExpiresAt:   time.Now().Add(s.config.AuthCodeLifetime),
	}
	openIdAuthRequest.Code = generateAuthenticationCode(openIdAuthRequest)
	if err := s.authStorage.StoreAuthorizationRequest(openIdAuthRequest); err != nil {
		return nil, err
	}
	return openIdAuthRequest, nil
}

func generateAuthenticationCode(openIdAuthRequest *model.OpenIdAuthenticationRequest) string {
	builder := strings.Builder{}
	builder.WriteString(openIdAuthRequest.ID.String())
	builder.WriteString(openIdAuthRequest.UserID.String())
	builder.WriteString(openIdAuthRequest.CliendID.String())
	builder.WriteString(openIdAuthRequest.Scope)
	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}

func (s *authService) GenerateTokenWithCodeAuth(creds model.CodeAuthCreds, sender model.SenderInfo) (*model.Token, error) {
	if _, err := s.authenticateOpenIdClient(creds.ClientID, creds.ClientSecret); err != nil {
		return nil, err
	}

	openIdAuthRequest, err := s.authStorage.FindAuthorizationRequestByCode(creds.Code)
	if err == storage2.ErrNotFound {
		return nil, fmt.Errorf("%w: code not found or expired", ErrInvalidCredentials)
	}
	if err != nil {
		return nil, err
	}

	if openIdAuthRequest.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("%w: code not found or expired", ErrInvalidCredentials)
	}
	if creds.RedirectURI != openIdAuthRequest.RedirectURI {
		return nil, fmt.Errorf("%w: do not match with auth request", ErrInvalidCredentials)
	}
	if creds.ClientID != openIdAuthRequest.CliendID {
		return nil, fmt.Errorf("%w: do not match with auth request", ErrInvalidCredentials)
	}

	user, err := s.userService.GetUserById(openIdAuthRequest.UserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidCredentials, err)
	}

	scope := strings.Split(openIdAuthRequest.Scope, ",")

	accessToken := s.createAccessToken(user, &creds, scope)
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken).SignedString(s.config.SecretKey)
	if err != nil {
		return nil, err
	}

	idToken := s.createIDToken(user, &creds, scope)
	idTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, idToken).SignedString(s.config.SecretKey)
	if err != nil {
		return nil, err
	}

	s.authStorage.StoreAuthenticatedSession(&model.AuthenticatedSession{
		ID:           uuid.New(),
		UserID:       user.ID,
		SessionID:    uuid.New(),
		Version:      1,
		AccessToken:  accessTokenString,
		RefreshToken: uuid.NewString(), // Replace with a real refresh token
		LoginIP:      sender.IP.String(),
		LoginDevice:  sender.Device,
		LoginDate:    time.Now(),
		ExpireAt:     accessToken.ExpiresAt.Time,
		IsValid:      true,
	})

	return &model.Token{
		AccessToken:  accessTokenString,
		RefreshToken: "",
		IDToken:      idTokenString,
		ExpiresIn:    s.config.AccessTokenLifetime,
		TokenType:    TokenType,
		Scope:        scope,
	}, nil
}

func (s *authService) authenticateOpenIdClient(id uuid.UUID, secret string) (*model.OpenIDClient, error) {
	openIdClient, err := s.authStorage.FindOpenIdClientByID(id)
	if err == storage2.ErrNotFound {
		return nil, fmt.Errorf("%w: unknown client", ErrInvalidCredentials)
	}
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(openIdClient.Secret), []byte(secret)); err != nil {
		return nil, fmt.Errorf("%w: client secret missmatch", ErrInvalidCredentials)
	}
	return openIdClient, nil
}

func (s *authService) createAccessToken(user *model.User, creds *model.CodeAuthCreds, scope []string) *model.AccessToken {
	now := time.Now()
	token := model.AccessToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        AccessTokenIDPrefix + uuid.NewString(),
			Subject:   user.ID.String(),
			Audience:  jwt.ClaimStrings{creds.ClientID.String()},
			Issuer:    s.config.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.config.AccessTokenLifetime)),
		},
		Scope: scope,
	}
	return &token
}

func (s *authService) createIDToken(user *model.User, creds *model.CodeAuthCreds, scope []string) *model.IDToken {
	now := time.Now()

	token := model.IDToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        IDTokenIDPrefix + uuid.NewString(),
			Subject:   user.ID.String(),
			Audience:  jwt.ClaimStrings{creds.ClientID.String()},
			Issuer:    s.config.TokenIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.config.IDTokenLifetime)),
		},
		Scope: scope,
	}

	scopeSet := toSet(scope)
	if _, ok := scopeSet[AuthScopeEmail]; ok {
		token.Email = &user.Email
	}
	if _, ok := scopeSet[AuthScopeProfile]; ok {
		token.FirstName = &user.FirstName
		token.SecondName = &user.LastName
	}
	if _, ok := scopeSet[AuthScopePhone]; ok {
		token.Phone = &user.Phone
	}

	return &token
}

func toSet(values []string) map[string]struct{} {
	valueSet := make(map[string]struct{}, len(values))
	for _, value := range values {
		valueSet[value] = struct{}{}
	}
	return valueSet
}

func (s *authService) GetSessionForAccessToken(accessToken string) (*model.AuthenticatedSession, error) {
	session, err := s.authStorage.FindAuthenticatedSessionByAccessToken(accessToken)
	if err == storage2.ErrNotFound {
		return nil, ErrSessionNotFound
	}
	if err != nil {
		return nil, err
	}
	if !session.IsValid {
		return nil, ErrSessionInvalid
	}
	if session.ExpireAt.Before(time.Now()) {
		return nil, ErrSessionExpired
	}
	return session, nil
}
