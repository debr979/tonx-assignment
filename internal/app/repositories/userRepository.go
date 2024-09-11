package repositories

import (
	"tonx-assignment/internal/app/db"
	"tonx-assignment/internal/app/models"
)

type userRepository struct {
}

var UserRepository userRepository

func (r *userRepository) Register(model any) error {
	dbConn := db.Conn()
	return dbConn.Create(model).Error
}

func (r *userRepository) ChangePassword(username, password, newPassword string) error {
	dbConn := db.Conn()
	return dbConn.Where("username = ? and password = ?", username, password).Model(&models.User{}).Update("password", newPassword).Error
}

func (r *userRepository) DeleteAccount(username, password string) error {
	dbConn := db.Conn()
	return dbConn.Where("username = ? and password = ?", username, password).Model(&models.User{}).Update("is_deleted", false).Error
}

func (r *userRepository) Login(model any, username, password string) error {
	dbConn := db.Conn()
	return dbConn.Where("username = ? and password = ?", username, password).First(model).Error
}

func (r *userRepository) IsUser(model any, username string) error {
	dbConn := db.Conn()
	return dbConn.Where("username = ?", username).First(model).Error
}
