package grago

import (
	"fmt"
	"testing"
)

func createGraphd() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphNegative() *Graph {
	graph := NewGraph(true, true, true)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), -5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	graph.AddLink(stringer("4"), stringer("5"), 10)
	graph.AddLink(stringer("3"), stringer("2"), 2)
	return graph
}

func createGraphNegative2() *Graph {
	graph := NewGraph(false, true, true)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), -5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func ExampleGraph_Floyd() {
	paths := createGraphd().Floyd()

	fmt.Println(paths[stringer("2")][stringer("5")])
	fmt.Println(paths[stringer("3")][stringer("4")])

	// Output:
	// 10
	// 7
}

func ExampleGraph_MinPaths_dijkstra() {
	graph := createGraphd()

	fmt.Println(graph.MinPaths(stringer("2"))[stringer("5")])
	fmt.Println(graph.MinPaths(stringer("3"))[stringer("4")])

	// Output:
	// 10
	// 7
}

func ExampleGraph_MinPaths_fordBellman() {
	graph := createGraphNegative()

	fmt.Println(graph.MinPaths(stringer("2"))[stringer("5")])
	fmt.Println(graph.MinPaths(stringer("3"))[stringer("4")])

	// Output:
	// 5
	// -3
}

func TestFloyd(t *testing.T) {
	paths := createGraphd().Floyd()

	if paths[stringer("2")][stringer("5")] != 10 {
		t.Fail()
	}

	if paths[stringer("3")][stringer("4")] != 7 {
		t.Fail()
	}
}

func TestDijkstra(t *testing.T) {
	paths := createGraphd().MinPaths(stringer("2"))

	if paths[stringer("5")] != 10 {
		t.Fail()
	}
}

func TestDijkstra2(t *testing.T) {
	paths := createGraphd().MinPaths(stringer("3"))

	if paths[stringer("4")] != 7 {
		t.Fail()
	}
}

func TestFordBellman(t *testing.T) {
	paths := createGraphNegative().MinPaths(stringer("2"))

	if paths[stringer("5")] != 5 {
		t.Fail()
	}
}

func TestFordBellman2(t *testing.T) {
	paths := createGraphNegative().MinPaths(stringer("3"))

	if paths[stringer("4")] != -3 {
		t.Fail()
	}
}

func TestFordBellman3(t *testing.T) {
	paths := createGraphNegative2().MinPaths(stringer("2"))

	if paths[stringer("5")] != -25 {
		t.Fail()
	}
}

func TestFordBellman4(t *testing.T) {
	paths := createGraphNegative2().MinPaths(stringer("3"))

	if paths[stringer("4")] != -23 {
		t.Fail()
	}
}
