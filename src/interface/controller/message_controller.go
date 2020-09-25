package controller

import (
	_ "api.com/rest-base-api/docs"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto/input"
	"api.com/rest-base-api/src/interface/error_handling"
	"api.com/rest-base-api/src/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"unicode/utf8"
)

type MessageController struct {
	Usecase *usecase.MessageUsecase
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
		Usecase: usecase.NewMessageUsecase(sqlHandler),
	}
}

func (controller MessageController) SearchMessage(c echo.Context) error {
	// リクエストパラメータと構造体をバインドする
	req := new(input.MessageSearchInput)
	if err := c.Bind(req); err != nil {
		return errors.WithStack(err) // 必ずstacktraceをつけてエラーを返す
	}

	// バリデーションチェック
	var validErrors = make([]error_handling.ValidationErrorDetail, 0)
	if req.Title != nil && utf8.RuneCountInString(*req.Title) > 20 {
		validErrors = append(validErrors, error_handling.ValidationErrorDetail{
			Item:    "title",
			Message: fmt.Sprintf("タイトルは20文字以内で入力してください。: %d", utf8.RuneCountInString(*req.Title)),
		})
	}
	if req.Message != nil && utf8.RuneCountInString(*req.Message) > 50 {
		validErrors = append(validErrors, error_handling.ValidationErrorDetail{
			Item:    "message",
			Message: fmt.Sprintf("メッセージは50文字以内で入力してください。: %d", utf8.RuneCountInString(*req.Message)),
		})
	}
	if len(validErrors) > 0 {
		e := &error_handling.ValidationError{
			Err:     errors.New("validation error in message_controller.SearchMessage"),
			Details: validErrors,
		}
		return errors.WithStack(e)
	}

	result, err := controller.Usecase.Search(req)
	// Controller側でエラーハンドリングする
	if err != nil {
		return errors.WithStack(err)
	}

	return c.JSON(http.StatusOK, result)
}
