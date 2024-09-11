package models

type JWToken struct {
	AccessToken  TokenValue `json:"access_token"`
	RefreshToken TokenValue `json:"refresh_token"`
}

type TokenValue struct {
	Token     string `json:"token"`
	ExpiredAt int64  `json:"expiredAt"`
}
