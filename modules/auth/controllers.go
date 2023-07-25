package auth

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
	return c.SendStatus(200)
}

func LoginHandler(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func CreateAccountHandler(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(200)
}
