package dto

type BaseDto interface {
	GetExternalId() string
	SetExternalId(externalId string)
}
