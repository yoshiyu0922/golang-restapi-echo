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
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	var apierr APIError
	switch err.(type) {
	case *echo.HTTPError:
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			msg = he.Message.(string)
		}
		apierr.Code = code
		apierr.Message = msg
	case *ValidationError:
		if he, ok := err.(*ValidationError); ok {
			apierr.Code = he.Code
			apierr.Message = he.Details[0].Item
			apierr.Details = he.Details
		}
	}

	c.JSON(code, apierr)
	log.Debug(err)
}
