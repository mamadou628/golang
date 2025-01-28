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
	reader := bufio.NewReader(os.Stdin)

	// Demander le chemin du fichier à l'utilisateur
	fmt.Print("Entrez le chemin complet du fichier : ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	var matrice [][]float64
	var numColumns int

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Fichier non trouvé :", filePath)
		return
	}

	// Lire la matrice depuis le fichier
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		if line == "" {
			continue
		}
		values := strings.Fields(line)
		var row []float64
		for _, val := range values {
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Printf("Erreur de conversion à la ligne %d : %v\n", lineNumber, err)
				return
			}
			row = append(row, num)
		}

		if numColumns == 0 {
			numColumns = len(row)
		} else if len(row) != numColumns {
			fmt.Printf("Erreur : Ligne %d a %d colonnes, mais une ligne précédente a %d colonnes.\n",
				lineNumber, len(row), numColumns)
			return
		}
		matrice = append(matrice, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	m := len(matrice)
	n := len(matrice[0])

	fmt.Println("Matrice lue depuis le fichier :")
	for _, val := range matrice {
		fmt.Println(val)
	}

	// Vérification du type de système
	if m < n-1 {
		fmt.Println("Système sous-déterminé : il y a plus d'inconnues que d'équations.")
	} else if m > n-1 {
		fmt.Println("Système surdéterminé : il y a plus d'équations que d'inconnues.")
	}

	// Appliquer l'élimination de Gauss
	gaussElimination(matrice, m, n)

	fmt.Println("Matrice après élimination de Gauss :")
	for _, val := range matrice {
		fmt.Println(val)
	}

	if m+1 == n {
		sol := solution(matrice, m)
		fmt.Println("\nSolution unique :")
		for i, s := range sol {
			fmt.Printf("x%d = %.2f\n", i+1, s)
		}
	} else if m < n-1 {
		fmt.Println("\nSolution générale :")
		generalSolution(matrice, m, n)
	} else {
		fmt.Println("Le système est incompatible ou possède une infinité de solutions.")
	}
}

// Élimination de Gauss avec gestion des pivots nuls
func gaussElimination(matrice [][]float64, m, n int) {
	for i := 0; i < m; i++ {
		// Trouver le plus grand pivot dans la colonne
		maxRow := i
		for k := i + 1; k < m; k++ {
			if math.Abs(matrice[k][i]) > math.Abs(matrice[maxRow][i]) {
				maxRow = k
			}
		}

		// Permuter les lignes si nécessaire
		if maxRow != i {
			matrice[i], matrice[maxRow] = matrice[maxRow], matrice[i]
		}

		// Vérifier si le pivot est nul
		if math.Abs(matrice[i][i]) < 1e-9 {
			fmt.Println("Pivot nul détecté, le système peut être singulier ou avoir des solutions infinies.")
			continue
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

// Résolution du système triangulaire supérieur
func solution(matrice [][]float64, m int) []float64 {
	sol := make([]float64, m)
	for i := m - 1; i >= 0; i-- {
		sol[i] = matrice[i][m]
		for j := i + 1; j < m; j++ {
			sol[i] -= matrice[i][j] * sol[j]
		}
	}
	return sol
}

// Afficher une solution générale pour un système sous-déterminé
func generalSolution(matrice [][]float64, m, n int) {
	vars := make([]string, n-1)
	for i := range vars {
		vars[i] = fmt.Sprintf("x%d", i+1)
	}

	fmt.Println("the solve general :")
	for i := 0; i < m; i++ {
		fmt.Printf("%s = %.2f", vars[i], matrice[i][n-1])
		for j := i + 1; j < n-1; j++ {
			fmt.Printf(" - %.2f*%s", matrice[i][j], vars[j])
		}
		fmt.Println()
	}

	// Les variables libres
	for i := m; i < n-1; i++ {
		fmt.Printf("%s is free.\n", vars[i])
	}
}
