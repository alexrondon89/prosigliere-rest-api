package service

import (
	"context"
	"github.com/alexrondon89/prosigliere-rest-api/internal/model"
	"github.com/alexrondon89/prosigliere-rest-api/internal/util"
	"log"
	"time"
)

type postStorage interface {
	CreatePostInDb(ctx context.Context, input *model.Post) (*model.Post, error)
	GetPostFromDb(ctx context.Context, id string) (*model.Post, error)
	GetAllPostsFromDb(ctx context.Context) ([]model.Post, error)
}

type PostService struct {
	postStorage postStorage
}

func NewPostService(postStorage postStorage) *PostService {
	return &PostService{
		postStorage: postStorage,
	}
}

func (ps *PostService) CreatePost(ctx context.Context, input *model.Post) (*model.Post, error) {
	ctx = util.SafeCtx(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := ps.postStorage.CreatePostInDb(ctx, input)
	if err != nil {
		log.Printf("[Service-CreatePost] create post in database failed: %v", err)
		return nil, err
	}
	return resp, nil
}

func (ps *PostService) GetPost(ctx context.Context, id string) (*model.Post, error) {
	ctx = util.SafeCtx(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := ps.postStorage.GetPostFromDb(ctx, id)
	if err != nil {
		log.Printf("[Service-GetPost] get post from database failed: %v", err)
		return nil, err
	}
	return resp, nil
}

func (ps *PostService) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	ctx = util.SafeCtx(ctx)
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	resp, err := ps.postStorage.GetAllPostsFromDb(ctx)
	if err != nil {
		log.Printf("[Service-GetAllPosts] get all posts from database failed: %v", err)
	}
	return resp, err
}
