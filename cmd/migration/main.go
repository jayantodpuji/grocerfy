package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jayantodpuji/grocerfy/config"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	mainDB, err := InitiateDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := migrateDatabase(mainDB, "main"); err != nil {
		log.Fatal(err)
	}

	if cfg.App.Env == "development" {
		testDB, err := InitiateTestDatabase(cfg)
		if err != nil {
			log.Fatal(err)
		}

		if err := migrateDatabase(testDB, "test"); err != nil {
			log.Fatal(err)
		}
	}
}

func migrateDatabase(gormDB *gorm.DB, dbType string) error {
	db, err := gormDB.DB()
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("No new migrations to apply for %s database.", dbType)
		} else {
			return fmt.Errorf("migration failed for %s database: %v", dbType, err)
		}
	} else {
		log.Printf("Migration applied successfully for %s database.", dbType)
	}

	return nil
}
