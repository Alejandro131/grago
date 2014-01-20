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
func (g *Graph) MinPaths(start string) map[string]int {
	if g.HasNegativeWeights {
		return g.fordBellman(start)
	} else {
		return g.dijkstra(start)
	}
}

// Dijkstra's implementation of the shortest path algorithm.
func (g *Graph) dijkstra(start string) map[string]int {
	result := make(map[string]int)
	
	result[start] = 0
	marked := make(map[string]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}
	marked[start] = true
	
	return result
}

// Ford Bellman's implementation of the shortest path algorithm.
func (g *Graph) fordBellman(start string) map[string]int {
}
