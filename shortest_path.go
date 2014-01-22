package grago

// Returns a matrix of node names, where you can find
// the minimum distance between two nodes, given their names.
func (g *Graph) Floyd() map[string]map[string]int {
	result := make(map[string]map[string]int)
	
	nodeList := g.Nodes()
	for _, node := range nodeList {
		result[node] = make(map[string]int)
		result[node][node] = 0
		for endNode, weight := range g.nodes[node].Adjacent {
			result[node][endNode] = weight
		}
	}
	
	for _, k := range nodeList {
		for _, i := range nodeList {
			for _, j := range nodeList {
				_, existsI := result[i][k]
				_, existsJ := result[k][j]
				if existsI && existsJ {
					if _, exists := result[i][j]; !exists || result[i][j] > result[i][k] + result[k][j] {
						result[i][j] = result[i][k] + result[k][j]
					}
				}
			}
		}
	}
	
	return result
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
	distanceQueue := NewPriorityQueue(0)
	distanceQueue.Enqueue(NewLink(start, start, g.Weighed, 0))
	
	for !distanceQueue.IsEmpty() {
		link := *(distanceQueue.Dequeue().(*Link))
		for endNode, weight := range g.nodes[link.End].Adjacent {
			if _, exists := result[endNode]; !exists || result[endNode] > result[link.End] + weight {
				result[endNode] = result[link.End] + weight
				distanceQueue.Enqueue(NewLink(link.End, endNode, g.Weighed, result[link.End] + weight))
			}
		}
	}
	
	return result
}

// Ford Bellman's implementation of the shortest path algorithm.
func (g *Graph) fordBellman(start string) map[string]int {
	result := make(map[string]int)
	
	result[start] = 0
	links := g.Links()
	nodeCount := len(g.Nodes())
	
	for i := 0; i < nodeCount - 1; i++ {
		for _, link := range links {
			if _, startExists := result[link.Start]; startExists {
				if _, endExists := result[link.End]; !endExists || result[link.Start] + link.Weight < result[link.End] {
					result[link.End] = result[link.Start] + link.Weight
				}
			}
		}
	}
	
	return result
}
