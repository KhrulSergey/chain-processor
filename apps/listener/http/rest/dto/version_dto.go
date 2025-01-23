package dto

type VersionResponseDTO struct {
	AppVersion       string `json:"app_version"`
	CurrentTimestamp int64  `json:"current_timestamp"`
}
