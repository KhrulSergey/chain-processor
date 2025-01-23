package model

import (
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/google/uuid"
)

const (
	OpenIDClientTableName                = "open_id_clients"
	OpenIDAuthenticationRequestTableName = "open_id_authentication_requests"
	AuthenticatedSessionTableName        = "sessions"
)

type OpenIDClient struct {
	ID     uuid.UUID `gorm:"column:id"`
	Secret string    `gorm:"column:secret"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
}

func (OpenIDClient) TableName() string {
	return OpenIDClientTableName
}

type OpenIdAuthenticationRequest struct {
	ID          uuid.UUID `gorm:"column:id"`
	UserID      uuid.UUID `gorm:"column:user_id"`
	CliendID    uuid.UUID `gorm:"column:client_id"`
	Code        string    `gorm:"column:code"`
	Scope       string    `gorm:"column:scope"`
	RedirectURI string    `gorm:"column:redirect_uri"`
	ExpiresAt   time.Time `gorm:"column:expires_at"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime:true"`
}

func (OpenIdAuthenticationRequest) TableName() string {
	return OpenIDAuthenticationRequestTableName
}

type AuthenticatedSession struct {
	ID     uuid.UUID `gorm:"column:id"`
	UserID uuid.UUID `gorm:"column:user_id"`

	SessionID uuid.UUID `gorm:"column:session_id"`
	Version   int       `gorm:"column:version"`

	AccessToken  string `gorm:"column:access_token"`
	RefreshToken string `gorm:"column:refresh_token"`

	LoginDate   time.Time `gorm:"column:login_date"`
	LoginIP     string    `gorm:"column:login_ip"`
	LoginDevice string    `gorm:"column:login_device"`

	ExpireAt time.Time `gorm:"column:expires_at"`
	IsValid  bool      `gorm:"column:is_valid"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
}

func (AuthenticatedSession) TableName() string {
	return AuthenticatedSessionTableName
}

type AuthenticationRequest struct {
	Username    string
	Password    string
	ClientID    uuid.UUID
	RedirectURI string
	Scope       []string
}

type CodeAuthCreds struct {
	ClientID     uuid.UUID
	ClientSecret string
	RedirectURI  string
	Code         string
}

type AccessToken struct {
	jwt.RegisteredClaims
	Scope []string `json:"scope"`
}

type IDToken struct {
	jwt.RegisteredClaims
	FirstName  *string  `json:"first_name,omitempty"`
	SecondName *string  `json:"second_name,omitempty"`
	Email      *string  `json:"email,omitempty"`
	Phone      *string  `json:"phone,omitempty"`
	Scope      []string `json:"scope"`
}

type Token struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
	TokenType    string
	ExpiresIn    time.Duration
	Scope        []string
}
