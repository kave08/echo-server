package service

import (
	"echo-server/model"
	"echo-server/repository"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewUserService(repository repository.Repository) UserService {
	return Service{
		repo: repository,
	}
}

func (u Service) GetUserList() ([]model.User, error) {

	userList, err := u.repo.GetUserList()
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (u Service) CreateUser(user model.User) (string, error) {

	userInput := model.User{
		FirstName:  user.FirstName,
		LastName:   user.UserName,
		Age:        user.Age,
		Email:      user.Email,
		Phone:      user.Phone,
		UserName:   user.UserName,
		Password:   user.Password,
		Created_at: time.Now(),
	}

	userId, err := u.repo.InsertUser(userInput)
	if err != nil {
		return "", err
	}

	return userId, nil
}
