package main

import (
	_ "spotifyAPI/app/docs"
	"spotifyAPI/config"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
)

func init() {
	config.InitLogger()
}

func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", HelloWorld)
	app.Listen(":5000")
}

// @Summary init
// @Description init
// @Tags init

// @Router / [get]
func HelloWorld(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}
