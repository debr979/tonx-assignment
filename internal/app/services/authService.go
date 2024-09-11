package services

import (
	"errors"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/pkg/utils"
)

type authService struct{}

var AuthService authService

func (r *authService) RefreshToken(req models.RefreshTokenRequest) (any, error) {
	isValid, username, err := utils.JsonWebToken.VerifyJWToken(req.RefreshToken, "refresh_token")
	if !isValid || username == "" || err != nil {
		return nil, err
	}

	if req.Username != username {
		return nil, errors.New("invalid refresh token")
	}

	return utils.JsonWebToken.GenerateJWToken(username)
}
