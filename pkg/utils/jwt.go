package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"tonx-assignment/internal/app/models"
)

type jsonWebToken struct{}

var JsonWebToken jsonWebToken

func (r *jsonWebToken) GenerateJWToken(username string) (*models.JWToken, error) {

	if username == "" {
		return nil, errors.New(`token generate failed`)
	}

	accessExpiredAt := time.Now().Add(30 * time.Minute)
	refreshExpiredAt := time.Now().Add(8 * time.Hour)

	accessClaim := &jwt.RegisteredClaims{
		Issuer:    "tonx_assignment",
		Subject:   "tonx_assignment",
		ExpiresAt: jwt.NewNumericDate(accessExpiredAt),
		ID:        username,
	}

	refreshClaim := &jwt.RegisteredClaims{
		Issuer:    "tonx_assignment",
		Subject:   "tonx_assignment",
		ExpiresAt: jwt.NewNumericDate(refreshExpiredAt),
		ID:        username,
	}

	secretKey := []byte(fmt.Sprintf(`@!access_token!@`))
	refreshKey := []byte(fmt.Sprintf(`@!refresh_token!@`))
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaim).SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaim).SignedString(refreshKey)
	if err != nil {
		return nil, err
	}

	return &models.JWToken{
		AccessToken: models.TokenValue{
			Token:     accessToken,
			ExpiredAt: accessExpiredAt.Unix(),
		},
		RefreshToken: models.TokenValue{
			Token:     refreshToken,
			ExpiredAt: refreshExpiredAt.Unix(),
		},
	}, nil
}

func (r *jsonWebToken) VerifyJWToken(jwtToken, verifyType string) (bool, string, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(fmt.Sprintf(`@!%s!@`, verifyType)), nil
	})
	if err != nil {
		return false, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token.Valid, claims["jti"].(string), nil
	}

	return token.Valid, "", nil
}
