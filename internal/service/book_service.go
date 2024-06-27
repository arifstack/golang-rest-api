package service

import (
	"golang-rest-api/internal/domain"
	"golang-rest-api/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) GetBookByID(id int) (*domain.Book, error) {
	// Implement business logic here, e.g., validate input, call repository, etc.
	return s.repo.FindByID(id)
}
