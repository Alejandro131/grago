package grago

import (
	"strconv"
	"fmt"
)

// Export the graph structure in a format specified
// by the dot language used in graphviz, which can
// be loaded by the software and visualized.
// 'highlights' is a slice of links representing edges
// of the graph you would like to outline, for example
// the results of a DFS search and they will be colored red.
// 'ordered' is a boolean indicating whether or not the index
// of the links would be shown next to their name which is
// useful for distinguishing DFS from MST for example.
// 'clusters' is a 2d array of fmt.Stringers representing nodes
// that should be ordered in separate groups visually, if it's
// empty, the graph will be displayed as is.
func (g *Graph) Export(highlights []Link, ordered bool, clusters [][]fmt.Stringer) string {
	result := ``

	if g.Oriented {
		result += `digraph {`
	} else {
		result += `graph {`
	}

	graph := NewGraph(g.Oriented, g.Weighed, g.HasNegativeWeights)

	if len(clusters) != 0 {
		for clusterID, cluster := range clusters {
			result += `subgraph cluster` + strconv.Itoa(clusterID) + ` {`

			for _, node := range cluster {
				result += `"` + node.String() + `" `
				graph.AddNode(node)
			}

			result += `}`
		}
	}

	if len(highlights) != 0 {
		for linkID, link := range highlights {
			graph.AddLink(link.Start, link.End, link.Weight)

			result += `"` + link.Start.String() + `"`
			if g.Oriented {
				result += `->`
			} else {
				result += `--`
			}
			result += `"` + link.End.String() + `"`

			result += ` [fontcolor=red color=red `

			if g.Weighed || ordered {
				result += `label="`

				if g.Weighed && ordered {
					result += strconv.Itoa(link.Weight) + ` (` + strconv.Itoa(linkID+1) + `)`
				} else if g.Weighed {
					result += strconv.Itoa(link.Weight)
				} else {
					result += `(` + strconv.Itoa(linkID+1) + `)`
				}

				result += `"`
			}

			result += `]; `
		}
	}

	// Add all nodes that aren't present from the clusters and highlights
	for _, node := range g.Nodes() {
		if graph.AddNode(node) {
			result += `"` + node.String() + `" `
		}
	}

	// Add all links that aren't present from the highlights
	for _, link := range g.Links() {
		if graph.AddLink(link.Start, link.End, link.Weight) {
			result += `"` + link.Start.String() + `"`
			if g.Oriented {
				result += `->`
			} else {
				result += `--`
			}
			result += `"` + link.End.String() + `"`

			if g.Weighed {
				result += ` [label="` + strconv.Itoa(link.Weight) + `"]; `
			}
		}
	}

	return result + `}`
}
