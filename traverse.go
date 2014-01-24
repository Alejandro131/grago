package grago

// Traverses the graph in a manner of breadth first search
// and returns a slice of slices of node names, representing
// the layers of nodes during the search, with the first layer
// with index 0 containing the initial node.
func (g *Graph) BFS(start string) [][]string {
	result := [][]string{}
	
	marked := make(map[string]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}
	marked[start] = true
	
	currentLevel := []string{start}
	result = append(result, []string{start})
	for len(currentLevel) != 0 {
		nextLevel := []string{}
		for _, node := range currentLevel {
			for adjacentNode := range g.nodes[node].Adjacent {
				if !marked[adjacentNode] {
					marked[adjacentNode] = true
					nextLevel = append(nextLevel, adjacentNode)
				}
			}
		}

		if len(nextLevel) != 0 {
			result = append(result, nextLevel)
		}
		currentLevel = currentLevel[:0] //clear the level and prepare it for the next iteration
		currentLevel = append(currentLevel, nextLevel...)
	}
	
	return result
}

// Helper function for recursive calling of dfs
func (g *Graph) dfs(start string, marked *map[string]bool, links *[]Link) {
	for node := range g.nodes[start].Adjacent {
		if !(*marked)[node] {
			*links = append(*links, *NewLink(start, node, g.Weighed, g.nodes[start].Adjacent[node]))
			(*marked)[node] = true
			g.dfs(node, marked, links)
		}
	}
}

// Traverses the graph in a manner of depth first search
// and returns a slice of links through which it goes
// during the search.
func (g *Graph) DFS(start string) []Link {
	result := []Link{}

	marked := make(map[string]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}
	marked[start] = true
	g.dfs(start, &marked, &result)
	
	return result
}
