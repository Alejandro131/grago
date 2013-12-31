package grago

// Export the graph structure in a format specified
// by the dot language used in graphviz, which can
// be loaded by the software and visualized.
// highlights is a slice of links representing edges
// of the graph you would like to outline, for example
// the results of a DFS search and they will be colored red
func (g Graph) Export(highlights []Link) string {
}