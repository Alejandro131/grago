package grago

import (
	"regexp"
	"strings"
	"strconv"
)

type Link struct {
	Start string
	End string
	Weighed bool
	Weight int
}

func NewLink(start string, end string, weighed bool, weight int) *Link {
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
		return l.Start + "-(" + strconv.Itoa(l.Weight) + ")->" + l.End
	} else {
		return l.Start + "--" + l.End
	}
}

type Node struct {
	// A container for all the nodes, that this node has an
	// outgoing connection to.
	Adjacent map[string]int
}

func NewNode() *Node {
	var n *Node = new(Node)
	
	n.Adjacent = make(map[string]int)
	
	return n
}

// Returns a slice of the names of the nodes this node
// has an edge to.
func (n *Node) AdjacentNodes() []string {
	result := []string{}
	
	for name := range n.Adjacent {
		result = append(result, name)
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

	// A container for all the nodes in the graph, accessible by
	// their given names.
	nodes map[string]*Node
}

func NewGraph(oriented bool, weighed bool, hasNegativeWeights bool) *Graph {
	var g *Graph = new(Graph)
	
	g.Oriented = oriented
	g.Weighed = weighed
	g.HasNegativeWeights = hasNegativeWeights
	g.nodes = make(map[string]*Node)
	
	return g
}

// Creates a graph, read from a string with an expected
// input format such as:
// <oriented> <weighed> <hasNegativeWeights> - booleans
// <node> - for adding a node
// <node> -- <node> [<weight>] - for adding a link
func ReadGraph(in string) *Graph {
	var g *Graph = new(Graph)
	
	g.nodes = make(map[string]*Node)
	
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
				g.AddLink(linkMatch[1], linkMatch[2], weight)
			} else {
				g.AddLink(linkMatch[1], linkMatch[2], 1)
			}
		} else { //this line describes a node
			g.AddNode(line)
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
		result += node + "\n"
	}
	
	for node, adjacents := range g.nodes {
		adjacentList := adjacents.AdjacentNodes()
		for _, adjNode := range adjacentList {
			if g.Weighed {
				result += node + " -- " + adjNode + " " + strconv.Itoa(adjacents.Adjacent[adjNode]) + "\n"
			} else {
				result += node + " -- " + adjNode + "\n"
			}
		}
	}
	
	return result
}

// Tries to add a node to the graph and returns true
// if successful, otherwise returns false if the
// node already exists.
func (g *Graph) AddNode(node string) bool {
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

// Note: If the graph isn't oriented adding a link from A to B
// effectively adds a link from B to A.
func (g *Graph) AddLink(startNode string, endNode string, weight int) bool {
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
func (g *Graph) RemoveNode(node string) bool {
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

// Note: If the graph isn't oriented removing the link from A to B
// effectively removes the link from B to A.
func (g *Graph) RemoveLink(startNode string, endNode string) bool {
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

// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g *Graph) OutgoingLinksCount(node string) int {
	return len(g.nodes[node].Adjacent)
}

// Returns the count of the links which have as an ending
// node the one specified as a parameter.

// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g *Graph) IncomingLinksCount(node string) int {
	result := 0
	
	for _, otherNode := range g.nodes {
		if _, exists := otherNode.Adjacent[node]; exists { //check if there is an incoming link from another node
			result++
		}
	}
	
	return result
}

// Returns a slice with the names of all the nodes
// in the graph.
func (g *Graph) Nodes() []string {
	result := []string{}
	
	for name := range g.nodes {
		result = append(result, name)
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
