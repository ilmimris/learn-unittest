package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilmimris/learn-unittest/internal/interfaces"
)

func GetPosts(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		app := rest.GetAppController()
		posts, err := app.Posts.GetPosts(c.Context())
		if err != nil {
			return c.Status(500).JSON(err)
		}

		return c.JSON(posts)
	}
}
