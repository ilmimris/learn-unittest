package controller

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/service"
	"github.com/ilmimris/learn-unittest/usecase/posts/interactor"
)

type PostController interface {
	CreatePost(ctx context.Context, title, content string, authorID int64) (*service.Post, error)
	GetPosts(ctx context.Context) ([]*service.Post, error)
	GetPost(ctx context.Context, id int64) (*service.Post, error)
}

type PostControllerImpl struct {
	postInteractor interactor.PostInteractor
}

func NewPostController(postInteractor interactor.PostInteractor) PostController {
	return &PostControllerImpl{postInteractor: postInteractor}
}

func (c *PostControllerImpl) GetPosts(ctx context.Context) ([]*service.Post, error) {
	return c.postInteractor.GetPosts(ctx)
}

func (c *PostControllerImpl) GetPost(ctx context.Context, id int64) (*service.Post, error) {
	return c.postInteractor.GetPost(ctx, id)
}

func (c *PostControllerImpl) CreatePost(ctx context.Context, title, content string, authorID int64) (*service.Post, error) {
	return c.postInteractor.CreatePost(ctx, title, content, authorID)
}
