package config

import (
	"fmt"
	"os"
)

type App struct {
	Port   string
	Env    string
	JWTKey string
}

func LoadApplicationConfig() App {
	return App{
		Port:   getAppPort(),
		Env:    getAppEnv(),
		JWTKey: getJWTKey(),
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

func getJWTKey() string {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		key = "secret"
	}

	return key
}
