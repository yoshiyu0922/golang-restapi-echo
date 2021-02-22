package controller

import (
	_ "api.com/go-echo-rest-api/docs"
	"api.com/go-echo-rest-api/src/adapter/dto/input"
	"api.com/go-echo-rest-api/src/core/error_handling"
	"api.com/go-echo-rest-api/src/infrastructure/rest_api/context"
	"api.com/go-echo-rest-api/src/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"unicode/utf8"
)

type MessageController struct{}

func NewMessageController() *MessageController {
	return &MessageController{}
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
func (controller MessageController) SearchMessage(c echo.Context) error {
	cc, _ := c.(*context.CustomContext)
	// リクエストパラメータと構造体をバインドする
	req := new(input.MessageSearchInput)
	if err := cc.Bind(req); err != nil {
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
		e := error_handling.NewValidationError(validErrors)
		return errors.WithStack(e)
	}

	result, err := usecase.NewMessageUsecase(cc.DB).Search(req)
	// Controller側でエラーハンドリングする
	if err != nil {
		return errors.WithStack(err)
	}

	return c.JSON(http.StatusOK, result)
}
