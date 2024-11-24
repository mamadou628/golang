package main

import "fmt"

// ----- Définition du modèle Livre -----
type Livre struct {
	ID     int
	Title  string
	Author string
	Year   int
}

// ----- Définition du Storage (mémorisation des livres) -----
type Storage struct {
	books []Livre
}

// GetBooks retourne tous les livres dans le stockage
func (s *Storage) GetBooks() []Livre {
	return s.books
}

// AddBook ajoute un livre au stockage
func (s *Storage) AddBook(book Livre) {
	s.books = append(s.books, book)
}

// FindBook recherche un livre par son ID
func (s *Storage) FindBook(id int) Livre {
	for _, book := range s.books {
		if book.ID == id {
			return book
		}
	}
	return Livre{} // Retourne un livre vide si non trouvé
}

// ----- Définition du Service de bibliothèque -----
type LibraryService struct {
	Storage *Storage
}

// NewLibraryService crée une nouvelle instance du service de bibliothèque
func NewLibraryService(storage *Storage) *LibraryService {
	return &LibraryService{Storage: storage}
}

// AddBook ajoute un livre à la bibliothèque
func (s *LibraryService) AddBook(title, author string, year int) {
	book := Livre{
		ID:     len(s.Storage.GetBooks()) + 1, // Génère un ID unique
		Title:  title,
		Author: author,
		Year:   year,
	}
	s.Storage.AddBook(book)
}

// ListBooks liste tous les livres dans la bibliothèque
func (s *LibraryService) ListBooks() []Livre {
	return s.Storage.GetBooks()
}

// FindBook recherche un livre par son ID
func (s *LibraryService) FindBook(id int) Livre {
	return s.Storage.FindBook(id)
}

// ----- Fonction principale (main) -----
func main() {
	// Création du stockage en mémoire
	memoryStorage := &Storage{}

	// Création du service de bibliothèque
	libraryService := NewLibraryService(memoryStorage)

	// Ajouter des livres à la bibliothèque
	libraryService.AddBook("Le Go en un coup d'œil", "Jean Dupont", 2020)
	libraryService.AddBook("Programmation Go pour les nuls", "Alice Dupont", 2021)

	// Lister les livres
	books := libraryService.ListBooks()
	fmt.Println("Liste des livres :")
	for _, book := range books {
		fmt.Printf("ID: %d, Titre: %s, Auteur: %s, Année: %d\n", book.ID, book.Title, book.Author, book.Year)
	}

	// Recherche d'un livre par ID
	fmt.Println("\nRecherche du livre avec l'ID 1 :")
	book := libraryService.FindBook(1)
	if (book == Livre{}) {
		fmt.Println("Livre non trouvé.")
	} else {
		fmt.Printf("Livre trouvé : ID: %d, Titre: %s, Auteur: %s, Année: %d\n", book.ID, book.Title, book.Author, book.Year)
	}
}
