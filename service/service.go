package service

import (
	"github.com/hesahesa/materi-rakamin-e2e-crud/repository"
)

type BookService struct {
	BookRepo repository.BookRepository
}

//