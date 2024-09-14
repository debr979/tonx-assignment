package migration

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
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
	password := `doug123`

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		utils.Logger.LogOutput(err)
	}

	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		utils.Logger.LogOutput(err)
	}

	if err := dbConn.Create(&models.Manager{
		ManagerName: "doug123", Password: fmt.Sprintf("%s", string(hash)),
	}).Error; err != nil {
		utils.Logger.LogOutput(err)
	}

	if err := dbConn.Create(&models.User{
		Username: "doug123", Password: fmt.Sprintf("%s", string(hash)),
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
