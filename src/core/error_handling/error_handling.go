package error_handling

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"net/http"
)

/**
Application Error Interface
*/
type ApplicationErrorInterface interface {
	AppError() *ApplicationError
}

/**
Application Error
*/
type ApplicationError struct {
	Err     error  `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details interface{} `json:"details"`
}

func (err *ApplicationError) AppError() *ApplicationError {
	return err
}

func JSONErrorHandler(err error, c echo.Context) {
	var apiError ApplicationError

	cause := errors.Cause(err)
	switch cause.(type) {
	case *echo.HTTPError:
		if he, ok := cause.(*echo.HTTPError); ok {
			apiError.Code = he.Code
			apiError.Message = he.Message.(string)
		}
	case *ValidationError:
		if he, ok := cause.(*ValidationError); ok {
			apiError.Code = 404
			apiError.Message = "入力内容に誤りがあります。"
			apiError.Details = he.Details
			err = he.Err
		}
	default:
		if validationError, ok := cause.(ApplicationErrorInterface); ok {
			apiError = *validationError.AppError()
		} else {
			apiError.Code = http.StatusInternalServerError
			apiError.Err = err
			apiError.Message = err.Error()
		}
	}

	log.Error(fmt.Sprintf("%+v", err)) // log.Fatalはプログラムを強制終了するので気を付ける
	c.JSON(apiError.Code, apiError)
}
