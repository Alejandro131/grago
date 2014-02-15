package grago

import "fmt"

// Return a slice of all nodes to which a path
// exists from the node provided as a parameter.
func (g *Graph) ReachableNodes(node fmt.Stringer) []fmt.Stringer {
	result := []fmt.Stringer{}
	levels := g.BFS(node)

	if len(levels) > 1 {
		for _, nodeList := range levels[1:] {
			result = append(result, nodeList...)
		}
	}

	return result
}

// Helper function implementing the traversing part of
// Kosaraju's algorithm for strongly connected components.
func (g *Graph) connectedComponentsDFS(currentNode fmt.Stringer, marked *map[fmt.Stringer]bool, stack *[]fmt.Stringer) {
	(*marked)[currentNode] = true

	for node := range g.nodes[currentNode].Adjacent {
		if _, exists := (*marked)[node]; !exists {
			g.connectedComponentsDFS(node, marked, stack)
		}
	}

	*stack = append(*stack, currentNode)
}

// Return a slice of slices where are given the nodes
// in each separate strongly connected component.
func (g *Graph) ConnectedComponents() [][]fmt.Stringer {
	result := [][]fmt.Stringer{}

	stack := []fmt.Stringer{}
	marked := make(map[fmt.Stringer]bool)
	nodeCount := len(g.Nodes())

	for len(marked) != nodeCount {
		for _, node := range g.Nodes() {
			if _, exists := marked[node]; !exists {
				g.connectedComponentsDFS(node, &marked, &stack)
			}
		}
	}

	reverseGraph := ReadGraph(g.String(), true)

	for len(stack) != 0 {
		currentNode := stack[len(stack)-1]
		if _, exists := marked[currentNode]; exists {
			set := reverseGraph.ReachableNodes(currentNode)
			set = append(set, currentNode)

			for _, node := range set {
				delete(marked, node)
				reverseGraph.RemoveNode(node)
			}

			result = append(result, set)
		}
		stack = stack[:len(stack)-1]
	}

	return result
}
