package service

import "echo-server/model"

type userService struct {
}

func NewUserService() UserService{
	return userService{}
}


func (u userService)GetUserList() ([]model.User, error) {
		userList := []model.User{
		{
			Name:   "kave",
			Family: "haj",
			Age:    31,
			Phone:  "09364646448",
		},
		{
			Name:   "james",
			Family: "hetfaild",
			Age:    45,
			Phone:  "0912345678",
		},
		{
			Name:   "anna",
			Family: "bani",
			Age:    32,
			Phone:  "0911345678",
		},
	}
	return userList, nil
}
