package handler

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (p *Handler) CreateComment(c *fiber.Ctx) error {
	comment := c.Locals("commentInput").(*model.Comment)
	resp, err := p.commentService.CreateComment(c.Context(), comment)
	if err != nil {
		log.Printf("[Handler-CreateComment] request to create new comment failed: %v", err)
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}
