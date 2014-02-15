package grago

import (
	"fmt"
	"testing"
)

func createGraphsp() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphsp2() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphsp3() *Graph {
	graph := NewGraph(false, false, false)
	graph.AddLink(stringer("1"), stringer("2"), 1)
	graph.AddLink(stringer("1"), stringer("5"), 1)
	graph.AddLink(stringer("2"), stringer("3"), 1)
	graph.AddLink(stringer("2"), stringer("4"), 1)
	graph.AddLink(stringer("2"), stringer("5"), 1)
	graph.AddLink(stringer("3"), stringer("4"), 1)
	graph.AddLink(stringer("4"), stringer("5"), 1)
	return graph
}

func ExampleGraph_EulerPath() {
	fmt.Println(createGraphsp().EulerPath())
	fmt.Println(createGraphsp2().EulerPath())

	// Output:
	// [3-(8)->5 5-(10)->4 4-(5)->2 2-(2)->3]
	// []
}

func ExampleGraph_HamiltonPath() {
	graph := createGraphsp()

	fmt.Println(graph.HamiltonPath())

	graph.RemoveNode(stringer("alpha"))

	fmt.Println(graph.HamiltonPath())

	// Output:
	// []
	// [2-(2)->3 3-(8)->5 5-(10)->4]
}

func TestEulerPath(t *testing.T) {
	path := createGraphsp().EulerPath()

	if len(path) != 4 {
		t.Fail()
	}

	for _, link := range path {
		if link.Start == stringer("alpha") || link.End == stringer("alpha") {
			t.Fail()
		}
	}
}

func TestNoEulerPath(t *testing.T) {
	if len(createGraphsp2().EulerPath()) != 0 {
		t.Fail()
	}
}

func TestHamiltonPath(t *testing.T) {
	graph := createGraphsp()
	graph.RemoveNode(stringer("alpha"))

	if len(graph.HamiltonPath()) != 3 {
		t.Fail()
	}
}

func TestNoHamiltonPath(t *testing.T) {
	if len(createGraphsp().HamiltonPath()) != 0 {
		t.Fail()
	}
}

func TestHamiltonPath2(t *testing.T) {
	graph := createGraphsp3()

	if len(graph.HamiltonPath()) != 4 {
		t.Fail()
	}
}

func TestEulerPath2(t *testing.T) {
	path := createGraphsp3().EulerPath()

	if len(path) != 7 {
		t.Fail()
	}
}
