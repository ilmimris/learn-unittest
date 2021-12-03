package repository

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
)

type AuthorRepository interface {
	FetchAuthorByID(ctx context.Context, id int64) (*model.Author, error)
}
