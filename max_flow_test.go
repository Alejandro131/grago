package grago

import (
	"fmt"
	"testing"
)

func createGraphmf() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func createGraphmf2() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink("s", "o", 3)
	graph.AddLink("s", "p", 3)
	graph.AddLink("o", "p", 2)
	graph.AddLink("p", "r", 2)
	graph.AddLink("o", "q", 3)
	graph.AddLink("q", "r", 4)
	graph.AddLink("q", "t", 2)
	graph.AddLink("r", "t", 3)
	return graph
}

func createGraphmf3() *Graph {
	graph := NewGraph(true, true, false)
	graph.AddLink("s", "u", 10)
	graph.AddLink("s", "v", 5)
	graph.AddLink("u", "v", 15)
	graph.AddLink("v", "t", 10)
	graph.AddLink("u", "t", 5)
	return graph
}

func ExampleMaxFlow_Graph() {
	graph := createGraphmf()

	fmt.Println(graph.MaxFlow("2", "5"))
	fmt.Println(graph.MaxFlow("3", "4"))
	fmt.Println(graph.MaxFlow("alpha", "4"))

	// Output:
	// 7
	// 10
	// 0
}

func TestMaxFlow(t *testing.T) {
	if createGraphmf().MaxFlow("2", "5") != 7 {
		t.Fail()
	}
}

func TestMaxFlow2(t *testing.T) {
	if createGraphmf().MaxFlow("3", "4") != 10 {
		t.Fail()
	}
}

func TestMaxFlow3(t *testing.T) {
	if createGraphmf().MaxFlow("alpha", "4") != 0 {
		t.Fail()
	}
}

func TestMaxFlow4(t *testing.T) {
	if createGraphmf2().MaxFlow("s", "t") != 5 {
		t.Fail()
	}
}

func TestMaxFlow5(t *testing.T) {
	if createGraphmf3().MaxFlow("s", "t") != 15 {
		t.Fail()
	}
}
