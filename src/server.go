package src

import (
	"api.com/rest-base-api/src/infrastructure"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/error_handling"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
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

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Error Handling
	e.HTTPErrorHandler = error_handling.JSONErrorHandler

	// Routes
	infrastructure.Routing(e, sqlHandler)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
