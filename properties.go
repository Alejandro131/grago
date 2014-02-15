package grago

import "fmt"

// Helper function for HasCycle doing a DFS and
// searching for already marked nodes indicating
// a present cycle.
func (g *Graph) cycleDFS(start fmt.Stringer, marked *map[fmt.Stringer]bool, hasCycle *bool) {
	for node := range g.nodes[start].Adjacent {
		if !(*marked)[node] {
			(*marked)[node] = true

			g.cycleDFS(node, marked, hasCycle)
			if *hasCycle {
				break
			}

			(*marked)[node] = false
		} else {
			*hasCycle = true
			break
		}
	}
}

// Returns a boolean value of whether or not the
// graph contains atleast one cycle, which means
// that some edges form a closed circle.
func (g *Graph) HasCycle() bool {
	hasCycle := false

	marked := make(map[fmt.Stringer]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}

	for _, node := range g.Nodes() {
		marked[node] = true
		g.cycleDFS(node, &marked, &hasCycle)
		marked[node] = false

		if hasCycle {
			return hasCycle
		}
	}

	return hasCycle
}

// Helper function for IsPlanar, determining
// whether the given graph is the full bipartite
// graph K_3_3 or not.
func (g *Graph) isK33() bool {
	return len(g.Nodes()) == 6 && len(g.Links()) == 9*2 && g.IsBipartite()
}

// Helper function for IsPlanar, determining whether
// the given graph is the full K_5 graph or not
func (g *Graph) isK5() bool {
	return len(g.Nodes()) == 5 && len(g.Links()) == 10*2
}

// Helper function for IsPlanar, merging two nodes
// into one.
func (g *Graph) contractLink(start fmt.Stringer, end fmt.Stringer) {
	g.RemoveLink(start, end)

	// Make every incoming connection of end node an
	// incoming connection of start node.

	for _, otherNode := range g.nodes {
		if _, exists := otherNode.Adjacent[end]; exists {
			otherNode.Adjacent[start] = otherNode.Adjacent[end]
		}
	}

	// Make every outgoing connection from end node an
	// outgoing connection from start node.

	for otherNode, weight := range g.nodes[end].Adjacent {
		g.nodes[start].Adjacent[otherNode] = weight
	}

	g.RemoveNode(end)
}

// Returns a boolean value of whether or not the
// graph is planar, which means that it can be drawn
// on a piece of paper with no two edges crossing.
func (g *Graph) IsPlanar() bool {
	if len(g.Nodes()) <= 5 && len(g.Links()) < 10 {
		return true
	}

	if g.isK33() || g.isK5() {
		return false
	}

	for _, node := range g.Nodes() {
		graph := ReadGraph(g.String(), false)
		graph.RemoveNode(node)
		if !graph.IsPlanar() {
			return false
		}
	}

	for _, link := range g.Links() {
		graph := ReadGraph(g.String(), false)
		graph.RemoveLink(link.Start, link.End)
		if !graph.IsPlanar() {
			return false
		}
	}

	for _, link := range g.Links() {
		graph := ReadGraph(g.String(), false)
		graph.contractLink(link.Start, link.End)
		if !graph.IsPlanar() {
			return false
		}
	}

	return true
}

// Helper function for the IsBipartite function,
// creating a subgraph with only the nodes provided
// and links between them.
func (g *Graph) subGraph(nodes []fmt.Stringer) *Graph {
	toDelete := make(map[fmt.Stringer]bool)

	for _, node := range g.Nodes() {
		toDelete[node] = true
	}

	for _, node := range nodes {
		delete(toDelete, node)
	}

	graph := ReadGraph(g.String(), false)

	for node := range toDelete {
		graph.RemoveNode(node)
	}

	return graph
}

// Returns a boolean value of whether or not the
// graph is bipartite, which means that it's vertices
// can be split into two groups where there aren't any
// edges between vertices of the same group.
func (g *Graph) IsBipartite() bool {
	components := g.ConnectedComponents()

	if len(components) > 1 {
		for _, nodes := range components {
			if !((g.subGraph(nodes)).IsBipartite()) {
				return false
			}
		}
		return true
	}

	if len(g.Nodes()) <= 2 {
		return true
	}

	sets := make(map[fmt.Stringer]int)
	for _, node := range g.Nodes() {
		sets[node] = -1
	}

	currentSet := 0
	start := g.Nodes()[0]
	sets[start] = currentSet

	currentLevel := []fmt.Stringer{start}
	for len(currentLevel) != 0 {
		nextLevel := []fmt.Stringer{}

		for _, node := range currentLevel {
			for adjacentNode := range g.nodes[node].Adjacent {
				if sets[adjacentNode] == -1 {
					sets[adjacentNode] = (currentSet + 1) % 2
					nextLevel = append(nextLevel, adjacentNode)
				} else if sets[adjacentNode] == currentSet {
					return false
				}
			}
		}

		currentLevel = currentLevel[:0]
		currentLevel = append(currentLevel, nextLevel...)

		currentSet = (currentSet + 1) % 2
	}

	return true
}
