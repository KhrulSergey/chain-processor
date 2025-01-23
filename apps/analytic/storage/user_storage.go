package storage

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/khrulsergey/chain-processor/pkg/model"

	"gorm.io/gorm"
)

const (
	postgresUniqueViolationCode = "23505"

	dbLoginUniqueConstraintName = "users_login_key"
)

var (
	ErrUniqueLoginConstraintViolation = errors.New("")
)

type UserStorage interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	FindUserByID(id uuid.UUID) (*model.User, error)
	FindUserByLogin(login string) (*model.User, error)
	FindUserByPersonalData(personalData model.UserPersonalData) (*model.User, error)
	FindUserByActivationCode(activationCode string) (*model.User, error)
}

var (
	_ UserStorage = &UserStorageGorm{}
)

type UserStorageGorm struct {
	db *gorm.DB
}

func NewUserStorageGorm(db *gorm.DB) *UserStorageGorm {
	return &UserStorageGorm{
		db: db,
	}
}

func (s *UserStorageGorm) CreateUser(user *model.User) error {
	return s.db.Create(user).Error
}

func (s *UserStorageGorm) UpdateUser(user *model.User) error {
	err := s.db.Save(user).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	if pgErr, ok := err.(*pgconn.PgError); ok {
		if pgErr.Code == postgresUniqueViolationCode && pgErr.ConstraintName == dbLoginUniqueConstraintName {
			return ErrUniqueLoginConstraintViolation
		}
	}
	return err
}

func (s *UserStorageGorm) FindUserByLogin(login string) (*model.User, error) {
	user := &model.User{}
	err := s.db.
		Table(model.UserTableName).
		Where("login = ?", login).
		First(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStorageGorm) FindUserByID(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	err := s.db.
		Table(model.UserTableName).
		Where("id = ?", id).
		First(user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStorageGorm) FindUserByPersonalData(personalData model.UserPersonalData) (*model.User, error) {
	user := &model.User{}
	err := s.db.
		Table(model.UserTableName).
		Where("first_name = ?", personalData.FirstName).
		Where("last_name = ?", personalData.LastName).
		Where("gender = ?", personalData.Gender).
		Where("date_of_birth = ?", personalData.DateOfBirth).
		Where("country = ?", personalData.Country).
		Where("city = ?", personalData.City).
		Where("address = ?", personalData.Address).
		Where("postal_code = ?", personalData.PostalCode).
		First(&user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStorageGorm) FindUserByActivationCode(activationCode string) (*model.User, error) {
	user := &model.User{}
	err := s.db.
		Table(model.UserTableName).
		Where("activation_code = ?", activationCode).
		First(&user).
		Error
	if err == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}
