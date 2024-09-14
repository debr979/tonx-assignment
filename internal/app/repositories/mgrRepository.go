package repositories

import (
	"tonx-assignment/internal/app/db"
)

type mgrRepository struct{}

var MgrRepository mgrRepository

func (r *mgrRepository) Login(model any, mgrName string) error {
	dbConn := db.Conn()
	return dbConn.Where("manager_name = ?", mgrName).First(model).Error
}

func (r *mgrRepository) IsManager(model any, managerName string) error {

	dbConn := db.Conn()
	return dbConn.Where("manager_name = ?", managerName).First(model).Error
}
