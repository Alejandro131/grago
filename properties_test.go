package grago

import (
	"fmt"
	"testing"
)

func createGraphpr() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphpr2() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	return graph
}

func createGraphpr3() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	return graph
}

func ExampleGraph_HasCycle() {
	fmt.Println(createGraphpr().HasCycle())
	fmt.Println(createGraphpr2().HasCycle())
	fmt.Println(createGraphpr3().HasCycle())

	// Output:
	// true
	// true
	// false
}

func ExampleGraph_IsPlanar() {
	graph := createGraphpr()
	graph.RemoveNode(stringer("alpha"))

	fmt.Println(graph.IsPlanar())

	graph.AddLink(stringer("2"), stringer("5"), 20)
	graph.AddLink(stringer("3"), stringer("4"), 15)
	graph.AddLink(stringer("2"), stringer("alpha"), 1)
	graph.AddLink(stringer("3"), stringer("alpha"), 1)
	graph.AddLink(stringer("4"), stringer("alpha"), 1)
	graph.AddLink(stringer("5"), stringer("alpha"), 1)

	fmt.Println(graph.IsPlanar())

	// Output:
	// true
	// false
}

func ExampleGraph_IsBipartite() {
	graph := createGraphpr()

	fmt.Println(createGraphpr2().IsBipartite())
	fmt.Println(graph.IsBipartite())

	graph.AddLink(stringer("2"), stringer("5"), 20)

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
	graph := createGraphpr()
	graph.RemoveNode(stringer("alpha"))

	if graph.IsPlanar() == false {
		t.Fail()
	}
}

func TestNonPlanar(t *testing.T) {
	graph := createGraphpr()
	graph.AddLink(stringer("2"), stringer("5"), 20)
	graph.AddLink(stringer("3"), stringer("4"), 15)
	graph.AddLink(stringer("2"), stringer("alpha"), 1)
	graph.AddLink(stringer("3"), stringer("alpha"), 1)
	graph.AddLink(stringer("4"), stringer("alpha"), 1)
	graph.AddLink(stringer("5"), stringer("alpha"), 1)

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

	graph.AddLink(stringer("3"), stringer("4"), 7)

	if graph.IsBipartite() == true {
		t.Fail()
	}
}
