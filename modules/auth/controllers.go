package auth

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

var DbName string = os.Getenv("MONGO_DB_NAME")
var DbUrl string = os.Getenv("MONGO_URI")

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

	return c.SendStatus(200)
}
