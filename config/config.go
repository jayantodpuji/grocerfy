package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	App    App
	DB     DB
	TestDB DB
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		App:    LoadApplicationConfig(),
		DB:     LoadDatabaseConfig(),
		TestDB: LoadTestDatabaseConfig(),
	}, nil
}
