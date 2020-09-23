package infrastructure

import (
	_ "api.com/rest-base-api/docs"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/controller"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Routing(e *echo.Echo, sqlHandler *database.SqlHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	msg := controller.NewMessageController(sqlHandler)
	user := controller.NewUserController(sqlHandler)
	e.GET("/message", msg.SearchMessage)
	e.GET("/user", user.Search)
}
