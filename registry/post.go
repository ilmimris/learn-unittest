package registry

import (
	"github.com/ilmimris/learn-unittest/adapter/controller"
	pp "github.com/ilmimris/learn-unittest/adapter/presenter"
	pr "github.com/ilmimris/learn-unittest/adapter/repository"
	"github.com/ilmimris/learn-unittest/usecase/posts/interactor"
	"github.com/ilmimris/learn-unittest/usecase/posts/presenter"
	"github.com/ilmimris/learn-unittest/usecase/posts/repository"
)

func (r *registry) NewPostController() controller.PostController {
	return controller.NewPostController(r.NewPostInteractor())
}

func (r *registry) NewPostRepository() repository.PostRepository {
	return pr.NewPostRepository()
}

func (r *registry) NewPostPresenter() presenter.PostPresenter {
	return pp.NewPostPresenter()
}

func (r *registry) NewPostInteractor() interactor.PostInteractor {
	return interactor.NewPostInteractor(r.NewPostPresenter(), r.NewPostRepository(), r.NewAuthorRepository())
}
