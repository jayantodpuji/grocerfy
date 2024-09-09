package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jayantodpuji/grocerfy/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := InitiateDatabase(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		log.Println("Migration applied successfully.")
	}
}
