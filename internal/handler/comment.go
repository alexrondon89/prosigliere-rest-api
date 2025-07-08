package handler

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (p *Handler) CreateComment(c *fiber.Ctx) error {
	comment, ok := c.Locals("commentInput").(*model.Comment)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "missing comment input")
	}
	postId, ok := c.Locals("postId").(string)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "missing post id input")
	}
	comment.PostId = postId
	_, err := p.commentService.CreateComment(c.Context(), comment)
	if err != nil {
		log.Printf("[Handler-CreateComment] request to create new comment failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(comment)
}
