package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/shareaf/store"
)

func main() {
	app := fiber.New()
	store := store.NewRedisStore()

	store.Set()
	store.Get()
	store.Get()
	store.Get()
	store.Get()

	app.Get("/", handleHome)

	log.Print("runnning server on port 3000")
	app.Listen(":3000")
}

func handleHome(c *fiber.Ctx) error {
	return c.Render("index.html", fiber.Map{
		"hello": "world",
	})
}
