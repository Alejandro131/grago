package grago

// Returns a matrix of node names, where you can find
// the minimum distance between two nodes, given their names.
func (g *Graph) Floyd() map[string]map[string]int {
}

// Returns a map bound by node name, containing the minimum
// distance between the given node and all of the other nodes.
// Internally it chooses between Dijkstra's and Ford Bellman's
// algorithms depending on whether the graph has negative weights
// or not.
func (g *Graph) MinPaths(node string) map[string]int {
}

// Dijkstra's implementation of the shortest path algorithm.
func (g *Graph) dijkstra(node string) map[string]int {
}

// Ford Bellman's implementation of the shortest path algorithm.
func (g *Graph) fordBellman(node string) map[string]int {
}
