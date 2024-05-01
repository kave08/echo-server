package main

import (
	"echo-server/common/security"
	"echo-server/config"
	"echo-server/database"
	"echo-server/handler"
	"echo-server/repository"
	"echo-server/routing"
	"echo-server/service"
	"echo-server/utility"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//database
	_, err := database.InitMongo()
	if err != nil {
		log.Fatal(err)
	}

	//config
	err = config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	//init repo
	repo := repository.NewUserService()

	//init service
	srv := service.NewUserService(repo)

	//init handler
	hndlr := handler.NewUserHandler(srv)

	//init server
	server := echo.New()

	// routing
	err = routing.SetRouting(server, hndlr)
	if err != nil {
		log.Fatal(err)
	}

	//middleware
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &utility.ApiContext{Context: c}
			return next(apiContext)
		}
	})

	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:             []byte("secret"),
		Claims:                 &security.JtwClaims{},
		ContinueOnIgnoredError: true,
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return nil
		},
	}))

	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	//validation
	server.Validator = utility.CustomValidator{
		Validator: validator.New(),
	}

	//start server
	err = server.Start(":" + config.Cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
