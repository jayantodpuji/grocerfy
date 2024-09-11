package test

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDBInstance *gorm.DB

func InitTestDB() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=grocerfy_test port=5432 sslmode=disable"
	TestDBInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := TestDBInstance.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func CloseTestDB() {
	if TestDBInstance != nil {
		sqlDB, err := TestDBInstance.DB()
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	InitTestDB()
}
