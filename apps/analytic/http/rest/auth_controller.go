package rest

import (
	"analytic/http/rest/dto"
	"analytic/service"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/pkg/http/httppath"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"github.com/khrulsergey/chain-processor/util"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/url"
	"strings"
)

const (
	oidcResponseType = "response_type"
	oidcClientId     = "client_id"
	oidcClientSecret = "client_secret"
	oidcRedirectURI  = "redirect_uri"
	oidcScope        = "scope"
	oidcCode         = "code"
	oidcGrantType    = "grant_type"

	formUsernameFieldName = "username"
	formPasswordFieldName = "password"
)

var (
	supportedResponeTypes = map[string]struct{}{
		"code": {},
	}

	supportedGrantTypes = map[string]struct{}{
		"authorization_code": {},
	}
)

type AuthController struct {
	authService service.AuthService

	log *zap.SugaredLogger
}

func NewAuthController(authService service.AuthService, log *zap.SugaredLogger) *AuthController {
	return &AuthController{
		authService: authService,
		log:         log,
	}
}

// RedirectToLogin godoc
//
//	@Summary		Redirect to login
//	@Description	redirects user to login page
//	@Tags			auth
//
//	@Param			response_type	query	string	true	"authentication response type"						Enums(code)
//	@Param			client_id		query	string	true	"id of the client (application)"					Format(uuid)
//	@Param			redirect_uri	query	string	true	"url to redirect after successful authentication"	Format(uri)
//	@Param			scope			query	string	true	"a list of scopes (open_id,profile,phone,email) for ID token separated by ' '"
//
//	@Success		302				"redirects to login page"
//	@Failure		400
//	@Failure		422
//	@Failure		500
//
//	@Router			/authorize [get]
func (c *AuthController) RedirectToLogin(ctx *fasthttp.RequestCtx) {
	_, err := util.GetRequiredQueryArgs(ctx, oidcResponseType, oidcClientId, oidcRedirectURI, oidcScope)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	loginURL, err := url.ParseRequestURI(string(ctx.Request.URI().FullURI()))
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	loginURL.Path = httppath.AuthGroup
	ctx.Redirect(loginURL.String(), fasthttp.StatusFound)
}

// Authenticate godoc
//
//	@Summary		Authenticates user
//	@Description	authenticates user using OpenId Listen
//	@Tags			auth
//
//	@Accept			x-www-form-urlencoded
//
//	@Param			response_type	query		string	true	"authentication response type"						Enums(code)
//	@Param			client_id		query		string	true	"id of the client (application)"					Format(uuid)
//	@Param			redirect_uri	query		string	true	"url to redirect after successful authentication"	Format(uri)
//	@Param			scope			query		string	true	"a list of scopes (open_id,profile,phone,email) for ID token separated by ' '"
//
//	@Param			username		formData	string	true	"user login"	Format(email)
//	@Param			password		formData	string	true	"user password"	Format(password)
//
//	@Success		302				"redirects to 'redirect_uri' with authentication credentials in query based on response type"
//	@Failure		400
//	@Failure		422
//	@Failure		500
//
//	@Router			/authorize [post]
func (c *AuthController) Authenticate(ctx *fasthttp.RequestCtx) {
	requiredArgs, err := util.GetRequiredQueryArgs(ctx, oidcResponseType, oidcClientId, oidcRedirectURI, oidcScope)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}
	requiredFormFields, err := util.GetRequiredFormFields(ctx, formUsernameFieldName, formPasswordFieldName)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if _, ok := supportedResponeTypes[requiredArgs[oidcResponseType]]; !ok {
		ctx.Error("unsupported response type", fasthttp.StatusBadRequest)
		return
	}

	parsedRedirectURI, err := url.ParseRequestURI(requiredArgs[oidcRedirectURI])
	if err != nil {
		ctx.Error("invalid redirect uri", fasthttp.StatusBadRequest)
		return
	}
	parsedClientId, err := uuid.Parse(requiredArgs[oidcClientId])
	if err != nil {
		ctx.Error("invalid client id format", fasthttp.StatusBadRequest)
	}

	authCode, err := c.authService.GenerateAuthenticationCode(model.AuthenticationRequest{
		ClientID:    parsedClientId,
		Username:    requiredFormFields[formUsernameFieldName],
		Password:    requiredFormFields[formPasswordFieldName],
		Scope:       strings.Split(strings.TrimSpace(requiredArgs[oidcScope]), " "),
		RedirectURI: requiredArgs[oidcRedirectURI],
	})
	if err != nil {
		c.log.Infof("Fail to generate authenification code for client '%v' with: %v", parsedClientId, err)
		if errors.Is(err, service.ErrInvalidCredentials) {
			ctx.Error(err.Error(), fasthttp.StatusUnauthorized)
		} else if errors.Is(err, service.ErrInvalidValue) {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			c.log.Error(err)
			ctx.Error("something went wrong", fasthttp.StatusInternalServerError)
		}
		return
	}

	parsedRedirectURI.RawQuery = appendQuery(
		parsedRedirectURI.RawQuery,
		oidcCode, url.QueryEscape(authCode),
	)

	ctx.Redirect(parsedRedirectURI.String(), fasthttp.StatusFound)
}

