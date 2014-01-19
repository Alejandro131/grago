package grago

import (
	"fmt"
	"testing"
)

func createGraph() Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func createGraphNegative() Graph {
	graph := NewGraph(false, true, true)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", -5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleFloyd_Graph() {
	paths := createGraph().Floyd()

	fmt.Println(paths["2"]["5"])
	fmt.Println(paths["3"]["4"])

	// Output:
	// 10
	// 7
}

func ExampleMinPaths_Graph_dijkstra() {
	graph := createGraph()

	fmt.Println(graph.MinPaths("2")["5"])
	fmt.Println(graph.MinPaths("3")["4"])

	// Output:
	// 10
	// 7
}

func ExampleMinPaths_Graph_fordBellman() {
	graph := createGraphNegative()

	fmt.Println(graph.MinPaths("2")["5"])
	fmt.Println(graph.MinPaths("3")["4"])

	// Output:
	// 5
	// -3
}
