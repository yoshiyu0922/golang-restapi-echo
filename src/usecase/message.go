package usecase

import (
	"api.com/rest-base-api/src/domain/models"
)

type MessageUsecase struct {
	// interface/databaseのリポジトリに直接参照しないようにする
	MessageRepository MessageRepository
}

func (usecase *MessageUsecase) FindByTitle(title *string) (messages []models.Message, err error) {
	messages, err = usecase.MessageRepository.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return
}
