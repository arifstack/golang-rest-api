package repository

import "myapp/internal/domain"

type BookRepository interface {
	GetAll() ([]domain.Book, error)
	GetByID(id string) (*domain.Book, error)
	Create(book *domain.Book) error
}
