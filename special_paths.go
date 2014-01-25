package grago

// Returns a slice of links representing an Eulearian
// path. If no such path exists, returns an empty slice.
// This algorithm works for undirected graphs for now.

// Note: An Eulerian path is such that it traverses
// through every edge exactly once.
func (g Graph) EulerPath() []Link {
	var start string

	// The graph must be connected in order for an
	// eulerian path to exist. Even if it has many
	// connected components if all of them but 1 have
	// one node each, an eulerian path could exist.

	if len(g.ConnectedComponents()) > 1 {
		componentCount := 0

		for _, component := range g.ConnectedComponents() {
			if len(component) > 1 {
				start = component[0]
				componentCount++
			}
		}

		if componentCount > 1 {
			return []Link{}
		}
	}

	oddNodes := 0

	for _, node := range g.Nodes() {
		if g.OutgoingLinksCount(node)%2 == 1 {
			oddNodes++
			start = node
		}
	}

	if oddNodes > 2 {
		return []Link{}
	}

	// Create a copy of the graph as we will manipulate it.
	graph := ReadGraph(g.String())

	stack := []string{start}
	pathNodes := []string{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]

		if _, exists := graph.nodes[node]; exists && len(graph.nodes[node].AdjacentNodes()) != 0 {
			nextNode := graph.nodes[node].AdjacentNodes()[0]
			stack = append(stack, nextNode)

			graph.RemoveLink(node, nextNode)
		} else {
			pathNodes = append(pathNodes, node)
			stack = stack[:len(stack)-1]
		}
	}

	result := []Link{}

	if len(pathNodes) > 1 {
		for index := 0; index < len(pathNodes)-1; index++ {
			start, end := pathNodes[index], pathNodes[index+1]
			result = append(result, *NewLink(start, end, g.Weighed, g.nodes[start].Adjacent[end]))
		}
	}

	return result
}

// Helper function for the hamilton path algorithm, traversing
// nodes for a possible solution to the problem.
func (g *Graph) hamiltonDFS(currentNode string, found *bool, marked *map[string]bool, path *[]Link, nodeCount int) {
	if len(*path) == nodeCount-1 {
		*found = true
		return
	}

	for _, node := range g.nodes[currentNode].AdjacentNodes() {
		if !(*marked)[node] {
			(*marked)[node] = true
			*path = append(*path, *NewLink(currentNode, node, g.Weighed, g.nodes[currentNode].Adjacent[node]))

			g.hamiltonDFS(node, found, marked, path, nodeCount)

			if *found {
				return
			}

			*path = (*path)[:len(*path)-1]
			(*marked)[node] = false
		}
	}
}

// Returns a slice of links representing an Hamiltonian
// path. If no such path exists, returns an empty slice.

// Note: A Hamiltonian path is such that it traverses
// through every vertex exactly once.
func (g *Graph) HamiltonPath() []Link {
	if len(g.ConnectedComponents()) > 1 {
		return []Link{}
	}

	pathFound := false
	result := []Link{}
	marked := make(map[string]bool)
	nodeCount := len(g.Nodes())

	for _, node := range g.Nodes() {
		marked[node] = true
		g.hamiltonDFS(node, &pathFound, &marked, &result, nodeCount)
		marked[node] = false

		if pathFound {
			break
		}
	}

	if !pathFound {
		return []Link{}
	}

	return result
}
