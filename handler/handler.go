package handler

import (
	"echo-server/model"
	"echo-server/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	services service.UserService
}

type UserResponse struct {
	NewUserId string
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		services: service,
	}
}

func (u *UserHandler) CreateUser() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		userInput := new(model.User)
		err := ctx.Bind(userInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}

		id, err := u.services.CreateUser(*userInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		res := UserResponse{
			NewUserId: id,
		}

		return ctx.JSON(http.StatusOK, res.NewUserId)
	}
}

func (u *UserHandler) GetUserList() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userList, err := u.services.GetUserList()
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, userList)
	}
}
