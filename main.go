package main

import (
	"github.com/gofiber/fiber"
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
)

func helloWorld(c *fiber.Ctx) {
	book := model.Book{
		Title: "test title",
	}
	c.JSON(book)
}

func helloPost(c *fiber.Ctx) {
	book := &model.Book{}

	if err := c.BodyParser(book); err != nil {
		c.Status(403).Send(err)
		return
	}

	c.JSON(book)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)
	app.Post("/", helloPost)

	app.Listen(3000)
}
