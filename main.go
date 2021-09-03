package main

import (
	"github.com/gofiber/fiber"
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
	"github.com/hesahesa/materi-rakamin-e2e-crud/repository"
	"github.com/hesahesa/materi-rakamin-e2e-crud/router"
	"github.com/hesahesa/materi-rakamin-e2e-crud/service"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
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

	gormDB, err := gorm.Open(mysql.Open("root:rakamin@tcp(127.0.0.1:3306)/rakamin?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		panic(err.Error())
	}

	model.MigrateBook(gormDB)

	bookRepo := repository.BookRepository{
		DB: gormDB,
	}
	bookService := service.BookService{
		BookRepo: bookRepo,
	}
	bookRouter := router.BookRouter{
		BookService: bookService,
	}

	app.Get("/book", bookRouter.GetBooks)
	app.Post("/book", bookRouter.InsertBook)

	app.Listen(3000)
}
