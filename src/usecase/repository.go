package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/interface/dto"
)

type MessageRepository interface {
	FindAll() (messages models.Messages, err error)

	Search(title *dto.MessageSearchInput) (messages models.Messages, err error)
}

type UserRepository interface {
	Search(*dto.UserSearchInput) (users models.Users, err error)
}
