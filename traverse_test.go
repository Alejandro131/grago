package grago

import (
	"fmt"
	"testing"
)

func createGrapht() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleGraph_BFS() {
	fmt.Println(createGrapht().BFS("2"))

	// Output:
	// [[2] [3 4] [5]]
}

func ExampleGraph_DFS() {
	fmt.Println(createGrapht().DFS("2"))

	// Output:
	// [2-(2)->3 3-(8)->5 5-(10)->4]
}

func TestBFS(t *testing.T) {
	layers := createGrapht().BFS("2")

	if len(layers) != 3 {
		t.Fail()
	}

	if layers[0][0] != "2" {
		t.Fail()
	}

	if layers[2][0] != "5" {
		t.Fail()
	}
}

func TestBFSNoWay(t *testing.T) {
	layers := createGrapht().BFS("alpha")

	if len(layers) != 1 {
		t.Fail()
	}
}

func TestDFS(t *testing.T) {
	links := createGrapht().DFS("2")

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
	links := createGrapht().DFS("alpha")

	if len(links) != 0 {
		t.Fail()
	}
}
