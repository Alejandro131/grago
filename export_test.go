package grago

import (
	"fmt"
	"testing"
)

func createGraphex() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func ExampleGraph_Export() {
	graph := createGraphex()

	fmt.Println(graph.Export([]Link{}, false, [][]fmt.Stringer{}))
	fmt.Println(graph.Export(graph.DFS(stringer("2")), true, [][]fmt.Stringer{}))

	// Output:
	// graph {"alpha" "2" "3" "4" "5" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "4"--"5" [label="10"]; }
	// graph {"2"--"3" [fontcolor=red color=red label="2 (1)"]; "3"--"5" [fontcolor=red color=red label="8 (2)"]; "5"--"4" [fontcolor=red color=red label="10 (3)"]; "alpha" "2"--"4" [label="5"]; }
}

func TestExport(t *testing.T) {
	exported := `graph {"alpha" "2" "3" "4" "5" "2"--"3" [label="2"]; "2"--"4" [label="5"]; "3"--"5" [label="8"]; "4"--"5" [label="10"]; }`

	graph := createGraphex()

	if exported != graph.Export([]Link{}, false, [][]fmt.Stringer{}) {
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

	if exported != graph.Export(graph.DFS(stringer("2")), true, [][]fmt.Stringer{}) {
		t.Fail()
	}
}

func TestExportHighlightsUnordered(t *testing.T) {
	exported := `graph {"2"--"3" [fontcolor=red color=red label="2"]; "2"--"4" [fontcolor=red color=red label="5"]; "5"--"3" [fontcolor=red color=red label="8"]; "alpha" "4"--"5" [label="10"]; }`

	graph := createGraphex()

	if exported != graph.Export(graph.MST(), false, [][]fmt.Stringer{}) {
		t.Fail()
	}
}

func TestExportHighlightsOrderedConnectedComponents(t *testing.T) {
	exported := `graph {subgraph cluster0 {"3" "4" "5" "2" }subgraph cluster1 {"alpha" }"2"--"3" [fontcolor=red color=red label="2 (1)"]; "3"--"5" [fontcolor=red color=red label="8 (2)"]; "5"--"4" [fontcolor=red color=red label="10 (3)"]; "2"--"4" [label="5"]; }`

	graph := createGraphex()

	if exported != graph.Export(graph.DFS(stringer("2")), true, graph.ConnectedComponents()) {
		t.Fail()
	}
}
