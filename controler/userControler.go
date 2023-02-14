package controler

import (
	"echo-server/model"
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

	return c.JSON(http.StatusOK, userList)
}
