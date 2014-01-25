package grago

import (
	"fmt"
	"testing"
)

func createGraphex() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleExport_Graph() {
	graph := createGraphex()

	fmt.Println(graph.Export([]Link{}, false, [][]string{}))
	fmt.Println(graph.Export(graph.DFS("2"), true, [][]string{}))

	// Output:
	// graph { "alpha" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "5"--"4" [label="10"];}
	// graph { "alpha" "2"--"3" [label="2 (1)" fontcolor=red color=red]; "2"--"4" [label="5"]; "3"--"5" [label="8 (2)" fontcolor=red color=red]; "5"--"4" [label="10 (3)" fontcolor=red color=red];}
}

func TestExport(t *testing.T) {
	exported := `graph { "alpha" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "5"--"4" [label="10"];}`

	graph := createGraphex()

	if exported != graph.ExportGraph([]Link{}) {
		t.Fail()
	}
}

func TestExportHighlights(t *testing.T) {
	exported := `graph { "alpha" "2"--"3" [label="2 (1)" fontcolor=red color=red]; "2"--"4" [label="5"]; "3"--"5" [label="8 (2)" fontcolor=red color=red]; "5"--"4" [label="10 (3)" fontcolor=red color=red];}`

	graph := createGraphex()

	if exported != graph.ExportGraph(graph.DFS("2")) {
		t.Fail()
	}
}
