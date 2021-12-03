package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/learn-unittest/adapter/controller"
)

type Rest interface {
	GetRouter() *fiber.App
	GetAppController() *controller.App
}
