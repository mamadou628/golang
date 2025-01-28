package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Exemple de matrice augmentée avec m lignes et n colonnes (n = m+1 pour les systèmes carrés)
	reader := bufio.NewReader(os.Stdin)

	// Demander le chemin du fichier à l'utilisateur
	fmt.Print("Entrez le chemin complet du fichier : ")
	filePath, _ := reader.ReadString('\n')

	// Supprimer les espaces ou sauts de ligne en trop
	filePath = strings.TrimSpace(filePath)
	var matrice [][]float64
	var numColumns int

	// Vérifier si le fichier existe
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Fichier non trouvé :", filePath)
	} else {
		fmt.Println("Fichier trouvé :", filePath)

		// Ouvrir le fichier
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Erreur lors de l'ouverture du fichier :", err)
			return
		}
		defer file.Close()

		// Lire le fichier ligne par ligne
		scanner := bufio.NewScanner(file)
		lineNumber := 0
		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			if line == "" {
				continue
			}
			// Diviser la ligne en valeurs séparées par des espaces
			values := strings.Fields(line)
			var row []float64

			// Convertir chaque valeur en float64 et ajouter à la ligne
			for _, val := range values {
				num, err := strconv.ParseFloat(val, 64)
				if err != nil {
					fmt.Printf("Erreur de conversion à la ligne %d : %v\n", lineNumber, err)
					return
				}
				row = append(row, num)
			}

			// Vérifier que toutes les lignes ont le même nombre de colonnes
			if numColumns == 0 {
				numColumns = len(row) // Initialiser le nombre de colonnes
			} else if len(row) != numColumns {
				fmt.Printf("Erreur : Ligne %d a %d colonnes, mais une ligne précédente a %d colonnes.\n",
					lineNumber, len(row), numColumns)
				return
			}

			// Ajouter la ligne à la matrice
			matrice = append(matrice, row)
		}

		// Vérifier les erreurs de lecture
		if err := scanner.Err(); err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}

		// Afficher la matrice
		fmt.Println("Matrice lue depuis le fichier :")
		for _, val := range matrice {
			fmt.Println(val)
		}
	}

	m := len(matrice)    // Nombre de lignes
	n := len(matrice[0]) // Nombre de colonnes

	// Affichage du nombre de lignes et de colonnes
	fmt.Println("Le nombre de lignes est de", m, "et le nombre de colonnes est de", n)

	// Affichage de la matrice avec la bonne dimension
	for i := 0; i < m; i++ { // Boucle sur les lignes
		for j := 0; j < n; j++ { // Boucle sur les colonnes
			fmt.Printf("%8.2f ", matrice[i][j]) // Affichage avec 2 décimales
		}
		fmt.Println() // Nouvelle ligne après chaque ligne de la matrice
	}
	gaussElimination(matrice, m, n)

	fmt.Println("la matrice apres elemination de gauss")
	for _, val := range matrice {
		fmt.Println(val)
	}

	solution(matrice, m)
	if m+1 == n {
		solution := solution(matrice, m)
		fmt.Println("\nSolution :")
		for i, sol := range solution {
			fmt.Printf("x%d = %.2f\n", i+1, sol)
		}
	} else {
		fmt.Println("Impossible de résoudre directement : matrice non carrée.")
	}
}

func gaussElimination(matrice [][]float64, m, n int) {
	for i := 0; i < m; i++ {
		// Trouver le pivot
		if math.Abs(matrice[i][i]) < 1e-9 {
			fmt.Println("Pivot nul détecté, recherche d'un pivot dans une autre ligne.")
			return
		}

		// Diviser la ligne i par le pivot
		pivot := matrice[i][i]
		for j := 0; j < n; j++ {
			matrice[i][j] /= pivot
		}

		// Annuler les éléments sous le pivot
		for k := i + 1; k < m; k++ {
			factor := matrice[k][i]
			for j := i; j < n; j++ {
				matrice[k][j] -= factor * matrice[i][j]
			}
		}
	}
}

func solution(matrice [][]float64, m int) []float64 {
	solution := make([]float64, m)

	// Partir de la dernière ligne vers la première
	for i := m - 1; i >= 0; i-- {
		solution[i] = matrice[i][m] // Dernière colonne
		for j := i + 1; j < m; j++ {
			solution[i] -= matrice[i][j] * solution[j]
		}
	}

	return solution
}
