package utility

import (
	"echo-server/common/security"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ApiContext struct {
	echo.Context
}

func (a ApiContext) GetUserId() (string, error) {
	token := a.Get("user").(*jwt.Token)

	claim := token.Claims.(*security.JtwClaims)

	return claim.UserId, nil
}
