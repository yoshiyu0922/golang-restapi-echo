package output

import (
	"api.com/go-echo-rest-api/src/domain/models"
)

type MessageSearchOutput struct {
	Data []*MessageSearch `json:"data"`
}
type MessageSearch struct {
	Id      models.MessageIdType `json:"message_id"`
	UserId  models.UserIdType    `json:"user_id"`
	Title   string               `json:"title"`
	Message string               `json:"message"`
}

func (output *MessageSearch) Update(message *models.Message) {
	output.Id = message.Id
	output.UserId = message.UserId
	output.Title = message.Title
	output.Message = message.Message
}
