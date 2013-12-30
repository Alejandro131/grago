package grago

import "testing"

func createGraph() Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func TestMaxFlow(t *testing.T) {
	if createGraph().MaxFlow("2", "5") != 7 {
		t.Fail()
	}
}

func TestMaxFlow2(t *testing.T) {
	if createGraph().MaxFlow("3", "4") != 10 {
		t.Fail()
	}
}

func TestMaxFlow3(t *testing.T) {
	if createGraph().MaxFlow("alpha", "4") != 0 {
		t.Fail()
	}
}