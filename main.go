package main

import (
	"go-hmis/config"
	"go-hmis/helpers"
	"go-hmis/modules/auth"
	"log"
	"os"
)

func main() {
	helpers.LoadEnv()

	app := config.FiberApp
	config.SetupMiddlewares(app)
	config.BaseRoutes(app)

	auth.SetupAuthRoutes(app)

	// Setup other routes

	log.Fatal(app.Listen(":" + os.Getenv("HTTP_PORT")))
}
