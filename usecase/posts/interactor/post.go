package interactor

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
	"github.com/ilmimris/learn-unittest/domain/service"
	ar "github.com/ilmimris/learn-unittest/usecase/author/repository"
	"github.com/ilmimris/learn-unittest/usecase/posts/presenter"
	"github.com/ilmimris/learn-unittest/usecase/posts/repository"
)

type PostInteractor interface {
	CreatePost(ctx context.Context, title, content string, authorID int64) (*service.Post, error)
	GetPost(ctx context.Context, id int64) (*service.Post, error)
	GetPosts(ctx context.Context) ([]*service.Post, error)
	UpdatePost(ctx context.Context, id int64, title, content string) (*service.Post, error)
	DeletePost(ctx context.Context, id int64) error
}

type PostInteractorImpl struct {
	PostRepository  repository.PostRepository
	AuthorRepostory ar.AuthorRepository
	PostPresenter   presenter.PostPresenter
}

func NewPostInteractor(p presenter.PostPresenter, r repository.PostRepository, a ar.AuthorRepository) PostInteractor {
	return &PostInteractorImpl{
		PostPresenter:   p,
		PostRepository:  r,
		AuthorRepostory: a,
	}
}

func (i *PostInteractorImpl) CreatePost(ctx context.Context, title, content string, authorID int64) (*service.Post, error) {
	var post = &model.Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}

	post, err := i.PostRepository.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return i.PostPresenter.ResponsePost(ctx, post)
}

func (i *PostInteractorImpl) GetPost(ctx context.Context, id int64) (*service.Post, error) {
	post, err := i.PostRepository.FetchPost(ctx, id)
	if err != nil {
		return nil, err
	}

	author, err := i.AuthorRepostory.FetchAuthorByID(ctx, post.AuthorID)
	if err != nil {
		return nil, err
	}

	post.Author = author
	return i.PostPresenter.ResponsePost(ctx, post)
}

func (i *PostInteractorImpl) GetPosts(ctx context.Context) ([]*service.Post, error) {
	posts, err := i.PostRepository.FetchPosts(ctx)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		author, err := i.AuthorRepostory.FetchAuthorByID(ctx, post.AuthorID)
		if err != nil {
			return nil, err
		}
		post.Author = author
	}

	return i.PostPresenter.ResponsePosts(ctx, posts)
}

func (i *PostInteractorImpl) UpdatePost(ctx context.Context, id int64, title, content string) (*service.Post, error) {
	var post = &model.Post{
		ID:      id,
		Title:   title,
		Content: content,
	}

	err := i.PostRepository.UpdatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return i.PostPresenter.ResponsePost(ctx, post)
}

func (i *PostInteractorImpl) DeletePost(ctx context.Context, id int64) error {
	return i.PostRepository.DeletePost(ctx, id)
}
