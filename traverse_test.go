package grago

import (
	"fmt"
	"testing"
)

func createGrapht() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func ExampleGraph_BFS() {
	fmt.Println(createGrapht().BFS(stringer("2")))

	// Output:
	// [[2] [3 4] [5]]
}

func ExampleGraph_DFS() {
	fmt.Println(createGrapht().DFS(stringer("2")))

	// Output:
	// [2-(2)->3 3-(8)->5 5-(10)->4]
}

func TestBFS(t *testing.T) {
	layers := createGrapht().BFS(stringer("2"))

	if len(layers) != 3 {
		t.Fail()
	}

	if layers[0][0] != stringer("2") {
		t.Fail()
	}

	if layers[2][0] != stringer("5") {
		t.Fail()
	}
}

func TestBFSNoWay(t *testing.T) {
	layers := createGrapht().BFS(stringer("alpha"))

	if len(layers) != 1 {
		t.Fail()
	}
}

func TestDFS(t *testing.T) {
	links := createGrapht().DFS(stringer("2"))

	if len(links) != 3 {
		t.Fail()
	}

	if links[0].Weight != 2 {
		t.Fail()
	}

	if links[1].Weight != 8 {
		t.Fail()
	}

	if links[2].Weight != 10 {
		t.Fail()
	}
}

func TestDFSNoLinks(t *testing.T) {
	links := createGrapht().DFS(stringer("alpha"))

	if len(links) != 0 {
		t.Fail()
	}
}
