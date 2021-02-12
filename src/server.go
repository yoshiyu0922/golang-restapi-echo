package src

import (
	 "api.com/go-echo-rest-api/src/infrastructure/config"
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"api.com/go-echo-rest-api/src/infrastructure/rest_api"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func Start() {
	// read Environment Variables
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Connect DB
	sqlHandler := database.NewSqlHandler(user, password, host, dbPort, dbName)
	defer sqlHandler.Conn.Close() // 遅延評価：アプリを終了したらクローズする

	appConfig := &config.Application{
		Environment: os.Getenv("ENV"),
		Port: os.Getenv("PORT"),
		AllowOrigins: strings.Split(os.Getenv("ALLOW_ORIGINS"), ","),
	}

	api := rest_api.Initialize(appConfig, sqlHandler)

	api.Logger.Fatal(api.Start(":" + port))
}
