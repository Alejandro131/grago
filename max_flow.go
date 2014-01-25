package grago

// Helper function for the max flow algorithm which
// finds a path through which we can establish a
// flow from the source to the sink
func flowDFS(startNode string, flowPath *[]string, newFlow *bool, marked *map[string]bool, sink string, adjacency *map[string]map[string]int, flow *map[string]map[string]int) {
	if *newFlow {
		return
	}

	if startNode == sink {
		*newFlow = true
		updateFlow(flowPath, adjacency, flow)
	} else {
		for node := range *adjacency {
			if !(*marked)[node] && (*adjacency)[startNode][node] > 0 {
				(*marked)[node] = true

				*flowPath = append(*flowPath, node)
				flowDFS(node, flowPath, newFlow, marked, sink, adjacency, flow)
				*flowPath = (*flowPath)[:len(*flowPath)-1]

				if *newFlow {
					return
				}
			}
		}

	}
}

// Helper function for the max flow algorithm which
// applies the flow from the source to the sink and
// adds a blocking flow backwards
func updateFlow(flowPath *[]string, adjacency *map[string]map[string]int, flow *map[string]map[string]int) {
	increaseFlow := 0

	if len(*flowPath) >= 2 {
		increaseFlow = (*adjacency)[(*flowPath)[0]][(*flowPath)[1]]
	}

	for index := 1; index < len(*flowPath)-1; index++ {
		if increaseFlow > (*adjacency)[(*flowPath)[index]][(*flowPath)[index+1]] {
			increaseFlow = (*adjacency)[(*flowPath)[index]][(*flowPath)[index+1]]
		}
	}

	for index := 0; index < len(*flowPath)-1; index++ {
		(*flow)[(*flowPath)[index]][(*flowPath)[index+1]] += increaseFlow
		(*flow)[(*flowPath)[index+1]][(*flowPath)[index]] -= increaseFlow
		(*adjacency)[(*flowPath)[index]][(*flowPath)[index+1]] -= increaseFlow
		(*adjacency)[(*flowPath)[index+1]][(*flowPath)[index]] += increaseFlow
	}
}

// Returns the maximum ammount that can flow from
// the source to the sink for 1 period of time
// according to the min-cut max-flow algorithm.
func (g *Graph) MaxFlow(source string, sink string) int {
	flow := make(map[string]map[string]int)
	adjacency := make(map[string]map[string]int)

	for _, startNode := range g.Nodes() {
		adjacency[startNode] = make(map[string]int)
		flow[startNode] = make(map[string]int)

		for _, endNode := range g.Nodes() {
			flow[startNode][endNode] = 0
			adjacency[startNode][endNode] = 0
		}
	}

	for _, startNode := range g.Nodes() {
		for endNode, weight := range g.nodes[startNode].Adjacent {
			adjacency[startNode][endNode] = weight
		}
	}

	marked := make(map[string]bool)
	for _, node := range g.Nodes() {
		marked[node] = false
	}
	newFlow := true

	for newFlow {
		for _, node := range g.Nodes() {
			marked[node] = false
		}
		newFlow = false
		marked[source] = true
		flowPath := []string{source}

		flowDFS(source, &flowPath, &newFlow, &marked, sink, &adjacency, &flow)
	}

	result := 0

	for _, node := range g.Nodes() {
		result += flow[node][sink]
	}

	return result
}
