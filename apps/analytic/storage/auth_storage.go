package storage

import (
	"github.com/khrulsergey/chain-processor/pkg/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthStorage interface {
	FindOpenIdClientByID(id uuid.UUID) (*model.OpenIDClient, error)
	StoreAuthorizationRequest(request *model.OpenIdAuthenticationRequest) error
	StoreAuthenticatedSession(session *model.AuthenticatedSession) error
	FindAuthenticatedSessionByAccessToken(token string) (*model.AuthenticatedSession, error)
	FindAuthorizationRequestByCode(code string) (*model.OpenIdAuthenticationRequest, error)
}

var (
	_ AuthStorage = &AuthStorageGorm{}
)

type AuthStorageGorm struct {
	db *gorm.DB
}

func NewAuthStorageGorm(db *gorm.DB) *AuthStorageGorm {
	return &AuthStorageGorm{
		db: db,
	}
}

func (s *AuthStorageGorm) FindOpenIdClientByID(id uuid.UUID) (*model.OpenIDClient, error) {
	client := &model.OpenIDClient{}
	err := s.db.
		Table(model.OpenIDClientTableName).
		Where("id = ?", id.String()).
		First(client).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *AuthStorageGorm) StoreAuthorizationRequest(request *model.OpenIdAuthenticationRequest) error {
	return s.db.Create(request).Error
}

func (s *AuthStorageGorm) StoreAuthenticatedSession(session *model.AuthenticatedSession) error {
	return s.db.Create(session).Error
}

func (s *AuthStorageGorm) FindAuthenticatedSessionByAccessToken(token string) (*model.AuthenticatedSession, error) {
	session := &model.AuthenticatedSession{}
	err := s.db.
		Table(model.AuthenticatedSessionTableName).
		Where("access_token = ?", token).
		First(session).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *AuthStorageGorm) FindAuthorizationRequestByCode(code string) (*model.OpenIdAuthenticationRequest, error) {
	request := &model.OpenIdAuthenticationRequest{}
	err := s.db.
		Table(model.OpenIDAuthenticationRequestTableName).
		Where("code = ?", code).
		First(request).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return request, nil
}
