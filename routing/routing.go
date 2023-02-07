package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {

	u := e.Group("/users")
	u.GET("/get/15", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/15")
	})

	u.GET("/get/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get")
	})

	u.GET("/get/:id/avatar", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/avatar")
	})


	return nil
}
