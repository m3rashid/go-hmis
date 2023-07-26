package auth

import (
	"fmt"
	"go-hmis/helpers"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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
	user := UserModel{}
	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := LoginValidator(&user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	db := helpers.GetLocal[*mongo.Client](c, "db")
	res, err := db.Database(DbName).Collection(UserModelName).InsertOne(c.Context(), user)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(DbName)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	fmt.Println(res)

	return c.SendStatus(200)
}
