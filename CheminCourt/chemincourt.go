package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Structure d'un nœud dans la file de priorité
type Node struct {
	id       int
	distance float64
	index    int
}

// File de priorité (min-heap)
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.index = len(*pq)
	*pq = append(*pq, n)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old[len(old)-1]
	*pq = old[:len(old)-1]
	return n
}

// Graph struct
type Graph struct {
	nodes map[int]map[int]float64
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int]map[int]float64)}
}

func (g *Graph) AddEdge(u, v int, weight float64) {
	if g.nodes[u] == nil {
		g.nodes[u] = make(map[int]float64)
	}
	if g.nodes[v] == nil {
		g.nodes[v] = make(map[int]float64)
	}
	g.nodes[u][v] = weight
	g.nodes[v][u] = weight // Si graphe non orienté
}

// Algorithme de Dijkstra
func (g *Graph) Dijkstra(start, end int) ([]int, float64) {
	dist := make(map[int]float64)
	prev := make(map[int]int)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialisation
	for node := range g.nodes {
		dist[node] = math.Inf(1)
		prev[node] = -1
	}
	dist[start] = 0
	heap.Push(pq, &Node{id: start, distance: 0})

	// Traitement des nœuds
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Node)

		// Si on atteint le nœud de destination
		if curr.id == end {
			break
		}

		for neighbor, weight := range g.nodes[curr.id] {
			alt := dist[curr.id] + weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = curr.id
				heap.Push(pq, &Node{id: neighbor, distance: alt})
			}
		}
	}

	// Reconstruction du chemin
	path := []int{}
	for at := end; at != -1; at = prev[at] {
		path = append([]int{at}, path...)
	}

	return path, dist[end]
}

func main() {
	graph := NewGraph()
	graph.AddEdge(1, 2, 4)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(3, 2, 2)
	graph.AddEdge(2, 4, 1)
	graph.AddEdge(3, 4, 5)

	start, end := 1, 4
	path, cost := graph.Dijkstra(start, end)
	fmt.Printf("Chemin le plus court de %d à %d : %v avec une distance de 4 kg %.2f\n", start, end, path, cost)
}
