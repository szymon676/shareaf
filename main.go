package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// store := store.NewRedisStore()

	app.Get("/", handleHome)

	log.Print("runnning server on port 3000")
	app.Listen(":3000")
}

func handleHome(c *fiber.Ctx) error {
	return c.Render("index.html", fiber.Map{
		"wow": "some text 123...",
	})
}
