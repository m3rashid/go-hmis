package auth

import (
	"fmt"
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
	user := CreateAccountBody{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body",
			"error":   err.Error(),
		})
	}

	// validate the user
	if err := user.Validate(); err != nil {
		// fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid body",
			"error":   err.Error(),
		})
	}
	fmt.Print("hello")

	// fmt.Println(user)

	return c.SendStatus(200)
}
