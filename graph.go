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
	// A container for all the nodes, that this node has an outgoing connection to.
	Adjacent map[string]int
}

func NewNode() Node {
}

func (n Node) AdjacentNodes() []string {
}

type Graph struct {
	Oriented bool
	Weighed bool
	HasNegativeWeights bool
	nodes map[string]Node
}

func NewGraph(oriented bool, weighed bool, hasNegativeWeights bool) Graph {
}

func (g *Graph) AddNode(node string) bool {
}

func (g *Graph) AddLink(startNode string, endNode string, weight int) bool {
}

func (g *Graph) RemoveNode(node string) bool {
}

func (g *Graph) RemoveLink(startNode string, endNode string, weight int) bool {
}

func (g Graph) OutgoingLinksCount(node string) int {
}

func (g Graph) IncomingLinksCount(node string) int {
}

func (g Graph) Nodes() []string {
}