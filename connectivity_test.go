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

func createGraphc2() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleGraph_ReachableNodes() {
	graph := createGraphc()

	fmt.Println(graph.ReachableNodes("2"))
	fmt.Println(graph.ReachableNodes("alpha"))

	// Output:
	// [3 4 5]
	// []
}

func ExampleGraph_ConnectedComponents() {
	fmt.Println(createGraphc().ConnectedComponents())
	fmt.Println(createGraphc2().ConnectedComponents())

	// Output:
	// [[3 4 5 2] [alpha]]
	// [[2] [3] [5] [4] [alpha]]
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

func TestConnectedComponents2(t *testing.T) {
	if len(createGraphc2().ConnectedComponents()) != 5 {
		t.Fail()
	}
}
