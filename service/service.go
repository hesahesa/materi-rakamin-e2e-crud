package service

import (
	"github.com/hesahesa/materi-rakamin-e2e-crud/model"
	"github.com/hesahesa/materi-rakamin-e2e-crud/repository"
)

type BookService struct {
	BookRepo repository.BookRepository
	Name string
}

func(s *BookService) GetAllBooks() ([]model.Book, error) {
	books, err := s.BookRepo.FindAll()

	return books, err
}

func(s *BookService) InsertBook(book model.Book) error {
	_, err := s.BookRepo.Save(book)

	return err
}