package registry

import (
	ar "github.com/ilmimris/learn-unittest/adapter/repository"
	"github.com/ilmimris/learn-unittest/usecase/author/repository"
)

func (r *registry) NewAuthorRepository() repository.AuthorRepository {
	return ar.NewAuthorRepository()
}
