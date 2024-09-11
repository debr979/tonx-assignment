package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/services"
)

type mgrController struct {
	baseControllers
}

var MgrController mgrController

// Login
// @Title Login
// @Tags Manager
// @Description Login member
// @Summary Login member
// @Param body body models.MgrLoginRequest true "login account"
// @Success 200 {object} models.JWToken ""
// @Failure 400 {object} string "api error"
// @Router /mgr/login [post]
func (r *mgrController) Login(c *gin.Context) {
	var req models.MgrLoginRequest
	if err := r.Request(c).ParseBody(&req); err != nil {
		r.Response(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
		return
	}

	if strings.TrimSpace(req.ManagerName) == "" || strings.TrimSpace(req.Password) == "" {
		r.Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	jwt, err := services.MgrService.Login(req)
	if err != nil {
		r.Response(http.StatusInternalServerError, err.Error(), nil)
		return
	}

	r.Response(http.StatusOK, http.StatusText(http.StatusOK), jwt)

}
