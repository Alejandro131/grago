package grago

import (
	"fmt"
	"testing"
)

func createGraphc() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleReachableNodes_Graph() {
	graph := createGraphc()

	fmt.Println(graph.ReachableNodes("2"))
	fmt.Println(graph.ReachableNodes("alpha"))

	// Output:
	// [3 4 5]
	// []
}

func ExampleConnectedComponents_Graph() {
	fmt.Println(createGraphc().ConnectedComponents())

	// Output:
	// [[alpha] [3 4 5 2]]
}

func TestReachableNodes(t *testing.T) {
	reachable := createGraphc().ReachableNodes("2")

	if len(reachable) != 3 {
		t.Fail()
	}

	for _, node := range reachable {
		if node == "alpha" {
			t.Fail()
		}
	}
}

func TestUnreachableNodes(t *testing.T) {
	if len(createGraphc().ReachableNodes("alpha")) != 0 {
		t.Fail()
	}
}

func TestConnectedComponents(t *testing.T) {
	if len(createGraphc().ConnectedComponents()) != 2 {
		t.Fail()
	}
}
