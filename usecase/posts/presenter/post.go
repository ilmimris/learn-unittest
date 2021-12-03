package presenter

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
	"github.com/ilmimris/learn-unittest/domain/service"
)

type PostPresenter interface {
	ResponsePost(ctx context.Context, post *model.Post) (*service.Post, error)
	ResponsePosts(ctx context.Context, posts []*model.Post) ([]*service.Post, error)
}
