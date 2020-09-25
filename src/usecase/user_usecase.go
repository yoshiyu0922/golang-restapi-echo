package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/infrastructure/database"
	"api.com/rest-base-api/src/interface/dto/input"
	"api.com/rest-base-api/src/interface/dto/output"
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
	ids := koazee.StreamOf(users).Map(func(u models.User) int {
		return u.Id
	}).RemoveDuplicates().Do().Out().Val().([]int)
	messages, err := u.MessageRepository.FindByUserIds(&ids)
	if err != nil {
		return nil, err
	}

	// usersの件数分、messagesとuser_idでマッピングし、レスポンスに変換する
	result := koazee.StreamOf(users).Map(func(u models.User) *output.UserSearch {
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
	}).Do().Out().Val().([]*output.UserSearch)

	// ポインターを返す：値を返すとコピーされてメモリを食うため
	return &output.UserSearchOutput{Data: result}, err
}
