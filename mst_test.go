package grago

import (
	"fmt"
	"testing"
)

func createGraph() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleMST_Graph() {
	fmt.Println(createGraph().MST())

	// Output:
	// [2-(2)->3 4-(5)->2 3-(8)->5]
}
