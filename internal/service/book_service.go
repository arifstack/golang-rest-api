package repository

import (
	"database/sql"
	"go1/internal/domain"
)

type MySQLBookRepository struct {
	db *sql.DB
}

func NewMySQLBookRepository(db *sql.DB) BookRepository {
	return &MySQLBookRepository{db: db}
}

func (r *MySQLBookRepository) GetAll() ([]domain.Book, error) {
	rows, err := r.db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *MySQLBookRepository) GetByID(id string) (*domain.Book, error) {
	var book domain.Book
	row := r.db.QueryRow("SELECT id, title, author, year FROM books WHERE id = ?", id)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func (r *MySQLBookRepository) Create(book *domain.Book) error {
	statement, err := r.db.Prepare("INSERT INTO books (id, title, author, year) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(book.ID, book.Title, book.Author, book.Year)
	return err
}
