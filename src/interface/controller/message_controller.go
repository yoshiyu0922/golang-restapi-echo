package controller

import (
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/repository"
	"api.com/rest-base-api/src/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MessageController struct {
	Usecase usecase.MessageUsecase
}

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
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}
