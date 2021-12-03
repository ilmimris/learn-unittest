package repository

import (
	"context"

	"github.com/ilmimris/learn-unittest/domain/model"
	ar "github.com/ilmimris/learn-unittest/usecase/author/repository"
)

type AuthorRepositoryImpl struct{}

func NewAuthorRepository() ar.AuthorRepository {
	return &AuthorRepositoryImpl{}
}

func (ar *AuthorRepositoryImpl) FetchAuthorByID(ctx context.Context, id int64) (*model.Author, error) {
	return &model.Author{
		ID:   id,
		Name: "John Doe",
	}, nil
}
