package service

import (
	storage2 "analytic/storage"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/pkg/model"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	registrationMailMessageTemplate = "Your PWG ID has been created. To activate it, follow the link to set your login and password. \n\n%s"
)

var (
	ErrUserNotFound              = errors.New("user not found")
	ErrUserInactive              = errors.New("user is not active")
	ErrUserCredentialsInvalid    = errors.New("user credentials invalid")
	ErrInvalidUserActivationCode = errors.New("invalid activation code")
	ErrUserLoginAlreadyTaken     = errors.New("user login is taken already")
)

type UserService interface {
	GetUserById(id uuid.UUID) (*model.User, error)
	GetAuthenticatedUser(login, password string) (*model.User, error)
	RegisterUserIfAbsent(personalData model.UserPersonalData, contactInfo model.UserContactInfo) (*model.User, error)
	ActivateUser(creds model.UserActivationCreds, sender model.SenderInfo) error
}

var (
	_ UserService = &userService{}
)

type UserServiceConfig struct {
	ServiceExternalURL        string
	ServiceProtocol           string
	AfterRegistrationEndpoint string
	MailSenderName            string
	MailSenderAddress         string
}

type userService struct {
	mailService MailService

	userStorage storage2.UserStorage

	config UserServiceConfig

	log *zap.SugaredLogger
}

func NewUserService(mailService MailService, userStorage storage2.UserStorage, config UserServiceConfig, log *zap.SugaredLogger) *userService {
	return &userService{
		mailService: mailService,
		userStorage: userStorage,
		config:      config,
		log:         log,
	}
}

func (s *userService) GetUserById(id uuid.UUID) (*model.User, error) {
	user, err := s.userStorage.FindUserByID(id)
	if err == storage2.ErrNotFound {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, ErrUserInactive
	}
	return user, nil
}

func (s *userService) GetAuthenticatedUser(login, password string) (*model.User, error) {
	user, err := s.userStorage.FindUserByLogin(login)
	if err == storage2.ErrNotFound {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, ErrUserInactive
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrUserCredentialsInvalid
	}
	return user, nil
}

func (s *userService) RegisterUserIfAbsent(personalData model.UserPersonalData, contactInfo model.UserContactInfo) (*model.User, error) {
	var (
		user *model.User
		err  error
	)

	user, err = s.userStorage.FindUserByPersonalData(personalData)
	if err == nil {
		return user, nil
	}
	if err != storage2.ErrNotFound {
		return nil, err
	}

	user = &model.User{
		ID:               uuid.New(),
		Login:            contactInfo.Email,
		Password:         "*",
		UserPersonalData: personalData,
		UserContactInfo:  contactInfo,
		IsActive:         false,
	}
	activationCode := createActivationCode(user)
	user.ActivationCode = &activationCode

	err = s.userStorage.CreateUser(user)
	if err != nil {
		return nil, err
	}

	message := s.createRegistartionMailMessage(user)
	err = s.mailService.SendMail(message)
	if err != nil {
		s.log.Errorf("registration mail send failed: %v", err)
	}

	return user, nil
}

func (s *userService) createRegistartionMailMessage(user *model.User) model.MailMessage {
	registrationLink := &url.URL{
		Scheme:   s.config.ServiceProtocol,
		Host:     s.config.ServiceExternalURL,
		Path:     s.config.AfterRegistrationEndpoint,
		RawQuery: fmt.Sprintf("activation_code=%s", *user.ActivationCode),
	}
	return model.MailMessage{
		From: mail.Address{
			Name:    user.FirstName + " " + user.LastName,
			Address: user.Email,
		},
		To: mail.Address{
			Name:    s.config.MailSenderName,
			Address: s.config.MailSenderAddress,
		},
		Subject: "PWG registration",
		Content: fmt.Sprintf(registrationMailMessageTemplate, registrationLink.String()),
	}
}

func createActivationCode(user *model.User) string {
	builder := strings.Builder{}
	builder.WriteString(user.ID.String())
	builder.WriteString(user.Login)
	builder.WriteString(user.Password)
	builder.WriteString(user.FirstName)
	builder.WriteString(user.LastName)
	builder.WriteString(user.Gender)
	builder.WriteString(user.DateOfBirth.String())
	builder.WriteString(user.Country)
	builder.WriteString(user.City)
	builder.WriteString(user.Address)
	builder.WriteString(user.PostalCode)
	builder.WriteString(user.Email)
	builder.WriteString(user.Phone)
	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}

func (s *userService) ActivateUser(creds model.UserActivationCreds, sender model.SenderInfo) error {
	user, err := s.userStorage.FindUserByActivationCode(creds.ActivationCode)
	if err == storage2.ErrNotFound {
		return ErrInvalidUserActivationCode
	}
	if err != nil {
		return err
	}

	user.Login = creds.Login
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.IsActive = true
	user.ActivationCode = nil

	user.RegistrationDevice = &sender.Device
	registrationIP := sender.IP.String()
	user.RegistrationIP = &registrationIP
	registrationDate := time.Now()
	user.RegistrationDate = &registrationDate

	err = s.userStorage.UpdateUser(user)
	if err == storage2.ErrNotFound {
		return ErrInvalidUserActivationCode
	}
	if err == storage2.ErrUniqueLoginConstraintViolation {
		return ErrUserLoginAlreadyTaken
	}
	return err
}
