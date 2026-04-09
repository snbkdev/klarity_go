package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("A", "B")
	addEdge("A", "C")
	addEdge("B", "C")
	addEdge("C", "A")

	fmt.Println(hasEdge("A", "B")) // true
	fmt.Println(hasEdge("A", "D")) // false
	fmt.Println(hasEdge("B", "A")) // false
	fmt.Println(hasEdge("C", "A")) // true

	fmt.Println(graph) // map[A:map[B:true C:true] B:map[C:true] C:map[A:true]]

	from := "A"
	fmt.Printf("Из %v можно попасть в: ", from) // Из A можно попасть в: B C
	for to := range graph[from] {
		fmt.Printf("%v ", to)
	}
}
