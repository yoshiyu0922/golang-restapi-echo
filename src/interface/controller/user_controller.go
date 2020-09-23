package controller

import (
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto"
	"api.com/rest-base-api/src/interface/repository"
	"api.com/rest-base-api/src/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	Usecase usecase.UserUsecase
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
func NewUserController(handler *database.SqlHandler) *UserController {
	return &UserController{
		Usecase: usecase.UserUsecase{
			Repository: repository.NewUserRepository(handler),
		},
	}
}

func (u *UserController) Search(c echo.Context) (err error) {
	input := new(dto.UserSearchInput)
	if err := c.Bind(input); err != nil {
		return err
	}
	res, err := u.Usecase.Search(input)
	// Controller側でエラーハンドリングする
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
