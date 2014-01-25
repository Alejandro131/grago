package grago

import (
	"fmt"
	"io/ioutil"
)

func createGraph() *Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "alpha", 2)
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleLink_String() {
	fmt.Println(NewLink("1", "4", true, 23))
	fmt.Println(NewLink("1", "4", false, 321321))

	// Output:
	// 1-(23)->4
	// 1--4
}

func ExampleNode_AdjacentNodes() {
	node := NewNode()
	node.Adjacent["2"] = 1
	node.Adjacent["5"] = 234

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

	fmt.Println(graph.AddNode("a"))
	fmt.Println(graph.AddNode("a"))
	fmt.Println(graph.AddNode("b"))

	// Output:
	// true
	// false
	// true
}

func ExampleGraph_AddLink() {
	graph := NewGraph(false, true, false)

	fmt.Println(graph.AddLink("a", "b", 2))
	fmt.Println(graph.AddLink("a", "b", 5))

	// Output:
	// true
	// false
}

func ExampleGraph_RemoveNode() {
	graph := NewGraph(false, true, false)

	graph.AddNode("a")

	fmt.Println(graph.RemoveNode("a"))
	fmt.Println(graph.RemoveNode("a"))
	fmt.Println(graph.RemoveNode("b"))

	// Output:
	// true
	// false
	// false
}

func ExampleGraph_RemoveLink() {
	graph := NewGraph(false, true, false)

	graph.AddLink("a", "b", 2)

	fmt.Println(graph.RemoveLink("a", "b"))
	fmt.Println(graph.RemoveLink("a", "b"))
	fmt.Println(graph.RemoveLink("a", "c"))

	// Output:
	// true
	// false
	// false
}

func ExampleGraph_OutgoingLinksCount() {
	graph := createGraph()

	fmt.Println(graph.OutgoingLinksCount("alpha"))
	fmt.Println(graph.OutgoingLinksCount("2"))
	fmt.Println(graph.OutgoingLinksCount("3"))

	// Output:
	// 1
	// 3
	// 2
}

func ExampleGraph_IncomingLinksCount() {
	graph := createGraph()

	fmt.Println(graph.IncomingLinksCount("alpha"))
	fmt.Println(graph.IncomingLinksCount("2"))
	fmt.Println(graph.IncomingLinksCount("3"))

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
