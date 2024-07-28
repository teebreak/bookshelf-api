package services

import (
	"bookshelf-api/models"
	"bookshelf-api/repositories"
)

type BookService interface {
	CreateBook(book *models.Book) error
	GetBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.bookRepo.Create(book)
}

func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	return s.bookRepo.FindByID(id)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	return s.bookRepo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	return s.bookRepo.Delete(id)
}
