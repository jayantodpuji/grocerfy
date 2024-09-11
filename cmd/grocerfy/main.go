package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	defer func() {
		if err := CloseDatabase(db); err != nil {
			log.Fatal(err)
		}
	}()

	app, err := internal.NewApplication(db, echo.New(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	internal.Routes(app)

	go func() {
		if err := app.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down the application")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(100)*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
