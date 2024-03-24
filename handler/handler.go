package handler

import (
	"echo-server/common/security"
	"echo-server/model"
	"echo-server/service"
	"echo-server/utility"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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

		apiContext := ctx.(*utility.ApiContext)

		userInput := new(model.User)
		err := ctx.Bind(userInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}

		if err := ctx.Validate(userInput); err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}
		creator, err := apiContext.GetUserId()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}
		userInput.CreateUser = creator

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

func (u *UserHandler) Login() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		loginInput := new(model.Login)

		err := ctx.Bind(loginInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}

		if err := ctx.Validate(loginInput); err != nil {
			return ctx.JSON(http.StatusBadRequest, "user not valid")
		}

		user, err := u.services.GetUserByUserNameAndPassword(*loginInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "user not found")
		}
		claims := &security.JtwClaims{
			UserName: user.UserName,
			UserId:   user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		st, err := token.SignedString([]byte("secret "))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, echo.Map{"token": st})
	}

}
