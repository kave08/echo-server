package controler

import (
	"echo-server/model"
	"echo-server/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	userInput := new(model.User)
	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	fmt.Println(userInput.Name, userInput.Family, userInput.Age, userInput.Phone)
	return c.String(http.StatusOK, "/users")
}

func GetUserId(c echo.Context) error {
	return c.String(http.StatusOK, "/user/get")
}

func GetUserAvatar(c echo.Context) error {
	idString := c.Param("id")
	return c.String(http.StatusOK, "/user/get/avatar/[with param ("+idString+")]/avatar")
}

func CreateUser(c echo.Context) error {
	userInput := new(model.User)
	err := c.Bind(userInput)
	if err != nil {
		return err
	}
	fmt.Println(userInput.Name, userInput.Family, userInput.Age, userInput.Phone)
	return c.String(http.StatusOK, "/users")
}

func GetUserList(c echo.Context) error {

	userService := service.NewUserService()
	
	userList, err := userService.GetUserList()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userList)
}
