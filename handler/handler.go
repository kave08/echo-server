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

// UserHandler is responsible for handling user-related HTTP requests.
type UserHandler struct {
	services service.UserService
}

// UserResponse represents the response structure for creating a new user
type UserResponse struct {
	NewUserID string
}

// NewUserHandler creates a new instance of UserHandler with the provided UserService.
func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		services: service,
	}
}

// CreateUser creates a new user in the system.
//
// It expects a valid user input in the request body and returns the ID of the newly created user.
func (u *UserHandler) CreateUser() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {

		apiContext := ctx.(*utility.APIContext)

		userInput := new(model.User)
		err := ctx.Bind(userInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}

		if err := ctx.Validate(userInput); err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}
		creator, err := apiContext.GetUserID()
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, " ")
		}
		userInput.CreateUser = creator

		id, err := u.services.CreateUser(*userInput)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		res := UserResponse{
			NewUserID: id,
		}

		return ctx.JSON(http.StatusOK, res.NewUserID)
	}
}

// GetUserList retrieves a list of all users from the database.
//
// It returns a function that handles the HTTP request and response.
func (u *UserHandler) GetUserList() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		userList, err := u.services.GetUserList()
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, userList)
	}
}

// Login authenticates a user based on the provided username and password.
//
// It returns a JWT token if the authentication is successful.
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
