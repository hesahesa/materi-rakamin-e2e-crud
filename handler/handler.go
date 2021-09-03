package handler

import (
	"github.com/gofiber/fiber"
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
	"github.com/hesahesa/materi-rakamin-e2e-crud/service"
)

type BookHandler struct {
	BookService service.BookService
}

func(r *BookHandler) GetBooks(c *fiber.Ctx) {
	books, err := r.BookService.GetAllBooks()

	if err != nil {
		c.Status(403).Send(err)
		return
	}

	c.JSON(books)
}

func(r *BookHandler) InsertBook(c *fiber.Ctx) {
	book := &model.Book{}

	if err := c.BodyParser(book); err != nil {
		c.Status(403).Send(err)
		return
	}

	r.BookService.InsertBook(*book)
}