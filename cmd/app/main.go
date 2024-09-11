package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"tonx-assignment/internal/app/cronJobs"
	"tonx-assignment/internal/app/migration"
	"tonx-assignment/internal/app/routes"
	"tonx-assignment/pkg/utils"
)

func init() {
	// Load environment variable
	if err := godotenv.Load("./config/dev.env"); err != nil {
		log.Fatalf("load .env file fail: %v", err)
	}
}

func main() {

	run := os.Getenv("RUN")

	switch run {
	case "init":
		migration.Migrate()
		return
	case "clear":
		migration.Clear()
		return
	}

	go func() {
		cronJobs.Cron.Run()
	}()

	// Handle api route
	route := routes.Router()

	// Start running service
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if err := route.Run(port); err != nil {
		utils.Logger.LogOutput(err)
		log.Fatalf("running services fail: %v", err)
	}
}
