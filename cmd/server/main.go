package main

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/dal/postgres"
	"github.com/alexrondon89/prosigliere-rest-api/internal/handler"
	"github.com/alexrondon89/prosigliere-rest-api/internal/middleware"
	"github.com/alexrondon89/prosigliere-rest-api/internal/service"
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

	// root ctx
	ctx := context.Background()
	// postgres storage
	pg := postgres.NewPostgresRepo(ctx, "postgres://postgres:postgres@db:5432/prosigliere?sslmode=disable")
	//services
	postSrv := service.NewPostService(pg)
	commentSrv := service.NewCommentService(pg)
	// handler
	h := handler.NewHandler(postSrv, commentSrv)
	// server
	app := fiber.New()
	//global middlewares
	app.Use(middleware.Recover())
	// group
	group := app.Group("/api/posts")
	// post routes
	group.Get("/", h.GetAllPosts)
	group.Post("/", middleware.ValidateNewPost(), h.CreateBlogPost)
	group.Get("/:id", middleware.ValidatePostIdFormat(), h.GetPostById)
	// comment routes
	group.Post("/:id/comments", middleware.ValidatePostIdFormat(), middleware.ValidateNewComment(), h.CreateComment)

	//init server
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
