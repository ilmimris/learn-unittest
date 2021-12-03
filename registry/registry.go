package registry

import "github.com/ilmimris/learn-unittest/adapter/controller"

type Registry interface {
	NewAppController() controller.App
}

type registry struct{}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.App {
	return controller.App{
		Posts: r.NewPostController(),
	}
}
