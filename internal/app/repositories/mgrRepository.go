package repositories

import (
	"tonx-assignment/internal/app/db"
)

type mgrRepository struct{}

var MgrRepository mgrRepository

func (r *mgrRepository) Login(model any, mgrName, password string) error {
	dbConn := db.Conn()
	return dbConn.Where("manager_name = ? and password = ?", mgrName, password).First(model).Error
}

func (r *mgrRepository) IsManager(model any, managerName string) error {

	dbConn := db.Conn()
	return dbConn.Where("manager_name = ?", managerName).First(model).Error
}
