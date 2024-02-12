package routing

import (
	"echo-server/controler"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	u := e.Group("/users")
	u.GET("/get",controler.GetUser)

	u.GET("/get/:id", controler.GetUserId)

	u.GET("/getlist", controler.GetUserList)

	u.GET("/get/:id/avatar", controler.GetUserAvatar)

	u.POST("/create", controler.CreateUser)

	return nil
}
