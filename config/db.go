package config

import "os"

type DB struct {
	Host     string
	User     string
	Password string
	Database string
	SSL      string
}

func LoadDatabaseConfig() DB {
	return DB{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSL:      os.Getenv("DB_SSL"),
	}
}

func LoadTestDatabaseConfig() DB {
	return DB{
		Host:     os.Getenv("TEST_DB_HOST"),
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASSWORD"),
		Database: os.Getenv("TEST_DB_NAME"),
		SSL:      os.Getenv("TEST_DB_SSL"),
	}
}
