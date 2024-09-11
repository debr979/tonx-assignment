package migration

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"tonx-assignment/internal/app/db"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/pkg/utils"
)

func Migrate() {
	log.Println("Database tables migrating")
	dbConn := db.Conn()

	if err := dbConn.AutoMigrate(&models.User{}, &models.Coupon{}, &models.UserCoupon{}, &models.Manager{}); err != nil {
		utils.Logger.LogOutput(err)
	}

	hashMgrPassword := sha256.New()
	_, _ = io.WriteString(hashMgrPassword, "doug123")

	if err := dbConn.Create(&models.Manager{
		ManagerName: "doug123", Password: fmt.Sprintf("%x", hashMgrPassword.Sum(nil)),
	}).Error; err != nil {
		utils.Logger.LogOutput(err)
	}

	if err := dbConn.Create(&models.User{
		Username: "doug123", Password: fmt.Sprintf("%x", hashMgrPassword.Sum(nil)),
	}).Error; err != nil {
		utils.Logger.LogOutput(err)
	}
}

func Clear() {
	dbConn := db.Conn()
	if err := dbConn.Migrator().DropTable(&models.User{}, &models.Coupon{}, &models.UserCoupon{}, &models.Manager{}); err != nil {
		utils.Logger.LogOutput(err)
	}
}
