package dto

type TokenResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresInSec int    `json:"expires_in"`
	Scope        string `json:"scope"`
}
