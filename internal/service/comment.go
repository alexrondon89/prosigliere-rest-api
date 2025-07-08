package service

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/alexrondon89/prosigliere-rest-api/internal/util"
	"log"
	"time"
)

type commentStorage interface {
	CreateComment(ctx context.Context, input *model.Comment) (*model.Comment, error)
}

type CommentService struct {
	commentStorage commentStorage
}

func NewCommentService(commentStorage commentStorage) *CommentService {
	return &CommentService{
		commentStorage: commentStorage,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	ctx = util.SafeCtx(ctx)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	resp, err := s.commentStorage.CreateComment(ctx, comment)
	if err != nil {
		log.Printf("[Service-CreateComment] create comment in database failed: %v", err)
		return nil, err
	}
	return resp, nil
}
