package config

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator(cfg *Config) *validator.Validate {
	return validator.New()
}
