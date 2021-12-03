package repository

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
)

type PostRepository interface {
	FetchPost(ctx context.Context, id int64) (*model.Post, error)
	FetchPosts(ctx context.Context) ([]*model.Post, error)
	CreatePost(ctx context.Context, post *model.Post) (*model.Post, error)
	UpdatePost(ctx context.Context, post *model.Post) error
	DeletePost(ctx context.Context, id int64) error
}
