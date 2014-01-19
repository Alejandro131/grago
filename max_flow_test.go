package grago

import (
	"fmt"
	"testing"
)

func createGraph() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleMaxFlow_Graph() {
	graph := createGraph()

	fmt.Println(graph.MaxFlow("2", "5"))
	fmt.Println(graph.MaxFlow("3", "4"))
	fmt.Println(graph.MaxFlow("alpha", "4"))

	// Output:
	// 7
	// 10
	// 0
}
