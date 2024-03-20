package routing

import (
	"echo-server/common/security"
	"echo-server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo, handler *handler.UserHandler) error {

	jwtCfg := middleware.JWTConfig{
		SigningKey: []byte("secret "),
		Claims:     security.JtwClaims{},
	}
	e.POST("/login", handler.Login(), middleware.JWTWithConfig(jwtCfg))

	u := e.Group("/users")
	u.GET("/getlist", handler.GetUserList())
	u.POST("/create", handler.CreateUser())

	return nil
}
