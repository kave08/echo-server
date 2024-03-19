package main

import (
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

	//middleware
	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	//validation
	server.Validator = utility.CustomValidator{
		Validator: validator.New(),
	}

	// routing
	err = routing.SetRouting(server, hndlr)
	if err != nil {
		log.Fatal(err)
	}

	//start server
	err = server.Start(":" + config.Cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
