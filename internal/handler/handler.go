package handler

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
)

type postServiceInterface interface {
	CreatePost(ctx context.Context, input *model.Post) (*model.Post, error)
	GetPost(ctx context.Context, id string) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
}

type commentServiceInterface interface {
	CreateComment(c context.Context, comment *model.Comment) (*model.Comment, error)
}

type Handler struct {
	postService    postServiceInterface
	commentService commentServiceInterface
}

func NewHandler(postService postServiceInterface, commentService commentServiceInterface) *Handler {
	return &Handler{
		postService:    postService,
		commentService: commentService,
	}
}
