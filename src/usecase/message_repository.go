package usecase

import "api.com/rest-base-api/src/domain/models"

type MessageRepository interface {
	FindAll() (messages models.Messages, err error)

	FindByTitle(title *string) (messages models.Messages, err error)
}
