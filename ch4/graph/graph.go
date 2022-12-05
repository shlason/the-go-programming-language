package main

import (
	"fmt"
)

var graph = map[string]map[string]bool{}

func main() {
	// graph["a"]["bottom"] = true
	fmt.Println(graph["b"] == nil)
}

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
