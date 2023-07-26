package config

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	AppName     string = "go-hmis"
	AppFullName string = "Hospital Management and Informatics System"
)

type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
}

var FiberApp = fiber.New(fiber.Config{
	BodyLimit:             1024 * 1024 * 10, // 10 MB
	ServerHeader:          AppName,
	ETag:                  true,
	DisableStartupMessage: true,
	PassLocalsToViews:     true,
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		/*
			code := fiber.StatusInternalServerError
			Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
		*/
		fmt.Println(err.Error())
		return ctx.SendStatus(fiber.StatusInternalServerError)
	},
})
