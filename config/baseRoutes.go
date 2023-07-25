package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func BaseRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello From " + AppFullName)
	})

	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/metrics", func(c *fiber.Ctx) error {
		/*
			 * TODO: read from user if it needs
			 	- With/Without Header
				- As an API or as a Webpage
		*/

		return monitor.New(monitor.Config{
			Title:      AppName + " " + "Server Metrics",
			CustomHead: AppName,
		})(c)
	})
}
