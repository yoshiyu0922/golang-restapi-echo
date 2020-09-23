package usecase

import (
	"api.com/rest-base-api/src/domain/models"
	"api.com/rest-base-api/src/interface/dto"
)

type UserUsecase struct {
	Repository UserRepository
}

func (u *UserUsecase) Search(user *dto.UserSearchInput) (users models.Users, err error) {
	users, err = u.Repository.Search(user)
	return
}
