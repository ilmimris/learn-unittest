package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/learn-unittest/internal/interfaces"
)

func InitRoot(rest interfaces.Rest) fiber.Router {
	router := rest.GetRouter()

	return router.Group("/")
}
