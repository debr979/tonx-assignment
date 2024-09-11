package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/services"
)

type userController struct {
	baseControllers
}

var UserController userController

// Register
// @Title Register
// @Tags User
// @Description Register member
// @Summary Register member
// @Param body body models.RegisterRequest true "register account"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /users/user [post]
func (r *userController) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if len(req.Username) > 50 || len(req.Password) > 50 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Username) == "" {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	userId, err := services.UserService.Register(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	result := map[string]any{
		"user_id": userId,
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), result)
	return
}

// ChangePassword
// @Title ChangePassword
// @Tags User
// @Description Change member password
// @Summary Change member password
// @Param body body models.ChangePasswordRequest true "modify member password"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /users/user [patch]
func (r *userController) ChangePassword(c *gin.Context) {
	var req models.ChangePasswordRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if req.Password == req.NewPassword {
		r.Response(http.StatusBadRequest, "same password", nil)
		return
	}

	username, isExist := c.Get(`username`)
	if !isExist {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	if req.Username != username.(string) {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	if err := services.UserService.ChangePassword(req); err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), true)

}

// DeleteAccount
// @Title DeleteAccount
// @Tags User
// @Description DeleteAccount member
// @Summary DeleteAccount member
// @Param body body models.DeleteAccountRequest true "delete account"
// @Success 200 {bool} bool ""
// @Failure 400 {object} string "api error"
// @Router /users/user [delete]
func (r *userController) DeleteAccount(c *gin.Context) {
	var req models.DeleteAccountRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	username, isExist := c.Get(`username`)
	if !isExist {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	if req.Username != username.(string) {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	if err := services.UserService.DeleteAccount(req); err != nil {
		r.Response(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), true)
}

// Login
// @Title Login
// @Tags User
// @Description Login member
// @Summary Login member
// @Param body body models.LoginRequest true "login account"
// @Success 200 {object} models.JWToken ""
// @Failure 400 {object} string "api error"
// @Router /user/login [post]
func (r *userController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if len(req.Username) > 100 || len(req.Password) > 100 {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	jwt, err := services.UserService.Login(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), jwt)

}
