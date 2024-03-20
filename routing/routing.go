package routing

import (
	"echo-server/handler"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo, handler *handler.UserHandler) error {

	e.POST("/login", handler.Login())

	u := e.Group("/users")
	u.GET("/getlist", handler.GetUserList())
	u.POST("/create", handler.CreateUser())

	return nil
}
