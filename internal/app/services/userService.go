package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
	"tonx-assignment/pkg/utils"
)

type userService struct{}

var UserService userService

func (r *userService) Register(req models.RegisterRequest) (int64, error) {
	var user models.User

	if err := repositories.UserRepository.IsUser(&user, req.Username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}

	if user.Id > 0 {
		return 0, errors.New("its registered")
	}

	h := sha256.New()
	_, _ = io.WriteString(h, req.Password)
	user.Username = req.Username
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	if err := repositories.UserRepository.Register(&user); err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (r *userService) ChangePassword(req models.ChangePasswordRequest) error {
	hashOldPassword := sha256.New()
	_, _ = io.WriteString(hashOldPassword, req.Password)
	req.Password = fmt.Sprintf("%x", hashOldPassword.Sum(nil))
	hashNewPassword := sha256.New()
	_, _ = io.WriteString(hashOldPassword, req.NewPassword)
	req.NewPassword = fmt.Sprintf("%x", hashNewPassword.Sum(nil))

	if err := repositories.UserRepository.ChangePassword(req.Username, req.Password, req.NewPassword); err != nil {
		return err
	}

	return nil
}

func (r *userService) DeleteAccount(req models.DeleteAccountRequest) error {
	hashPassword := sha256.New()
	_, _ = io.WriteString(hashPassword, req.Password)
	req.Password = fmt.Sprintf("%x", hashPassword.Sum(nil))

	if err := repositories.UserRepository.DeleteAccount(req.Username, req.Username); err != nil {
		return err
	}
	return nil
}

func (r *userService) Login(req models.LoginRequest) (any, error) {
	var user models.User

	h := sha256.New()
	_, _ = io.WriteString(h, req.Password)
	req.Password = fmt.Sprintf("%x", h.Sum(nil))

	if err := repositories.UserRepository.Login(&user, req.Username, req.Password); err != nil {
		return nil, err
	}

	if user.Id == 0 {
		return nil, errors.New("not member")
	}

	return utils.JsonWebToken.GenerateJWToken(user.Username)
}

func (r *userService) GetUserId(username string) (int64, error) {
	var user models.User

	if err := repositories.UserRepository.IsUser(&user, username); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
	}

	if user.Id == 0 {
		return 0, errors.New("not member")
	}

	return user.Id, nil
}
