package usecase

import (
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"api.com/go-echo-rest-api/src/interface/dto/input"
	"api.com/go-echo-rest-api/src/interface/dto/output"
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

	// DBから取得したメッセージを件数分、レスポンスに変換する（koazeeを使わないパターン）
	var d = make([]*output.MessageSearch, len(messages), len(messages))
	for i, m := range messages {
		s := new(output.MessageSearch)
		s.Update(&m)
		d[i] = s
	}

	// ポインターを返す：値を返すとコピーされてメモリを食うため
	return &output.MessageSearchOutput{Data: d}, err
}
