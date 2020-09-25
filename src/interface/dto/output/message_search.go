package output

import (
	"api.com/rest-base-api/src/domain/models"
)

type MessageSearchOutput struct {
	Data []*MessageSearch `json:"data"`
}
type MessageSearch struct {
	Id      int    `json:"message_id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (output *MessageSearch) Update(message *models.Message) {
	output.Id = message.Id
	output.UserId = message.UserId
	output.Title = message.Title
	output.Message = message.Message
}
