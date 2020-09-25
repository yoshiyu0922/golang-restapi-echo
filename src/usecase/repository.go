package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto/input"
	"api.com/rest-base-api/src/interface/repository"
)

type MessageRepository interface {
	FindAll() (messages models.Messages, err error)

	Search(title *input.MessageSearchInput) (messages models.Messages, err error)

	FindByUserIds(userIds *[]int) (messages models.Messages, err error)
}

func NewMessageRepository(sqlHandler *database.SqlHandler) MessageRepository {
	return &repository.MessageRepository{
		SqlHandler: sqlHandler,
	}
}

type UserRepository interface {
	Search(*input.UserSearchInput) (users models.Users, err error)
}

func NewUserRepository(sqlHandler *database.SqlHandler) UserRepository {
	return &repository.UserRepository{
		SqlHandler: sqlHandler,
	}
}
