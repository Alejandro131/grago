package grago

import (
	"fmt"
	"testing"
)

func createGraph() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleMST_Graph() {
	fmt.Println(createGraph().MST())

	// Output:
	// [2-(2)->3 4-(5)->2 3-(8)->5]
}

func TestMST(t *testing.T) {
	treeLinks := createGraph().MST()

	if len(treeLinks) != 3 {
		t.Fail()
	}

	if treeLinks[0].Weight != 2 {
		t.Fail()
	}

	if treeLinks[1].Weight != 5 {
		t.Fail()
	}

	if treeLinks[2].Weight != 8 {
		t.Fail()
	}
}
