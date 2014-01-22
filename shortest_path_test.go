package grago

import (
	"fmt"
	"testing"
)

func createGraphd() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func createGraphNegative() *Graph {
	graph := NewGraph(true, true, true)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", -5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	graph.AddLink("4", "5", 10)
	graph.AddLink("3", "2", 2)
	return graph
}

func createGraphNegative2() *Graph {
	graph := NewGraph(false, true, true)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", -5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleFloyd_Graph() {
	paths := createGraphd().Floyd()

	fmt.Println(paths["2"]["5"])
	fmt.Println(paths["3"]["4"])

	// Output:
	// 10
	// 7
}

func ExampleMinPaths_Graph_dijkstra() {
	graph := createGraphd()

	fmt.Println(graph.MinPaths("2")["5"])
	fmt.Println(graph.MinPaths("3")["4"])

	// Output:
	// 10
	// 7
}

func ExampleMinPaths_Graph_fordBellman() {
	graph := createGraphNegative()

	fmt.Println(graph.MinPaths("2")["5"])
	fmt.Println(graph.MinPaths("3")["4"])

	// Output:
	// 5
	// -3
}

func TestFloyd(t *testing.T) {
	paths := createGraphd().Floyd()

	if paths["2"]["5"] != 10 {
		t.Fail()
	}

	if paths["3"]["4"] != 7 {
		t.Fail()
	}
}

func TestDijkstra(t *testing.T) {
	paths := createGraphd().MinPaths("2")

	if paths["5"] != 10 {
		t.Fail()
	}
}

func TestDijkstra2(t *testing.T) {
	paths := createGraphd().MinPaths("3")

	if paths["4"] != 7 {
		t.Fail()
	}
}

func TestFordBellman(t *testing.T) {
	paths := createGraphNegative().MinPaths("2")

	if paths["5"] != 5 {
		t.Fail()
	}
}

func TestFordBellman2(t *testing.T) {
	paths := createGraphNegative().MinPaths("3")

	if paths["4"] != -3 {
		t.Fail()
	}
}

func TestFordBellman3(t *testing.T) {
	paths := createGraphNegative2().MinPaths("2")

	if paths["5"] != -25 {
		t.Fail()
	}
}

func TestFordBellman4(t *testing.T) {
	paths := createGraphNegative2().MinPaths("3")

	if paths["4"] != -23 {
		t.Fail()
	}
}
