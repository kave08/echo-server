package service

import "echo-server/model"

type UserService interface {
	GetUserList() ([]model.User, error)
}
