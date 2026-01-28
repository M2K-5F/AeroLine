package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetDatabaseDSNFromEnv() (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf(".env file not found")
	}
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_NAME", "testdb"),
		getEnv("MAX_CONNECTIONS", "15"),
	), nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal

}
