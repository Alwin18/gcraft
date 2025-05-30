package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the request method is OPTIONS, send a 200 response and end the middleware chain
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}

		// Call the next middleware in the chain
		return c.Next()
	}
}
