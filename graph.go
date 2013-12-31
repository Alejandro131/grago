package grago

type Link struct {
	// The name of the start node
	Start string
	
	// The name of the end node
	End string
	
	Weight int
}

func NewLink(start string, end string, weight int) {
}

type Node struct {
	// A container for all the nodes, that this node has an
	// outgoing connection to.
	Adjacent map[string]int
}

func NewNode() Node {
}

// Returns a slice of the names of the nodes this node
// has an edge to.
func (n Node) AdjacentNodes() []string {
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
	nodes map[string]Node
}

func NewGraph(oriented bool, weighed bool, hasNegativeWeights bool) Graph {
}

// Tries to add a node to the graph and returns true
// if successful, otherwise returns false if the
// node already exists.
func (g *Graph) AddNode(node string) bool {
}

// Tries to add a link between two nodes and returns true
// if successful, otherwise returns false if such a link
// already exists.
// Note: There can be many links between two nodes,
// but all of them have to have unique weights.
// If the graph isn't oriented adding a ling from A to B
// effectively adds a link from B to A.
func (g *Graph) AddLink(startNode string, endNode string, weight int) bool {
}

// Tries to remove the node from the graph and if
// successful removes all links between it and
// other nodes and returns true, otherwise if the
// node doesn't exist, returns false.
func (g *Graph) RemoveNode(node string) bool {
}

// Tries to remove the link from the graph and if
// successful returns true, otherwise if the link
// doesn't exist, returns false.
func (g *Graph) RemoveLink(startNode string, endNode string, weight int) bool {
}

// Tries to remove all links between node1 and node2
// and returns true if atleast one link was removed,
// otherwise returns false.
func (g *Graph) RemoveLinks(node1 string, node2 string) bool {
}

// Returns the count of the links which have as a starting
// node the one specified as a parameter.
// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g Graph) OutgoingLinksCount(node string) int {
}

// Returns the count of the links which have as an ending
// node the one specified as a parameter.
// Note: If the graph isn't oriented the outgoing links
// will always match the incoming links.
func (g Graph) IncomingLinksCount(node string) int {
}

// Returns a slice with the names of all the nodes
// in the graph.
func (g Graph) Nodes() []string {
}