package main

import (
	"echo-server/config"
	"echo-server/routing"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	//confi
	err := config.LoadConfig()
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
