package grago

// Traverses the graph in a manner of breadth first search
// and returns a slice of slices of node names, representing
// the layers of nodes during the search, with the first layer
// with index 0 containing the initial node.
func (g Graph) BFS(node string) [][]string {
}

// Traverses the graph in a manner of depth first search
// and returns a slice of links through which it goes
// during the search.
func (g Graph) DFS(node string) []Link {
}