package config

import "github.com/gofiber/fiber/v2"

func NewFiber(cfg *Config) *fiber.App {
	return fiber.New(fiber.Config{
		Network: fiber.NetworkTCP,
	})
}
