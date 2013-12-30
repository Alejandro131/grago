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

func TestConnectedComponents(t *testing.T) {
	components := createGraph().ConnectedComponents()
	
	if len(components) != 2 {
		t.Fail()
	}
}