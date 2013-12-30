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

func TestBFS(t *testing.T) {
	layers := createGraph().BFS("2")
	
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

func TestDFS(t *testing.T) {
	links := createGraph().DFS("2")
	
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