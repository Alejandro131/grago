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

func TestReachableNodes(t *testing.T) {
	reachable := createGraph().ReachableNodes("2")
	
	if len(reachable) != 3 {
		t.Fail()
	}
	
	for _, node := range reachable {
		if node == "alpha" {
			t.Fail()
		}
	}
}

func TestUnreachableNodes(t *testing.T) {
	reachable := createGraph().ReachableNodes("alpha")
	
	if len(reachable) != 0 {
		t.Fail()
	}
}

func TestConnectedComponents(t *testing.T) {
	components := createGraph().ConnectedComponents()
	
	if len(components) != 2 {
		t.Fail()
	}
}