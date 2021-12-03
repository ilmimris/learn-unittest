package repository

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
	pr "github.com/ilmimris/learn-unittest/usecase/posts/repository"
)

type PostRepositoryImpl struct{}

func NewPostRepository() pr.PostRepository {
	return &PostRepositoryImpl{}
}

func (pr *PostRepositoryImpl) FetchPosts(ctx context.Context) ([]*model.Post, error) {
	return []*model.Post{
		{
			ID:       1,
			Title:    "Hello 1",
			Content:  "Hello World",
			AuthorID: int64(1),
			Author: &model.Author{
				ID:   int64(1),
				Name: "John Doe",
			},
		},
		{
			ID:       2,
			Title:    "Hello 2",
			Content:  "Hello World",
			AuthorID: int64(1),
			Author: &model.Author{
				ID:   int64(1),
				Name: "John Doe",
			},
		},
	}, nil
}

func (pr *PostRepositoryImpl) FetchPost(ctx context.Context, id int64) (*model.Post, error) {
	return &model.Post{
		ID:       1,
		Title:    "Hello 1",
		Content:  "Hello World",
		AuthorID: int64(1),
	}, nil
}

func (pr *PostRepositoryImpl) CreatePost(ctx context.Context, post *model.Post) (*model.Post, error) {
	panic("implement me!")
}

func (pr *PostRepositoryImpl) UpdatePost(ctx context.Context, post *model.Post) error {
	panic("implement me!")
}

func (pr *PostRepositoryImpl) DeletePost(ctx context.Context, id int64) error {
	panic("implement me!")
}