func appendQuery(baseQuery, query, value string) string {
	builder := strings.Builder{}
	if baseQuery != "" {
		builder.WriteString(baseQuery)
		builder.WriteRune('&')
	}
	builder.WriteString(query)
	builder.WriteRune('=')
	builder.WriteString(value)
	return builder.String()
}

// GenerateToken godoc
//
//	@Summary		Token generation
//	@Description	generates token for user
//	@Tags			auth
//
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//
//	@Param			grant_type		formData	string	true	"authentication type"					Enums(authorization_code)
//	@Param			client_id		formData	string	true	"id of the client (application)"		Format(uuid)
//	@Param			client_secret	formData	string	true	"secret of the client (application)"	Format(password)
//	@Param			code			formData	string	true	"authentication code"
//	@Param			redirect_uri	formData	string	true	"redirect uri which was used for authentication"	Format(uri)
//
//	@Success		200				{object}	dto.TokenResponseDTO
//	@Failure		400
//	@Failure		422
//	@Failure		500
//
//	@Router			/token [post]
func (c *AuthController) GenerateToken(ctx *fasthttp.RequestCtx) {
	requiredArgs, err := util.GetRequiredFormFields(ctx, oidcGrantType, oidcClientId, oidcClientSecret, oidcCode, oidcRedirectURI)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if _, ok := supportedGrantTypes[requiredArgs[oidcGrantType]]; !ok {
		ctx.Error("unsupported grant type", fasthttp.StatusBadRequest)
		return
	}

	clientUUID, err := uuid.Parse(requiredArgs[oidcClientId])
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
		return
	}

	token, err := c.authService.GenerateTokenWithCodeAuth(
		model.CodeAuthCreds{
			ClientID:     clientUUID,
			ClientSecret: requiredArgs[oidcClientSecret],
			RedirectURI:  requiredArgs[oidcRedirectURI],
			Code:         requiredArgs[oidcCode],
		},
		util.GetSenderInfo(ctx),
	)
	if err != nil {
		c.log.Infof("Failed to generate token for with: %v", err)
		if errors.Is(err, service.ErrInvalidCredentials) {
			ctx.Error(err.Error(), fasthttp.StatusUnauthorized)
		} else if errors.Is(err, service.ErrInvalidValue) {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			c.log.Error(err)
			ctx.Error("something went wrong", fasthttp.StatusInternalServerError)
		}
		return
	}

	tokenDto := &dto.TokenResponseDTO{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		IDToken:      token.IDToken,
		TokenType:    token.TokenType,
		Scope:        strings.Join(token.Scope, " "),
		ExpiresInSec: int(token.ExpiresIn.Seconds()),
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.SetContentType("application/json")
	if err != json.NewEncoder(ctx.Response.BodyWriter()).Encode(tokenDto) {
		c.log.Error(err)
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	}
}
