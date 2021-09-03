package router

import (
	"github.com/gofiber/fiber"
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
	"github.com/hesahesa/materi-rakamin-e2e-crud/service"
)

type BookRouter struct {
	BookService service.BookService
}

func(r *BookRouter) GetBooks(c *fiber.Ctx) {
	books, err := r.BookService.GetAllBooks()

	if err != nil {
		c.Status(403).Send(err)
		return
	}

	c.JSON(books)
}

func(r *BookRouter) InsertBook(c *fiber.Ctx) {
	book := &model.Book{}

	if err := c.BodyParser(book); err != nil {
		c.Status(403).Send(err)
		return
	}

	r.BookService.InsertBook(*book)
}