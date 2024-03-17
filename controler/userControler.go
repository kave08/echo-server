package controler

import (
	"echo-server/model"
	"echo-server/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRequest struct {
	services service.UserService
}

type UserResponse struct {
	NewUserId string
}

func NewUserRequest(service service.UserService) *UserRequest {
	return &UserRequest{
		services: service,
	}
}

func (u UserRequest) CreateUser(c echo.Context) error {
	userInput := new(model.User)
	err := c.Bind(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, " ")
	}

	id, err := u.services.CreateUser(*userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	res := UserResponse{
		NewUserId: id,
	}

	return c.JSON(http.StatusOK, res.NewUserId)
}

func (u UserRequest) GetUserList(c echo.Context) error {

	userList, err := u.services.GetUserList()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userList)
}
