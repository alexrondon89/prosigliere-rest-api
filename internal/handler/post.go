package handler

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (p *Handler) GetAllPosts(c *fiber.Ctx) error {
	resp, err := p.postService.GetAllPosts(c.Context())
	if err != nil {
		log.Printf("[Handler-GetAllPosts] request to get all posts failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (p *Handler) CreateBlogPost(c *fiber.Ctx) error {
	input := c.Locals("postInput").(*model.Post)
	resp, err := p.postService.CreatePost(c.Context(), input)
	if err != nil {
		log.Printf("[Handler-CreateBlogPost] request to create post failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (p *Handler) GetPostById(c *fiber.Ctx) error {
	id := c.Locals("postId").(string)
	resp, err := p.postService.GetPost(c.Context(), id)
	if err != nil {
		log.Printf("[Handler-GetPostById] request to get post failed: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}
