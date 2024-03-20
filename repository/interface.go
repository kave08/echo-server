package repository

import "echo-server/model"

type Repository interface {
	GetUserList() ([]model.User, error)
	GetUserId(id string) (*model.User, error)
	InsertUser(user model.User) (string, error)
	UpdateUserById(user model.User) error
	DeleteUserById(id string) error
	GetUserByUserNameAndPassword(login model.Login) (*model.User, error)
}
