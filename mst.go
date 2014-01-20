package grago

// Interface for the priority queue for Link
func (l *Link) Less(other interface{}) bool {
	return l.Weight < other.(*Link).Weight
}

// Returns a slice of the links constructing the
// minimal spanning tree for the given graph,
// according to Kruskal's algorithm.
func (g *Graph) MST() []Link {
	result := []Link{}

	linkQueue := NewPriorityQueue(0)
	for _, startNode := range g.Nodes() {
		for endNode, weight := range g.nodes[startNode].Adjacent {
			linkQueue.Enqueue(NewLink(startNode, endNode, g.Weighed, weight))
		}
	}
	
	sets := make(map[string]int)
	setCount := 0
	for _, node := range g.Nodes() {
		sets[node] = -1
	}
	
	for !linkQueue.IsEmpty() {
		link := *(linkQueue.Dequeue().(*Link))
		if sets[link.Start] == -1 && sets[link.End] == -1 {
			sets[link.Start] = setCount
			sets[link.End] = setCount
			setCount++
			result = append(result, link)
		} else if sets[link.Start] == -1 || sets[link.End] == -1 {
			if sets[link.Start] == -1 {
				sets[link.Start] = sets[link.End]
			} else {
				sets[link.End] = sets[link.Start]
			}
			result = append(result, link)
		} else if sets[link.Start] != sets[link.End] {
			for node, set := range sets {
				if set == sets[link.End] {
					sets[node] = sets[link.Start]
				}
			}
			result = append(result, link)
		}
	}
	
	return result
}
