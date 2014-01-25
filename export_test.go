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
	// graph {"alpha" "2" "3" "4" "5" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "4"--"5" [label="10"]; }
	// graph {"2"--"3" [fontcolor=red color=red label="2 (1)"]; "3"--"5" [fontcolor=red color=red label="8 (2)"]; "5"--"4" [fontcolor=red color=red label="10 (3)"]; "alpha" "2"--"4" [label="5"]; }
}

func TestExport(t *testing.T) {
	exported := `graph {"alpha" "2" "3" "4" "5" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "4"--"5" [label="10"]; }`

	graph := createGraphex()

	if exported != graph.Export([]Link{}, false, [][]string{}) {
		t.Fail()
	}
}

func TestExportConnectedComponents(t *testing.T) {
	exported := `graph {subgraph cluster0 {"3" "4" "5" "2" }subgraph cluster1 {"alpha" }"2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "4"--"5" [label="10"]; }`

	graph := createGraphex()

	if exported != graph.Export([]Link{}, false, graph.ConnectedComponents()) {
		t.Fail()
	}
}

func TestExportHighlightsOrdered(t *testing.T) {
	exported := `graph {"2"--"3" [fontcolor=red color=red label="2 (1)"]; "3"--"5" [fontcolor=red color=red label="8 (2)"]; "5"--"4" [fontcolor=red color=red label="10 (3)"]; "alpha" "2"--"4" [label="5"]; }`

	graph := createGraphex()

	if exported != graph.Export(graph.DFS("2"), true, [][]string{}) {
		t.Fail()
	}
}

func TestExportHighlightsUnordered(t *testing.T) {
	exported := `graph {"2"--"3" [fontcolor=red color=red label="2"]; "2"--"4" [fontcolor=red color=red label="5"]; "5"--"3" [fontcolor=red color=red label="8"]; "alpha" "4"--"5" [label="10"]; }`

	graph := createGraphex()

	if exported != graph.Export(graph.MST(), false, [][]string{}) {
		t.Fail()
	}
}

func TestExportHighlightsOrderedConnectedComponents(t *testing.T) {
	exported := `graph {subgraph cluster0 {"3" "4" "5" "2" }subgraph cluster1 {"alpha" }"2"--"3" [fontcolor=red color=red label="2 (1)"]; "3"--"5" [fontcolor=red color=red label="8 (2)"]; "5"--"4" [fontcolor=red color=red label="10 (3)"]; "2"--"4" [label="5"]; }`

	graph := createGraphex()

	if exported != graph.Export(graph.DFS("2"), true, graph.ConnectedComponents()) {
		t.Fail()
	}
}
