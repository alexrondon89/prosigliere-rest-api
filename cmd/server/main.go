package main

import (
	"github.com/alexrondon89/prosigliere-rest-api/internal/handler"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"runtime/debug"
)

func main() {
	// to catch root panic
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[panic-main] %v\n%s", err, debug.Stack())
			os.Exit(1)
		}
	}()

	// handler
	h := handler.NewHandler()
	app := fiber.New()
	// group
	group := app.Group("/api/posts")
	// post routes
	group.Get("/", h.GetAllPosts)
	group.Post("/", h.CreateBlogPost)
	group.Get("/:id", h.GetPostById)
	// comment routes
	group.Post("/:id/comments", h.CreateComment)

	//init server
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
