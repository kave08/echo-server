package service

import (
	"echo-server/model"
	"echo-server/repository"

	"time"

	"golang.org/x/exp/slices"
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

func (u Service) GetUserByUserNameAndPassword(login model.Login) (*model.User, error) {

	user, err := u.repo.GetUserByUserNameAndPassword(login)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u Service) IsUserValidForAccess(userID, roleName string) bool {

	user, err := u.repo.GetUserId(userID)
	if err != nil {
		return false
	}

	if user.Roles == nil {
		return false
	}

	roleIndex := slices.IndexFunc(user.Roles, func(role string) bool {
		return role == roleName
	})

	return roleIndex >= 0
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
		CreateUser: user.CreateUser,
		Roles:      user.Roles,
		CreatedAt:  time.Now(),
	}

	userId, err := u.repo.InsertUser(userInput)
	if err != nil {
		return "", err
	}

	return userId, nil
}
