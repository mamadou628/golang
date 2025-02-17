package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

/*
	Task représente une tâche avec ses critères.

Creation de la strucure des taches
en golang on utilise les structures a la place des class
*/
type Task struct {
	Name       string  `json:"Name"`
	Urgency    float64 `json:"Urgency"`
	Importance float64 `json:"Importance"`
	Complexity float64 `json:"Complexity"`
	Priority   float64 // Calculé après chargement
}

// PriorityCalculator gère le calcul des priorités avec des pondérations.
type PriorityCalculator struct {
	UrgencyWeight    float64
	ImportanceWeight float64
	ComplexityWeight float64
}

// CalculatePriority calcule la priorité pour une tâche donnée.
// creation de la methode pour le calculatePriority
func (pc PriorityCalculator) CalculatePriority(task *Task) {
	task.Priority = (task.Urgency * pc.UrgencyWeight) +
		(task.Importance * pc.ImportanceWeight) +
		(task.Complexity * pc.ComplexityWeight)
}

/*
	trie les tâches en fonction de leur priorité.

Creation d'une methode de tri par priorité de traitement
*/
func (pc PriorityCalculator) SortTasks(tasks []*Task) {
	// Calculer la priorité de chaque tâche
	for _, task := range tasks {
		pc.CalculatePriority(task)
	}

	// Trier les tâches par priorité décroissante
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority > tasks[j].Priority
	})
}

// LoadTasksFromFile charge les tâches depuis un fichier JSON.
func LoadTasksFromFile(filename string) ([]*Task, error) {
	// Lire le contenu du fichier
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Lire et désérialiser le JSON
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var tasks []*Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func main() {
	// Charger les tâches depuis un fichier
	tasks, err := LoadTasksFromFile("fichier.json")
	if err != nil {
		fmt.Printf("Erreur lors du chargement des tâches : %v\n", err)
		return
	}

	// Définir les pondérations
	// en golang on doit initiliser une struture a l'appel
	calculator := PriorityCalculator{
		UrgencyWeight:    2,
		ImportanceWeight: 3,
		ComplexityWeight: 4,
	}
	fmt.Println("juste presentatin des elements de la taches")
	// Démarrer le chronométrage
	start := time.Now()

	// Trier les tâches par priorité
	calculator.SortTasks(tasks)

	// Arrêter le chronométrage
	elapsed := time.Since(start)

	var t []*Task
	fmt.Println("la presentation simple de la tache")
	for _, val := range t {
		fmt.Println("voice la tache", val)
	}

	// Afficher les tâches triées
	fmt.Println("The task for priorite :")
	for _, task := range tasks {
		fmt.Printf("%s - Priorité : %.2f\n", task.Name, task.Priority)
	}

	// Afficher le temps d'exécution
	fmt.Printf("\nTime : %s\n", elapsed)
}
