package services

import (
	"Library/pkg/models"
	"Library/pkg/storage"
	"fmt"
)

type Libraryservice struct {
	storage *storage.MemoryStorage
}

func Newlibraryservice(storage *storage.MemoryStorage) *Libraryservice {
	return &Libraryservice{storage: storage}
}

// AddBook ajoute un livre à la bibliothèque
func (s *Libraryservice) AddBook(title, author string, year int) error {
	books := s.storage.GetBooks()
	for _, book := range books {
		if book.Title == title && book.Author == author {
			return fmt.Errorf("book '%s' by '%s' already exists", title, author)
		}
	}

	book := models.Livre{
		ID:     len(books) + 1, // Génère un ID unique
		Title:  title,
		Author: author,
		Year:   year,
	}
	s.storage.AddBook(book)
	return nil
}

// Listbook retourne la liste des livres
func (s *Libraryservice) ListBook() []models.Livre {
	books := s.storage.GetBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return nil
	}

	for _, book := range books {
		fmt.Println(book)
	}
	return books
}

// Findbook recherche un livre par ID
func (s *Libraryservice) FindBook(id int) (models.Livre, error) {
	books := s.storage.GetBooks()
	for _, book := range books {
		if book.ID == id {
			return book, nil
		}
	}
	return models.Livre{}, fmt.Errorf("book with id %d not found", id)
}

// Deletebook supprime un livre par ID
func (s *Libraryservice) DeleteBook(id int) error {
	books := s.storage.GetBooks()
	for i, book := range books {
		if book.ID == id {
			// Suppression du livre
			s.storage.DeleteBook(i)
			return nil
		}
	}
	return fmt.Errorf("book with id %d not found", id)
}
