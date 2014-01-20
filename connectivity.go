package grago

// Return a slice of the names of all nodes to which a path
// exists from the node provided as a parameter.
func (g *Graph) ReachableNodes(node string) []string {
	result := []string{}
	levels := g.BFS(node)
	
	if len(levels) > 1 {
		for _, nodeList := range levels[1:] {
			result = append(result, nodeList...)
		}
	}
	
	return result
}

// Return a slice of slices where are given the node names
// in each separate connected component.
func (g *Graph) ConnectedComponents() [][]string {
	result := [][]string{}

	marked := make(map[string]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}
	
	for _, node := range g.Nodes() {
		if !marked[node] {
			nodeList := append(g.ReachableNodes(node), node)
			for _, otherNodes := range nodeList {
				marked[otherNodes] = true
			}
			result = append(result, nodeList)
		}
	}
	
	return result
}
