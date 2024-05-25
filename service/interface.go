package service

import "echo-server/model"

type UserService interface {
	CreateUser(user model.User) (string, error)
	GetUserList() ([]model.User, error)
	GetUserByUserNameAndPassword(login model.Login) (*model.User, error)
	IsUserValidForAccess(userID, roleName string) bool
}
