package rest

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/learn-unittest/adapter/controller"
	"github.com/ilmimris/learn-unittest/registry"
	"github.com/ilmimris/learn-unittest/transport/rest/group"
)

type RestImpl struct {
	port          int
	router        *fiber.App
	appController controller.App
	listenErrCh   chan error
}

func NewRest() *RestImpl {

	app := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		app.Shutdown()
	}()

	registry := registry.NewRegistry()
	appController := registry.NewAppController()

	r := &RestImpl{
		port:          8080,
		appController: appController,
		router:        app,
	}

	root := group.InitRoot(r)
	group.InitPost(r, root)

	return r
}

func (r *RestImpl) Serve() {
	if err := r.router.Listen(fmt.Sprintf(":%d", r.port)); err != nil {
		panic(err)
	}
}

func (r *RestImpl) ListernError() <-chan error {
	return r.listenErrCh
}

func (r *RestImpl) SignalCheck() {
	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Exiting gracefully...")
	case err := <-r.ListernError():
		log.Panic("Error starting web server, exiting gracefully:", err)
	}
}

func (r *RestImpl) GetRouter() *fiber.App {
	return r.router
}

func (r *RestImpl) GetAppController() *controller.App {
	return &r.appController
}
