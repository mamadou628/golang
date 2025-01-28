package storage

import (
	"Library/pkg/models"
	"errors"
)

type Error interface {
	Error() string
}

type MemoryStorage struct {
	books  []models.Livre
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		books:  []models.Livre{},
		nextID: 1,
	}
}

func (s *MemoryStorage) AddBook(book models.Livre) models.Livre {
	book.ID = s.nextID
	s.nextID++
	s.books = append(s.books, book)
	return book
}

func (s *MemoryStorage) GetBooks() []models.Livre {
	return s.books
}

func (s *MemoryStorage) FindBook(id int) (models.Livre, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Livre{}, errors.New("book not found")
}

func (s *MemoryStorage) DeleteBook(id int) error {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
