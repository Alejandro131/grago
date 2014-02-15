package grago

import (
	"fmt"
	"testing"
)

func createGraphm() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func createGraphm2() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddLink(stringer("alpha"), stringer("beta"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	graph.AddLink(stringer("5"), stringer("zeta"), 10)
	graph.AddLink(stringer("2"), stringer("zeta"), 1)
	graph.AddLink(stringer("alpha"), stringer("zeta"), 200)
	return graph
}

func ExampleGraph_MST() {
	fmt.Println(createGraphm().MST())

	// Output:
	// [2-(2)->3 2-(5)->4 5-(8)->3]
}

func TestMST(t *testing.T) {
	treeLinks := createGraphm().MST()

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

func TestMST2(t *testing.T) {
	treeLinks := createGraphm2().MST()

	if len(treeLinks) != 6 {
		t.Fail()
	}

	if treeLinks[0].Weight != 1 {
		t.Fail()
	}

	if treeLinks[1].Weight != 2 {
		t.Fail()
	}

	if treeLinks[2].Weight != 5 {
		t.Fail()
	}

	if treeLinks[3].Weight != 8 {
		t.Fail()
	}

	if treeLinks[4].Weight != 10 {
		t.Fail()
	}

	if treeLinks[5].Weight != 200 {
		t.Fail()
	}
}
