package controller

import (
	"api.com/go-echo-rest-api/src/adapter/dto/input"
	"api.com/go-echo-rest-api/src/infrastructure/rest_api/context"
	"api.com/go-echo-rest-api/src/usecase"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// search users.
// @Summary search users
// @Description search users
// @Accept  json
// @Produce  json
// @Param user_id query string false "ユーザーID"
// @Param name query string false "名前"
// @Param age query string false "年齢"
// @Param job_large_type_id query string false "職種大分類"
// @Param job_middle_type_id query string false "職種中分類"
// @Param job_small_type_id query string false "職種小分類"
// @Param job_name query string false "職種名"
// @Param job_term query string false "就業期間"
// @Param message_id query string false "メッセージID"
// @Success 200 {array} models.User
// @Failure 500 {object} error_handling.APIError
// @Router /user [get]
func (u *UserController) Search(c echo.Context) (err error) {
	cc, _ := c.(*context.CustomContext)
	req := new(input.UserSearchInput)
	if err := cc.Bind(req); err != nil {
		return errors.WithStack(err) // 必ずstacktraceをつけてエラーを返す
	}

	res, err := usecase.NewUserUsecase(cc.DB).Search(req)
	// Controller側でエラーハンドリングする
	if err != nil {
		return errors.WithStack(err)
	}

	return c.JSON(http.StatusOK, res)
}
