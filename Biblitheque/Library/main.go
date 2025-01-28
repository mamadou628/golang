package main

import (
	"Library/pkg/models"
	"Library/pkg/services"
	"Library/pkg/storage"
	"fmt"
)

func main() {
	// Initialiser le stockage et le service
	memoryStorage := storage.NewMemoryStorage()
	library := services.Newlibraryservice(memoryStorage)
	// ajouter des livre
	booksToAdd := []models.Livre{
		{Title: "Le Petit Prince", Author: "Antoine de Saint-Exupéry", Year: 1943},
		{Title: "1984", Author: "George Orwell", Year: 1949},
		{Title: "L'Étranger", Author: "Albert Camus", Year: 1942},
		{Title: "Les Misérables", Author: "Victor Hugo", Year: 1862},
		{Title: "Moby Dick", Author: "Herman Melville", Year: 1851},
		{Title: "Don Quichotte", Author: "Miguel de Cervantes", Year: 1605},
		// Ajoute autant de livres que tu veux ici
	}

	// Ajout des livres à la bibliothèque
	for _, book := range booksToAdd {
		err := library.AddBook(book.Title, book.Author, book.Year)
		if err != nil {
			fmt.Println("Erreur lors de l'ajout du livre:", err)
		}
	}
	// Lister les livres
	books := library.ListBook()
	fmt.Println("Liste des livres :")
	for _, book := range books {
		fmt.Printf("ID: %d, Titre: %s, Auteur: %s, Année: %d\n", book.ID, book.Title, book.Author, book.Year)
	}

	// Trouver un livre
	fmt.Println("\nRecherche d'un livre par ID :")
	book, err := library.FindBook(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Trouvé : ID: %d, Titre: %s, Auteur: %s, Année: %d\n", book.ID, book.Title, book.Author, book.Year)
	}

	// Supprimer un livre
	fmt.Println("\nSuppression d'un livre :")
	err = library.DeleteBook(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Livre supprimé avec succès !")
	}

	// Afficher la liste mise à jour
	fmt.Println("\nListe des livres après suppression :")
	books = library.ListBook()
	for _, book := range books {
		fmt.Printf("ID: %d, Titre: %s, Auteur: %s, Année: %d\n", book.ID, book.Title, book.Author, book.Year)
	}
}
