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

func createGraph2() *Graph {
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

func TestEulerPath(t *testing.T) {
	path := createGraph().EulerPath()

	if len(path) != 4 {
		t.Fail()
	}

	for _, link := range path {
		if link.Start == "alpha" || link.End == "alpha" {
			t.Fail()
		}
	}
}

func TestNoEulerPath(t *testing.T) {
	if len(createGraph2().EulerPath()) != 0 {
		t.Fail()
	}
}

func TestHamiltonPath(t *testing.T) {
	graph := createGraph()
	graph.RemoveNode("alpha")

	if len(graph.HamiltonPath()) != 4 {
		t.Fail()
	}
}

func TestNoHamiltonPath(t *testing.T) {
	if len(createGraph().HamiltonPath()) != 0 {
		t.Fail()
	}
}
