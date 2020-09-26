package output

import (
	"api.com/go-echo-rest-api/src/domain/models"
	"github.com/wesovilabs/koazee"
)

type UserSearchOutput struct {
	Data []*UserSearch `json:"data"`
}

type UserSearch struct {
	Id              int                  `json:"user_id"`
	Name            string               `json:"name"`
	Age             int                  `json:"age"`
	JobLargeTypeId  *string              `json:"job_large_type_id"`
	JobMiddleTypeId *string              `json:"job_middle_type_id"`
	JobSmallTypeId  *string              `json:"job_small_type_id"`
	JobName         string               `json:"job_name"`
	JobTerm         int                  `json:"job_term"`
	Messages        []*UserSearchMessage `json:"messages"`
}

type UserSearchMessage struct {
	Id      int    `json:"message_id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

func (output *UserSearchMessage) Update(message *models.Message) {
	output.Id = message.Id
	output.UserId = message.UserId
	output.Title = message.Title
	output.Message = message.Message
}

func (output *UserSearch) Update(user *models.User, messages []models.Message) {
	output.Id = user.Id
	output.Name = user.Name
	output.Age = user.Age
	if user.JobLargeTypeId.Valid {
		output.JobLargeTypeId = &user.JobLargeTypeId.String
	} else {
		output.JobLargeTypeId = nil
	}
	if user.JobMiddleTypeId.Valid {
		output.JobMiddleTypeId = &user.JobMiddleTypeId.String
	} else {
		output.JobMiddleTypeId = nil
	}
	if user.JobSmallTypeId.Valid {
		output.JobSmallTypeId = &user.JobSmallTypeId.String
	} else {
		output.JobSmallTypeId = nil
	}
	output.JobName = user.JobName
	output.JobTerm = user.JobTerm

	messageRes := koazee.StreamOf(messages).Map(func(m models.Message) *UserSearchMessage {
		res := new(UserSearchMessage)
		res.Update(&m)
		return res
	}).Do().Out().Val().([]*UserSearchMessage)
	output.Messages = messageRes
}
