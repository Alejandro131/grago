package grago

type Link struct {
	Start string
	End string
	Weight int
}

func NewLink(start string, end string, weight int) {
}

type Node struct {
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

func (g Graph) Nodes() []string {
}