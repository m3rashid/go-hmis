package main

import (
	"hmis/config"
	"hmis/helpers"
	"log"
	"os"
)

func main() {
	helpers.LoadEnv()

	app := config.FiberApp
	config.SetupMiddlewares(app)
	config.BaseRoutes(app)

	// Setup other routes

	log.Fatal(app.Listen(":" + os.Getenv("HTTP_PORT")))
}
