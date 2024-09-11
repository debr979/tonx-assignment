package models

type RefreshTokenRequest struct {
	Username     string `json:"username" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
