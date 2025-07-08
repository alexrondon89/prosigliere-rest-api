package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"runtime/debug"
)

// Recover from a panic for each request
func Recover() fiber.Handler {
	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[panic-request] %v\n%s", r, debug.Stack())
				_ = c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{"error": "internal server error"},
				)
			}
		}()
		return c.Next()
	}
}
