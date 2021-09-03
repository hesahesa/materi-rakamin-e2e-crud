package router

import (
	"github.com/gofiber/fiber"
	"github.com/hesahesa/materi-rakamin-e2e-crud/service"
)

type BookRouter struct {
	BookService service.BookService
}

func(r *BookRouter) GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}