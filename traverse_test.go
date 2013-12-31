package grago

import (
	"testing"
	"fmt"
)

func createGraph() Graph {
	graph := NewGraph(false, true, false)
	graph.AddNode("alpha")
	graph.AddLink("2", "3", 2)
	graph.AddLink("2", "4", 5)
	graph.AddLink("3", "5", 8)
	graph.AddLink("5", "4", 10)
	return graph
}

func ExampleBFS_Graph() {	
	fmt.Println(createGraph().BFS("2"))
	
	// Output:
	// [[2] [3 4] [5]]
}

func ExampleDFS_Graph() {
	fmt.Println(createGraph().DFS("2"))
	
	// Output:
	// [2-(2)->3 3-(8)->5 5-(10)->4]
}