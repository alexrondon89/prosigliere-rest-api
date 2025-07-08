package middleware

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/gofiber/fiber/v2"
)

func ValidateNewComment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input model.Comment
		if err := c.BodyParser(&input); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "body not valid JSON")
		}

		if input.PostId == "" {
			return fiber.NewError(fiber.StatusBadRequest, "post id is required")
		}

		if input.Username == "" {
			return fiber.NewError(fiber.StatusBadRequest, "username is required")
		}

		if input.Content == "" {
			return fiber.NewError(fiber.StatusBadRequest, "content is required")
		}

		c.Locals("commentInput", &input)
		return c.Next()
	}
}
