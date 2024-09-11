package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/services"
)

type authController struct {
	baseControllers
}

var AuthController authController

// RefreshToken
// @Title RefreshToken
// @Tags Auth
// @Description refresh access token
// @Summary refresh access token
// @Param body models.RefreshTokenRequest true "refresh access token"
// @Success 200 {object} models.JWToken ""
// @Failure 400 {object} string "api error"
// @Router /auth/refreshToken [post]
func (r *authController) RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	jwt, err := services.AuthService.RefreshToken(req)
	if err != nil {
		r.Response(http.StatusUnauthorized, err.Error(), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), jwt)
	return
}
