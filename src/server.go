package src

import (
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/infrastructure/rest_api"
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
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	allowOrigin := strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")

	// Connect DB
	sqlHandler := database.NewSqlHandler(user, password, host, dbPort, dbName)
	defer sqlHandler.Conn.Close() // 遅延評価：アプリを終了したらクローズする

	// Server Run
	rest_api.Run(env, port, allowOrigin, sqlHandler)
}
