package main

import (
	"echo-server/config"
	"echo-server/database"
	"echo-server/routing"
	"fmt"
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

	fmt.Println(config.Cfg.Port)

	//init server
	server := echo.New()

	// routing
	err = routing.SetRouting(server)

	//start server
	err = server.Start(":" + config.Cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
