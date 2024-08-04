package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs, _ = LoadConfig()

type Config struct {
	ServerAddress          string
	Postgres               PostgresConfig
	JWTExpirationInSeconds int64
	JWTSecret              string
}

type PostgresConfig struct {
	Server   string
	Port     string
	DBName   string
	User     string
	Password string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		ServerAddress: ":" + getEnv("SERVER_PORT", "8000"),
		Postgres: PostgresConfig{
			Server:   getEnv("POSTGRES_SERVER", "localhost"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
			DBName:   getEnv("POSTGRES_DB", "go_app"),
			User:     getEnv("POSTGRES_USER", "sensei"),
			Password: getEnv("POSTGRES_PASSWORD", "12345"),
		},
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-secret-secret-anymore"),
	}
	return cfg, nil
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.Postgres.User, c.Postgres.Password, c.Postgres.Server, c.Postgres.Port, c.Postgres.DBName)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
