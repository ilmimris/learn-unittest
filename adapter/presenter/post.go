package presenter

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
	"github.com/ilmimris/learn-unittest/domain/service"
	pp "github.com/ilmimris/learn-unittest/usecase/posts/presenter"
)

type PostPresenterImpl struct{}

func NewPostPresenter() pp.PostPresenter {
	return &PostPresenterImpl{}
}

func (p *PostPresenterImpl) ResponsePosts(ctx context.Context, posts []*model.Post) ([]*service.Post, error) {
	var response []*service.Post
	for _, post := range posts {
		response = append(response, &service.Post{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
			Author:  post.Author.Name,
		})
	}
	return response, nil
}

func (p *PostPresenterImpl) ResponsePost(ctx context.Context, post *model.Post) (*service.Post, error) {
	return &service.Post{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
		Author:  post.Author.Name,
	}, nil
}
