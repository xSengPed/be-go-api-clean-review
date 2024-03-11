package main

import (
	"be-nuxt-fiber/config"
	postcontroller "be-nuxt-fiber/controller/post"
	"be-nuxt-fiber/repositories"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	config := config.Config{}
	config.InitConfig()

	repositories.ConnectToMongoDB(&config.Mongo)
	v1 := app.Group("/v1")
	api := v1.Group("/api")
	// api.Get("/post")
	api.Post("/post", postcontroller.CreatePost)
	// api.Put("/post")
	// api.Delete("/post")
	app.Listen(":3000")
}
