package utility

import (
	"echo-server/common/security"
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// APIContext is a custom context that embeds the echo.Context and provides additional functionality.
type APIContext struct {
	echo.Context
}

// GetUserID retrieves the user ID from the JWT token stored in the context.
//
// It returns an empty string and an error if the user is not logged in.
func (a APIContext) GetUserID() (userID string, err error) {
	defer func() {
		r := recover()
		if r != nil {
			userID = ""
			err = errors.New("user is not login")

		}
	}()
	token := a.Get("user").(*jwt.Token)

	claim := token.Claims.(*security.JtwClaims)

	return claim.UserId, nil
}
