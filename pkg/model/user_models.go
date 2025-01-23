package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	UserTableName = "users"
)

type UserPersonalData struct {
	FirstName   string    `gorm:"column:first_name"`
	LastName    string    `gorm:"column:last_name"`
	Gender      string    `gorm:"column:gender"`
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	Country     string    `gorm:"column:country"`
	City        string    `gorm:"column:city"`
	Address     string    `gorm:"column:address"`
	PostalCode  string    `gorm:"column:postal_code"`
}

type UserContactInfo struct {
	Email string `gorm:"column:email"`
	Phone string `gorm:"column:phone"`
}

type User struct {
	UserPersonalData
	UserContactInfo

	ID uuid.UUID `gorm:"column:id"`

	Login    string `gorm:"column:login"`
	Password string `gorm:"column:password"`

	RegistrationDate   *time.Time `gorm:"column:registration_date"`
	RegistrationIP     *string    `gorm:"column:registration_ip"`
	RegistrationDevice *string    `gorm:"column:registration_device"`

	IsActive       bool    `gorm:"column:is_active"`
	ActivationCode *string `gorm:"column:activation_code"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime:true"`
}

func (User) TableName() string {
	return UserTableName
}

type UserActivationCreds struct {
	Login          string
	Password       string
	ActivationCode string
}
