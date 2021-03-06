// This package provides a basic graph structure, along with
// some of the popular algorithms in graph theory.
package grago

import (
	"regexp"
	"strconv"
	"strings"
	"fmt"
)

type stringer string

func (s stringer) String() string {
	return string(s)
} 

type Link struct {
	Start   fmt.Stringer
	End     fmt.Stringer
	Weighed bool
	Weight  int
}

func NewLink(start fmt.Stringer, end fmt.Stringer, weighed bool, weight int) *Link {
	var l *Link = new(Link)

	l.Start = start
	l.End = end
	l.Weighed = weighed
	l.Weight = weight

	return l
}

// A string representation of Link used for output.
func (l Link) String() string {
	if l.Weighed {
		return l.Start.String() + "-(" + strconv.Itoa(l.Weight) + ")->" + l.End.String()
	} else {
		return l.Start.String() + "--" + l.End.String()
	}
}

type Node struct {
	// A container for all the nodes, that this node has an
	// outgoing connection to.
	Adjacent map[fmt.Stringer]int
}

func NewNode() *Node {
	var n *Node = new(Node)

	n.Adjacent = make(map[fmt.Stringer]int)

	return n
}

// Returns a slice of the nodes this node has an edge to.
func (n *Node) AdjacentNodes() []fmt.Stringer {
	result := []fmt.Stringer{}

	for node := range n.Adjacent {
		result = append(result, node)
	}

	return result
}

type Graph struct {
	// Whether the graph is oriented or not takes effect on the
	// visual representation as well as edge addition to it.
	Oriented bool

	// If the graph is weighed, the weights will be shown on the
	// exported visualization, otherwise an edge between two nodes
	// will be represented by a weight of 1
	Weighed bool

	// Whether or not the graph has negative weights takes effect
	// on which algorithms to use in some cases like shorted path
	// finding.
	HasNegativeWeights bool

	// A container for all the nodes in the graph.
	nodes map[fmt.Stringer]*Node
}

func NewGraph(oriented bool, weighed bool, hasNegativeWeights bool) *Graph {
	var g *Graph = new(Graph)

	g.Oriented = oriented
	g.Weighed = weighed
	g.HasNegativeWeights = hasNegativeWeights
	g.nodes = make(map[fmt.Stringer]*Node)

	return g
}

// Creates a graph, read from a string with an expected
// input format such as:
// <oriented> <weighed> <hasNegativeWeights> - booleans
// <node> - for adding a node
// <node> -- <node> [<weight>] - for adding a link
// If 'reversed' is true, the graph will be constructed
// with reversed edges.
func ReadGraph(in string, reversed bool) *Graph {
	var g *Graph = new(Graph)

	g.nodes = make(map[fmt.Stringer]*Node)

	lines := strings.Split(in, "\n")
	attributes := strings.Split(lines[0], " ")
	g.Oriented, _ = strconv.ParseBool(attributes[0])
	g.Weighed, _ = strconv.ParseBool(attributes[1])
	g.HasNegativeWeights, _ = strconv.ParseBool(attributes[2])

	linkRegexString := `(.+) -- (.+)`
	if g.Weighed {
		linkRegexString = `(.+) -- (.+) (\d+)`
	}

	linkRegex := regexp.MustCompile(linkRegexString)

	for _, line := range lines[1:] {
		linkMatch := linkRegex.FindStringSubmatch(line)

		if len(linkMatch) != 0 { //this line describes a link
			if g.Weighed {
				weight, _ := strconv.Atoi(linkMatch[3])
				if reversed {
					g.AddLink(stringer(linkMatch[2]), stringer(linkMatch[1]), weight)
				} else {
					g.AddLink(stringer(linkMatch[1]), stringer(linkMatch[2]), weight)
				}
			} else {
				if reversed {
					g.AddLink(stringer(linkMatch[2]), stringer(linkMatch[1]), 1)
				} else {
					g.AddLink(stringer(linkMatch[1]), stringer(linkMatch[2]), 1)
				}
			}
		} else { //this line describes a node
			g.AddNode(stringer(line))
		}
	}

	return g
}

