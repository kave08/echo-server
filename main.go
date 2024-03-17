package main

import (
	"echo-server/config"
	"echo-server/handler"
	"echo-server/database"
	"echo-server/repository"
	"echo-server/routing"
	"echo-server/service"
	"log"

	"github.com/labstack/echo/v4"
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

	//init server
	server := echo.New()

	//init repo
	repo := repository.NewUserService()

	//init service
	srv := service.NewUserService(repo)

	//init handler
	hndlr := handler.NewUserHandler(srv)

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
