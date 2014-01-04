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

func createGraph2() Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleEulerPath_Graph() {
	fmt.Println(createGraph().EulerPath())
	fmt.Println(createGraph2().EulerPath())

	// Output:
	// [2-(2)->3 3-(8)->5 5-(10)->4 4-(5)->2]
	// []
}

func ExampleHamiltonPath_Graph() {
	graph := createGraph()

	fmt.Println(graph.HamiltonPath())

	graph.RemoveNode("alpha")

	fmt.Println(graph.HamiltonPath())

	// Output:
	// []
	// [2-(2)->3 3-(8)->5 5-(10)->4 4-(5)->2]
}
