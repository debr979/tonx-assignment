package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"time"

	"os"

	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dns := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=verify-full", username, password, host, dbName)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to db fail: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("connect to db fail: %v", err)
	}

	sqlDb.SetConnMaxLifetime(300 * time.Second)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(100)

	return db
}

func Conn() *gorm.DB {
	return initDB()
}
