package main

import (
	"be-nuxt-fiber/config"
	"be-nuxt-fiber/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	config := config.Config{}
	config.InitConfig()

	repositories.ConnectToMongoDB(&config.Mongo)

	app.Listen(":3000")
}
