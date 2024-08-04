package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/xtasysensei/go-poll/internal/config"
)

var DB *sql.DB

func Init(cfg *config.Config) {
	var err error
	DB, err = Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func Connect(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.DatabaseURL()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}
