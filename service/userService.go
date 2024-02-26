package service

import (
	"echo-server/model"
	"echo-server/repository"
)

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (u userService) GetUserList() ([]model.User, error) {

	userRepo := repository.NewUserService()
	userList, err := userRepo.GetUserList()
	if err != nil {
		return nil, err
	}
	return userList, nil
}
