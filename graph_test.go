package grago

import (
	"fmt"
	"io/ioutil"
)

func createGraph() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode(stringer("alpha"))
	graph.AddLink(stringer("2"), stringer("alpha"), 2)
	graph.AddLink(stringer("2"), stringer("3"), 2)
	graph.AddLink(stringer("2"), stringer("4"), 5)
	graph.AddLink(stringer("3"), stringer("5"), 8)
	graph.AddLink(stringer("5"), stringer("4"), 10)
	return graph
}

func ExampleLink_String() {
	fmt.Println(NewLink(stringer("1"), stringer("4"), true, 23))
	fmt.Println(NewLink(stringer("1"), stringer("4"), false, 321321))

	// Output:
	// 1-(23)->4
	// 1--4
}

func ExampleNode_AdjacentNodes() {
	node := NewNode()
	node.Adjacent[stringer("2")] = 1
	node.Adjacent[stringer("5")] = 234

	fmt.Println(node.AdjacentNodes())

	// Output:
	// [2 5]
}

func ExampleReadGraph() {
	data, _ := ioutil.ReadFile("exampleGraph.txt")
	graph := ReadGraph(string(data), false)

	fmt.Println(graph)

	// Output:
	// true true false
	// alpha
	// 2
	// 3
	// 2 -- 3 2
}

func ExampleGraph_String() {
	data, _ := ioutil.ReadFile("exampleGraph.txt")
	graph := ReadGraph(string(data), false)

	fmt.Println(graph)

	// Output:
	// true true false
	// alpha
	// 2
	// 3
	// 2 -- 3 2
}

func ExampleGraph_AddNode() {
	graph := NewGraph(false, true, false)

	fmt.Println(graph.AddNode(stringer("a")))
	fmt.Println(graph.AddNode(stringer("a")))
	fmt.Println(graph.AddNode(stringer("b")))

	// Output:
	// true
	// false
	// true
}

func ExampleGraph_AddLink() {
	graph := NewGraph(false, true, false)

	fmt.Println(graph.AddLink(stringer("a"), stringer("b"), 2))
	fmt.Println(graph.AddLink(stringer("a"), stringer("b"), 5))

	// Output:
	// true
	// false
}

func ExampleGraph_RemoveNode() {
	graph := NewGraph(false, true, false)

	graph.AddNode(stringer("a"))

	fmt.Println(graph.RemoveNode(stringer("a")))
	fmt.Println(graph.RemoveNode(stringer("a")))
	fmt.Println(graph.RemoveNode(stringer("b")))

	// Output:
	// true
	// false
	// false
}

func ExampleGraph_RemoveLink() {
	graph := NewGraph(false, true, false)

	graph.AddLink(stringer("a"), stringer("b"), 2)

	fmt.Println(graph.RemoveLink(stringer("a"), stringer("b")))
	fmt.Println(graph.RemoveLink(stringer("a"), stringer("b")))
	fmt.Println(graph.RemoveLink(stringer("a"), stringer("c")))

	// Output:
	// true
	// false
	// false
}

func ExampleGraph_OutgoingLinksCount() {
	graph := createGraph()

	fmt.Println(graph.OutgoingLinksCount(stringer("alpha")))
	fmt.Println(graph.OutgoingLinksCount(stringer("2")))
	fmt.Println(graph.OutgoingLinksCount(stringer("3")))

	// Output:
	// 1
	// 3
	// 2
}

func ExampleGraph_IncomingLinksCount() {
	graph := createGraph()

	fmt.Println(graph.IncomingLinksCount(stringer("alpha")))
	fmt.Println(graph.IncomingLinksCount(stringer("2")))
	fmt.Println(graph.IncomingLinksCount(stringer("3")))

	// Output:
	// 1
	// 3
	// 2
}

func ExampleGraph_Nodes() {
	fmt.Println(createGraph().Nodes())

	// Output:
	// [alpha 2 3 4 5]
}

func ExampleGraph_Links() {
	data, _ := ioutil.ReadFile("exampleGraph.txt")
	graph := ReadGraph(string(data), false)
	fmt.Println(graph.Links())

	// Output:
	// [2-(2)->3]
}
