package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
	"tonx-assignment/pkg/utils"
)

type mgrService struct{}

var MgrService mgrService

func (r *mgrService) Login(req models.MgrLoginRequest) (any, error) {
	var mgr models.Manager

	if err := repositories.MgrRepository.Login(&mgr, req.ManagerName); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(mgr.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	if mgr.Id == 0 {
		return nil, errors.New("not manager")
	}

	return utils.JsonWebToken.GenerateJWToken(mgr.ManagerName)
}
