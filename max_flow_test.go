package grago

import (
	"fmt"
	"testing"
)

func createGraphmf() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphmf2() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink(stringer("s"), stringer("o"), 3)
	graph.AddLink(stringer("s"), stringer("p"), 3)
	graph.AddLink(stringer("o"), stringer("p"), 2)
	graph.AddLink(stringer("p"), stringer("r"), 2)
	graph.AddLink(stringer("o"), stringer("q"), 3)
	graph.AddLink(stringer("q"), stringer("r"), 4)
	graph.AddLink(stringer("q"), stringer("t"), 2)
	graph.AddLink(stringer("r"), stringer("t"), 3)
	return graph
}

func createGraphmf3() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink(stringer("s"), stringer("u"), 10)
	graph.AddLink(stringer("s"), stringer("v"), 5)
	graph.AddLink(stringer("u"), stringer("v"), 15)
	graph.AddLink(stringer("v"), stringer("t"), 10)
	graph.AddLink(stringer("u"), stringer("t"), 5)
	return graph
}

func ExampleGraph_MaxFlow() {
	graph := createGraphmf()

	fmt.Println(graph.MaxFlow(stringer("2"), stringer("5")))
	fmt.Println(graph.MaxFlow(stringer("3"), stringer("4")))
	fmt.Println(graph.MaxFlow(stringer("alpha"), stringer("4")))

	// Output:
	// 7
	// 10
	// 0
}

func TestMaxFlow(t *testing.T) {
	if createGraphmf().MaxFlow(stringer("2"), stringer("5")) != 7 {
		t.Fail()
	}
}

func TestMaxFlow2(t *testing.T) {
	if createGraphmf().MaxFlow(stringer("3"), stringer("4")) != 10 {
		t.Fail()
	}
}

func TestMaxFlow3(t *testing.T) {
	if createGraphmf().MaxFlow(stringer("alpha"), stringer("4")) != 0 {
		t.Fail()
	}
}

func TestMaxFlow4(t *testing.T) {
	if createGraphmf2().MaxFlow(stringer("s"), stringer("t")) != 5 {
		t.Fail()
	}
}

func TestMaxFlow5(t *testing.T) {
	if createGraphmf3().MaxFlow(stringer("s"), stringer("t")) != 15 {
		t.Fail()
	}
}
