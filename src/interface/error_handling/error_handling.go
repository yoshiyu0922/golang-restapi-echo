package error_handling

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type APIError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func JSONErrorHandler(err error, c echo.Context) {
	var apierr APIError
	switch err.(type) {
	case *echo.HTTPError:
		if he, ok := err.(*echo.HTTPError); ok {
			apierr.Code = he.Code
			apierr.Message = he.Message.(string)
		}
	case *ValidationError:
		if he, ok := err.(*ValidationError); ok {
			apierr.Code = he.Code
			apierr.Message = he.Message
			apierr.Details = he.Details
		}
	default:
		apierr.Code = http.StatusInternalServerError
		apierr.Message = err.Error()
	}

	log.Error(err) // log.Fatalはプログラムを強制終了するので気を付ける
	c.JSON(apierr.Code, apierr)
}
