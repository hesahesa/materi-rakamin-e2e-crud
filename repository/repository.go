package repository

import (
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func(r *BookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	findResult := r.DB.Find(&books)

	return books, findResult.Error
}

func(r *BookRepository) FindById(id string) (model.Book, error) {
	return model.Book{}, nil
}

func(r *BookRepository) Save(book model.Book) (model.Book, error) {
	trx := r.DB.Create(&book)

	return book, trx.Error
}

func(r *BookRepository) Update(id string, book model.Book) {
	//
}

func(r *BookRepository) Delete(id string) {
	//
}
