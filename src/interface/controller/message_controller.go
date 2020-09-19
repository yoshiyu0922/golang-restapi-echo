package controller

import (
	_ "api.com/rest-base-api/docs"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/repository"
	"api.com/rest-base-api/src/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MessageController struct {
	Usecase usecase.MessageUsecase
}

// getUsers is getting users.
// @Summary search messages
// @Description search messages
// @Accept  json
// @Produce  json
// @Param title query string false "タイトル"
// @Success 200 {array} models.Message
// @Failure 500 {object} error_handling.APIError
// @Router /message [get]
func NewMessageController(sqlHandler *database.SqlHandler) *MessageController {
	return &MessageController{
		Usecase: usecase.MessageUsecase{
			MessageRepository: repository.NewMessageRepository(sqlHandler),
		},
	}
}

func (controller *MessageController) SearchMessage(c echo.Context) error {
	title := c.QueryParam("title")
	res, err := controller.Usecase.FindByTitle(&title)
	// Controller側でエラーハンドリングする
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
