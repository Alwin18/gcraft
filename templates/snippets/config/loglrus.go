package config

import (
	"github.com/sirupsen/logrus"
)

func NewLogger(cfg *Config) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
