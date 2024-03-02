package repository

import "echo-server/model"

type Repository interface {
	GetUserList() ([]model.User, error)
	GetUserId(id string) (*model.User, error)
	InsertUser(user model.User) (string, error)
}
