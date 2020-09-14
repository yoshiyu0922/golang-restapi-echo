package infrastructure

import (
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/controller"
	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo, sqlHandler *database.SqlHandler) {

	msg := controller.NewMessageController(sqlHandler)
	e.GET("/", msg.SearchMessage)
}
