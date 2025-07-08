package middleware

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

func ValidateNewPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.Post
		if err := c.BodyParser(&input); err != nil {
			log.Printf("[Middleware-ValidateNewPost] error: %v", err)
			return fiber.NewError(fiber.StatusBadRequest, "body not valid JSON")
		}

		if input.Title == "" {
			return fiber.NewError(fiber.StatusBadRequest, "title is required")
		}

		if input.Content == "" {
			return fiber.NewError(fiber.StatusBadRequest, "content is required")
		}

		c.Locals("postInput", &input)
		return c.Next()
	}
}

func ValidatePostIdFormat() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if _, err := uuid.Parse(id); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "invalid format for post id")
		}
		c.Locals("postId", id)
		return c.Next()
	}
}
