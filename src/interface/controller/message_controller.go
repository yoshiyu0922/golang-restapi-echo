package controller

import (
	_ "api.com/rest-base-api/docs"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto"
	"api.com/rest-base-api/src/interface/repository"
	"api.com/rest-base-api/src/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MessageController struct {
	Usecase usecase.MessageUsecase
}

// search messages.
// @Summary search messages
// @Description search messages
// @Accept  json
// @Produce  json
// @Param message_id query int false "メッセージID"
// @Param user_id query int false "ユーザID"
// @Param title query string false "タイトル"
// @Param message query string false "メッセージ"
// @Success 200 {array} models.Message
// @Failure 500 {object} error_handling.APIError
// @Router /message [get]
func NewMessageController(sqlHandler *database.SqlHandler) *MessageController {
	return &MessageController{
		Usecase: usecase.MessageUsecase{
			Repository: repository.NewMessageRepository(sqlHandler),
		},
	}
}

func (controller *MessageController) SearchMessage(c echo.Context) error {
	input := new(dto.MessageSearchInput)
	if err := c.Bind(input); err != nil {
		return err
	}
	res, err := controller.Usecase.Search(input)
	// Controller側でエラーハンドリングする
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
