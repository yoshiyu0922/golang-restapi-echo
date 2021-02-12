package usecase

import (
	"api.com/go-echo-rest-api/src/adapter/dto/input"
	"api.com/go-echo-rest-api/src/adapter/dto/output"
	"api.com/go-echo-rest-api/src/domain/models"
	"api.com/go-echo-rest-api/src/infrastructure/database"
	"github.com/wesovilabs/koazee"
)

type UserUsecase struct {
	Repository        UserRepository
	MessageRepository MessageRepository
}

func NewUserUsecase(sqlHandler *database.SqlHandler) *UserUsecase {
	return &UserUsecase{
		Repository:        NewUserRepository(sqlHandler),
		MessageRepository: NewMessageRepository(sqlHandler),
	}
}

func (u *UserUsecase) Search(user *input.UserSearchInput) (out *output.UserSearchOutput, err error) {
	users, err := u.Repository.Search(user)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return &output.UserSearchOutput{
			Data: make([]*output.UserSearch, 0),
		}, nil
	}

	// ユーザーIDのリストを抽出して、それらに紐づくメッセージを取得
	ids := koazee.StreamOf(users).
		Map(func(u models.User) models.UserIdType {
			return u.Id
		}).
		RemoveDuplicates().
		Do().Out().Val().([]models.UserIdType)
	messages, err := u.MessageRepository.FindByUserIds(&ids)
	if err != nil {
		return nil, err
	}

	// usersの件数分、messagesとuser_idでマッピングし、レスポンスに変換する
	result := koazee.StreamOf(users).Map(createOutput(messages)).Do().Out().Val().([]*output.UserSearch)

	// ポインターを返す：値を返すとコピーされてメモリを食うため
	return &output.UserSearchOutput{Data: result}, err
}

func createOutput(messages models.Messages) func(u models.User) *output.UserSearch {
	return func(u models.User) *output.UserSearch {
		res := new(output.UserSearch)
		// users.idに紐づくmessagesを抽出
		if messages != nil {
			target := koazee.StreamOf(messages).Filter(func(m models.Message) bool {
				return m.UserId == u.Id
			}).Do().Out().Val().([]models.Message)

			// レスポンスを更新
			res.Update(&u, target)
		} else {
			res.Update(&u, make([]models.Message, 0))
		}
		return res
	}
}
