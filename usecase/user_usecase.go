package usecase

import (
	"modulo/model"
	"modulo/repository"
)

type UserUsecase struct {
	// repository
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (uc *UserUsecase) GetUsers() ([]model.Users, error) {
	return uc.repository.GetUsers()
}
