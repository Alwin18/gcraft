package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	DBHost          string
	DBUser          string
	DBPassword      string
	DBName          string
	DBPort          string
	ServerPort      string
	SetMaxIdleConns string
	SetMaxOpenConns string
	SetMaxLifeTime  string
}

func LoadEnv() (*Config, error) {
	config := &Config{}
	envVars := map[string]*string{
		"DB_HOST":      &config.DBHost,
		"DB_PORT":      &config.DBPort,
		"DB_USER":      &config.DBUser,
		"DB_PASS":      &config.DBPassword,
		"DB_NAME":      &config.DBName,
		"SERVER_PORT":  &config.ServerPort,
		"SET_MAX_IDLE": &config.SetMaxIdleConns,
		"SET_MAX_OPEN": &config.SetMaxOpenConns,
		"SET_MAX_LIFE": &config.SetMaxLifeTime,
	}

	for key, ptr := range envVars {
		value := os.Getenv(key)
		if value == "" {
			log.Warnf("Missing environment variable: %s", key)
		}
		*ptr = value
	}

	return config, nil
}
