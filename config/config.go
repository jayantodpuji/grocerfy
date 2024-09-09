package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	DB
	App
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:  LoadDatabaseConfig(),
		App: LoadApplicationConfig(),
	}, nil
}
