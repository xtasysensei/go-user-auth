package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/xtasysensei/go-poll/internal/config"
	"github.com/xtasysensei/go-poll/pkg/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("No command provided. Supported commands are 'up' and 'down'.")
	}
	cmd := os.Args[1]

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Unable to initialize migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Unable to initialize migration manager: %v", err)
	}

	switch cmd {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "force":
		if len(os.Args) < 3 {
			log.Fatalf("No version provided for force command.")
		}
		version, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Invalid version number: %v", err)
		}
		if err = m.Force(version); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	default:
		log.Fatalf("Invalid command. Supported commands are 'up' and 'down'.")
	}

	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No change in migration.")
		} else {
			log.Fatalf("Migration failed: %v", err)
		}
	} else {
		fmt.Printf("Migration %s successful.\n", cmd)
	}
}
