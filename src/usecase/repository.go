package usecase

import (
	"api.com/go-echo-rest-api/src/adapter/dto/input"
	"api.com/go-echo-rest-api/src/adapter/repository"
	"api.com/go-echo-rest-api/src/domain/models"
	"api.com/go-echo-rest-api/src/infrastructure/database"
)

type MessageRepository interface {
	FindAll() (messages models.Messages, err error)

	Search(title *input.MessageSearchInput) (messages models.Messages, err error)

	FindByUserIds(userIds *[]models.UserIdType) (messages models.Messages, err error)
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
