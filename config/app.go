package config

import (
	"fmt"
	"os"
)

type App struct {
	Port string
	Env  string
}

func LoadApplicationConfig() App {
	return App{
		Port: getAppPort(),
		Env:  getAppEnv(),
	}
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "development"
	}

	return env
}

func getAppPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf(":%s", port)
}
