package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto/input"
	"api.com/rest-base-api/src/interface/dto/output"
	"github.com/wesovilabs/koazee"
)

type MessageUsecase struct {
	// interface/databaseのリポジトリに直接参照しないようにする
	Repository MessageRepository
}

func NewMessageUsecase(sqlHandler *database.SqlHandler) *MessageUsecase {
	return &MessageUsecase{
		Repository: NewMessageRepository(sqlHandler),
	}
}

func (usecase *MessageUsecase) Search(input *input.MessageSearchInput) (out *output.MessageSearchOutput, err error) {
	messages, err := usecase.Repository.Search(input)
	if err != nil {
		return nil, err
	}

	// DBから取得したメッセージを件数分、レスポンスに変換する
	var result = make([]*output.MessageSearch, 0)
	if messages != nil {
		result = koazee.StreamOf(messages).Map(func(msg models.Message) *output.MessageSearch {
			res := new(output.MessageSearch)
			res.Update(&msg)
			return res
		}).Do().Out().Val().([]*output.MessageSearch)
	}

	// ポインターを返す：値を返すとコピーされてメモリを食うため
	return &output.MessageSearchOutput{Data: result}, err
}
