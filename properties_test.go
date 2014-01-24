package grago

import (
	"fmt"
	"testing"
)

func createGraphpr() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func createGraphpr2() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	return graph
}

func createGraphpr3() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	return graph
}

func ExampleHasCycle_Graph() {
	fmt.Println(createGraphpr().HasCycle())
	fmt.Println(createGraphpr2().HasCycle())
	fmt.Println(createGraphpr3().HasCycle())

	// Output:
	// true
	// true
	// false
}

func ExampleIsPlanar_Graph() {
	graph := createGraphpr()

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
	graph := createGraphpr()

	fmt.Println(createGraphpr2().IsBipartite())
	fmt.Println(graph.IsBipartite())
	
	graph.AddLink("2", "5", 20)
	
	fmt.Println(graph.IsBipartite())

	// Output:
	// true
	// true
	// false
}

func TestCycle(t *testing.T) {
	if createGraphpr().HasCycle() == false {
		t.Fail()
	}
}

func TestCycle2(t *testing.T) {
	if createGraphpr2().HasCycle() == false {
		t.Fail()
	}
}

func TestNoCycle(t *testing.T) {
	if createGraphpr3().HasCycle() == true {
		t.Fail()
	}
}

func TestPlanar(t *testing.T) {
	if createGraphpr().IsPlanar() == false {
		t.Fail()
	}
}

func TestNonPlanar(t *testing.T) {
	graph := createGraphpr()
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
	if createGraphpr2().IsBipartite() == false {
		t.Fail()
	}
}

func TestBipartite2(t *testing.T) {
	if createGraphpr().IsBipartite() == false {
		t.Fail()
	}
}

func TestNonBipartite(t *testing.T) {
	graph := createGraphpr()
	
	graph.AddLink("3", "4", 7)
	
	if graph.IsBipartite() == true {
		t.Fail()
	}
}
