package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
	"tonx-assignment/pkg/utils"
)

type mgrService struct{}

var MgrService mgrService

func (r *mgrService) Login(req models.MgrLoginRequest) (any, error) {
	var mgr models.Manager
	h := sha256.New()
	_, _ = io.WriteString(h, req.Password)
	req.Password = fmt.Sprintf("%x", h.Sum(nil))

	if err := repositories.MgrRepository.Login(&mgr, req.ManagerName, req.Password); err != nil {
		return nil, err
	}

	if mgr.Id == 0 {
		return nil, errors.New("not manager")
	}

	return utils.JsonWebToken.GenerateJWToken(mgr.ManagerName)
}
