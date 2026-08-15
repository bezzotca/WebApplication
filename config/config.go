// Package config assembles app-wide configuration from the environment.
package config

import (
	"os"
	"strconv"
	"time"

	"webproject/pkg/database/postgres"
)

type Config struct {
	HTTPPort string
	Postgres postgres.Config
}

func Load() Config {
	return Config{
		HTTPPort: getEnv("HTTP_PORT", "8080"),
		Postgres: postgres.Config{
			Host:            getEnv("DB_HOST", "localhost"),
			Port:            getEnvInt("DB_PORT", 5432),
			User:            getEnv("DB_USER", "postgres"),
			Password:        getEnv("DB_PASSWORD", "postgres"),
			Database:        getEnv("DB_NAME", "webproject"),
			SSLMode:         getEnv("DB_SSLMODE", "disable"),
			MaxConns:        int32(getEnvInt("DB_MAX_CONNS", 10)),
			MinConns:        int32(getEnvInt("DB_MIN_CONNS", 2)),
			MaxConnLifetime: time.Hour,
			MaxConnIdleTime: 30 * time.Minute,
		},
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}
