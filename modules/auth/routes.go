package auth

import "github.com/gofiber/fiber/v2"

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/api/auth/login", LoginHandler)
	app.Post("/api/auth/logout", LogoutHandler)
	app.Post("/api/auth/create-account", CreateAccountHandler)
}
