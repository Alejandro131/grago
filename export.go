package grago

// Export the graph structure in a format specified
// by the dot language used in graphviz, which can
// be loaded by the software and visualized.
// 'highlights' is a slice of links representing edges
// of the graph you would like to outline, for example
// the results of a DFS search and they will be colored red.
// 'ordered' is a boolean indicating whether or not the index
// of the links would be shown next to their name which is
// useful for distinguishing DFS from MST for example.
// 'clusters' is a 2d array of strings representing node names
// that should be ordered in separate groups visually, if it's
// empty, the graph will be displayed as is.
func (g Graph) Export(highlights []Link, ordered bool, clusters [][]string) string {
}