// A string representation of Graph, with an output format
// like the input format:
// <oriented> <weighed> <hasNegativeWeights> - booleans
// <node> - all nodes in the graph each on a separate line
// <node> -- <node> [<weight>] - for all the links
func (g *Graph) String() string {
	result := ""

	result += strconv.FormatBool(g.Oriented)
	result += " " + strconv.FormatBool(g.Weighed)
	result += " " + strconv.FormatBool(g.HasNegativeWeights) + "\n"

	for node := range g.nodes {
		result += node.String() + "\n"
	}

	for node, adjacents := range g.nodes {
		adjacentList := adjacents.AdjacentNodes()
		for _, adjNode := range adjacentList {
			if g.Weighed {
				result += node.String() + " -- " + adjNode.String() + " " + strconv.Itoa(adjacents.Adjacent[adjNode]) + "\n"
			} else {
				result += node.String() + " -- " + adjNode.String() + "\n"
			}
		}
	}

	return result
}

// Tries to add a node to the graph and returns true
// if successful, otherwise returns false if the
// node already exists.
func (g *Graph) AddNode(node fmt.Stringer) bool {
	if _, exists := g.nodes[node]; exists { //check if the node is already in the graph
		return false
	} else {
		g.nodes[node] = NewNode()
		return true
	}
}

// Tries to add a link between two nodes and returns true
// if successful, otherwise returns false if such a link
// already exists.
// 
// Note: If the graph isn't oriented adding a link from A to B
// effectively adds a link from B to A.
func (g *Graph) AddLink(startNode fmt.Stringer, endNode fmt.Stringer, weight int) bool {
	g.AddNode(startNode)
	g.AddNode(endNode)
	if _, exists := g.nodes[startNode].Adjacent[endNode]; exists { //check if the link is already in the graph
		return false
	} else {
		g.nodes[startNode].Adjacent[endNode] = weight
		if !g.Oriented {
			g.nodes[endNode].Adjacent[startNode] = weight
		}
		return true
	}
}

// Tries to remove the node from the graph and if
// successful removes all links between it and
// other nodes and returns true, otherwise if the
// node doesn't exist, returns false.
func (g *Graph) RemoveNode(node fmt.Stringer) bool {
	if _, exists := g.nodes[node]; exists { //check if the node is already in the graph
		delete(g.nodes, node)
		for _, otherNode := range g.nodes {
			if _, existsNode := otherNode.Adjacent[node]; existsNode {
				//if there is another node pointing to the one we removed, delete the link
				delete(otherNode.Adjacent, node)
			}
		}
		return true
	} else {
		return false
	}
}

// Tries to remove the link from the graph and if
// successful returns true, otherwise if the link
// doesn't exist, returns false.
// 
// Note: If the graph isn't oriented removing the link from A to B
// effectively removes the link from B to A.
func (g *Graph) RemoveLink(startNode fmt.Stringer, endNode fmt.Stringer) bool {
	if _, existsNode := g.nodes[startNode]; existsNode { //check if the start node exists
		if _, existsLink := g.nodes[startNode].Adjacent[endNode]; existsLink { //check if the link exists
			delete(g.nodes[startNode].Adjacent, endNode)
			if !g.Oriented {
				delete(g.nodes[endNode].Adjacent, startNode)
			}
			return true
		} else {
			return false
		}
	}
	return false
}

// Returns the count of the links which have as a starting
// node the one specified as a parameter.
// 
// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g *Graph) OutgoingLinksCount(node fmt.Stringer) int {
	return len(g.nodes[node].Adjacent)
}

// Returns the count of the links which have as an ending
// node the one specified as a parameter.
// 
// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g *Graph) IncomingLinksCount(node fmt.Stringer) int {
	result := 0

	for _, otherNode := range g.nodes {
		if _, exists := otherNode.Adjacent[node]; exists { //check if there is an incoming link from another node
			result++
		}
	}

	return result
}

// Returns a slice with all the nodes in the graph.
func (g *Graph) Nodes() []fmt.Stringer {
	result := []fmt.Stringer{}

	for node := range g.nodes {
		result = append(result, node)
	}

	return result
}

// Returns a slice with the all the links
// in the graph.
func (g *Graph) Links() []Link {
	result := []Link{}

	for _, startNode := range g.Nodes() {
		for endNode, weight := range g.nodes[startNode].Adjacent {
			result = append(result, *NewLink(startNode, endNode, g.Weighed, weight))
		}
	}

	return result
}
