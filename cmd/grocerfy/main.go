package main

import (
	"log"

	"github.com/jayantodpuji/grocerfy/config"
	"github.com/jayantodpuji/grocerfy/internal"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := InitiateDatabase(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	app, err := internal.NewApplication(db, echo.New(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Start()
}
