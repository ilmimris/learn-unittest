package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/learn-unittest/internal/interfaces"
	"github.com/ilmimris/learn-unittest/transport/rest/handler"
)

func InitPost(rest interfaces.Rest, r fiber.Router) {
	postGroup := r.Group("/posts")
	postGroup.Get("/", handler.GetPosts(rest))
}
