package main

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type apiHandler struct {
	store  Store
	addr   string
	engine *html.Engine
}

func NewApiHandler(store Store, addr string) *apiHandler {
	engine := html.New("./views", ".html")

	engine.Reload(true)

	return &apiHandler{
		store:  store,
		addr:   addr,
		engine: engine,
	}
}

func (ah *apiHandler) Run() {
	app := fiber.New(fiber.Config{
		Views: ah.engine,
	})

	app.Get("/", ah.handleHome)
	app.Get("/:name", ah.handleGetPaste)
	app.Post("/pastes", ah.handleSavePaste)
	app.Delete("/:name", ah.handleDeletePaste)

	log.Print("api running on port: ", ah.addr)
	app.Listen(ah.addr)
}

func (ah *apiHandler) handleHome(c *fiber.Ctx) error {
	return c.Render("home", nil)
}

func (ah *apiHandler) handleGetPaste(c *fiber.Ctx) error {
	name := c.Params("name")

	paste, err := ah.store.RetrievePaste(name)
	if err != nil {
		return err
	}

	if paste == nil {
		return errors.New("couldn't retrieve a paste ;c")
	}

	return c.Render("paste", fiber.Map{"paste": paste, "name": name})
}

func (ah *apiHandler) handleSavePaste(c *fiber.Ctx) error {
	var paste Paste
	err := c.BodyParser(&paste)
	if err != nil {
		return err
	}

	result, _ := ah.store.RetrievePaste(paste.Name)
	if len(result.(string)) > 0 {
		return errors.New("paste already exists")
	}

	err = ah.store.SavePaste(paste)
	if err != nil {
		return err
	}

	return c.Redirect("/")
}

func (ah *apiHandler) handleDeletePaste(c *fiber.Ctx) error {
	name := c.Params("name")
	err := ah.store.DeletePaste(name)
	if err != nil {
		return err
	}
	return c.SendStatus(204)
}
