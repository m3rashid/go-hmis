package auth

import (
	"go-hmis/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	user := db.CreateUserParams{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	user.Password = string(password)

	return c.SendStatus(200)
}
