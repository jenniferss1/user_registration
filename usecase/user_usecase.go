package usecase

import (
	"fmt"
	"modulo/model"
	"modulo/repository"
)

type UserUsecase struct {
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

func (uc *UserUsecase) CreateUser(user model.Users) (model.Users, error) {
	userId, err := uc.repository.CreateUser(user)
	if err != nil {
		return model.Users{}, err
	}

	user.ID = userId
	return user, nil
}

func (uc *UserUsecase) GetUserByID(userID int) (*model.Users, error) {
	user, err := uc.repository.GetUserById(userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
