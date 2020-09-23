package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/interface/dto"
)

type MessageUsecase struct {
	// interface/databaseのリポジトリに直接参照しないようにする
	Repository MessageRepository
}

func (usecase *MessageUsecase) Search(input *dto.MessageSearchInput) (messages []models.Message, err error) {
	messages, err = usecase.Repository.Search(input)
	if err != nil {
		return nil, err
	}
	//for m := range messages {
	// TODO: ユーザーIDからメッセージ一覧を取得
	//}
	return
}
