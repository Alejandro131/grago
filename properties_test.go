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
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	return graph
}

func ExampleHasCycle_Graph() {
	fmt.Println(createGraph().HasCycle())
	fmt.Println(createGraph2().HasCycle())

	// Output:
	// true
	// false
}

func ExampleIsPlanar_Graph() {
	graph := createGraph()

	fmt.Println(graph.IsPlanar())

	graph.AddLink("2", "5", 20)
	graph.AddLink("3", "4", 15)
	graph.AddLink("2", "alpha", 1)
	graph.AddLink("3", "alpha", 1)
	graph.AddLink("4", "alpha", 1)
	graph.AddLink("5", "alpha", 1)

	fmt.Println(graph.IsPlanar())

	// Output:
	// true
	// false
}

func ExampleIsBipartite_Graph() {
	fmt.Println(createGraph2().IsBipartite())
	fmt.Println(createGraph().IsBipartite())

	// Output:
	// true
	// false
}

func TestCycle(t *testing.T) {
	if createGraph().HasCycle() == false {
		t.Fail()
	}
}

func TestNoCycle(t *testing.T) {
	if createGraph2().HasCycle() == true {
		t.Fail()
	}
}

func TestPlanar(t *testing.T) {
	if createGraph().IsPlanar() == false {
		t.Fail()
	}
}

func TestNonPlanar(t *testing.T) {
	graph := createGraph()
	graph.AddLink("2", "5", 20)
	graph.AddLink("3", "4", 15)
	graph.AddLink("2", "alpha", 1)
	graph.AddLink("3", "alpha", 1)
	graph.AddLink("4", "alpha", 1)
	graph.AddLink("5", "alpha", 1)

	if graph.IsPlanar() == true {
		t.Fail()
	}
}

func TestBipartite(t *testing.T) {
	if createGraph2().IsBipartite() == false {
		t.Fail()
	}
}

func TestNonBipartite(t *testing.T) {
	if createGraph().IsBipartite() == true {
		t.Fail()
	}
}
