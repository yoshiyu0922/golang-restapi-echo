package rest_api

import (
	_ "api.com/go-echo-rest-api/docs"
	"api.com/go-echo-rest-api/src/infrastructure/config"
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"api.com/go-echo-rest-api/src/adapter/controller"
	"api.com/go-echo-rest-api/src/core/error_handling"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

func Initialize(
	appConfig *config.Application,
	sqlHandler *database.SqlHandler,
) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     appConfig.AllowOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// リクエストヘッダー認証
	e.Use(requestHeaderMiddleware(appConfig))

	// Error Handling
	e.HTTPErrorHandler = error_handling.JSONErrorHandler

	// instance Controllers
	msg := controller.NewMessageController(sqlHandler)
	user := controller.NewUserController(sqlHandler)

	// Routing
	if appConfig.Environment != "production" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	e.GET("/message", msg.SearchMessage)
	e.GET("/user", user.Search)

	return e
}

func requestHeaderMiddleware(applicationConfig *config.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Originヘッダの中身を取得
			origin := c.Request().Header.Get(echo.HeaderOrigin)
			// 許可しているOriginの中で、リクエストヘッダのOriginと一致するものがあれば処理を継続
			for _, o := range applicationConfig.AllowOrigins {
				if origin == o || o == "*" {
					return next(c)
				}
			}
			// 一致しているものがなかった場合は403(Forbidden)を返却する
			return c.String(http.StatusForbidden, "forbidden")
		}
	}
}
