package grago

import (
	"testing"
	"fmt"
)

func createGraph() Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "alpha", 2)
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleString_Link() {
	fmt.Println(NewLink("1", "4", 23))
	
	// Output:
	// 1-(23)->4
}

func ExampleAdjacentNodes_Node() {
	node := NewNode()
	node.Adjacent["2"] = 1
	node.Adjacent["5"] = 234
	
	fmt.Println(node.AdjacentNodes())
	
	// Output:
	// [2 5]
}

func ExampleAddNode_Graph() {
	graph := NewGraph(false, true, false)
	
	fmt.Println(graph.AddNode("a"))
	fmt.Println(graph.AddNode("a"))
	fmt.Println(graph.AddNode("b"))
	
	// Output:
	// true
	// false
	// true
}

func ExampleAddLink_Graph() {
	graph := NewGraph(false, true, false)
	
	fmt.Println(graph.AddLink("a", "b", 2))
	fmt.Println(graph.AddLink("a", "b", 2))
	fmt.Println(graph.AddLink("a", "b", 5))
	
	// Output:
	// true
	// false
	// true
}

func ExampleRemoveNode_Graph() {
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

func ExampleRemoveLink_Graph() {
	graph := NewGraph(false, true, false)
	
	graph.AddLink("a", "b", 2)
	
	fmt.Println(graph.RemoveLink("a", "b", 2))
	fmt.Println(graph.RemoveLink("a", "b", 2))
	fmt.Println(graph.RemoveLink("a", "c", 2))
	
	// Output:
	// true
	// false
	// false
}

func ExampleRemoveLinks_Graph() {
	graph := NewGraph(false, true, false)
	
	graph.AddLink("a", "b", 2)
	
	fmt.Println(graph.RemoveLinks("a", "b"))
	fmt.Println(graph.RemoveLinks("a", "b"))
	fmt.Println(graph.RemoveLinks("a", "c"))
	
	// Output:
	// true
	// false
	// false
}

func ExampleOutgoingLinksCount_Graph() {
	graph := createGraph()
	
	fmt.Println(graph.OutgoingLinksCount("alpha"))
	fmt.Println(graph.OutgoingLinksCount("2"))
	fmt.Println(graph.OutgoingLinksCount("3"))
	
	// Output:
	// 1
	// 3
	// 2
}

func ExampleIncomingLinksCount_Graph() {
	graph := createGraph()
	
	fmt.Println(graph.IncomingLinksCount("alpha"))
	fmt.Println(graph.IncomingLinksCount("2"))
	fmt.Println(graph.IncomingLinksCount("3"))
	
	// Output:
	// 1
	// 3
	// 2
}

func ExampleNodes_Graph() {
	fmt.Println(createGraph().Nodes())
	
	// Output:
	// [alpha 2 3 4 5]
}