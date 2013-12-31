# About grago

Implementing various graph algorithms, `grago` is a library writen in Go.

### Setup

##### Prerequisites

Make sure you have a working Go installation with version `Go 1.1.x` or later.
See [Getting Started](http://golang.org/doc/install.html)

##### To Install
Run `go get github.com/Alejandro131/grago`

### Examples

Before starting with the code, we have to import the `grago` package:
`import "github.com/Alejandro131/grago"`

##### Creating a graph

To create a graph, simply use the `NewGraph` function which accepts as arguments
3 boolean variables indicating whether the graph will be oriented, weighed and
if it is weighed, whether or not there can be negative edges. The following code
will create a graph that isn't oriented and is weighed with positive values only:
`graph := NewGraph(false, true, false)`

##### Adding nodes and edges to it

To add a node or edges, use the functions `AddNode` and `AddLink` accordingly:
`graph.AddNode("alpha")
graph.AddLink("2", "alpha", 2)`
Adding a link which has a node (or 2 nodes) not previously added to the graph,
effectively adds the node to the graph, so you don't have to previously initialize
all the nodes and add the links later, but rather construct the graph from links.

##### Using algorithms on the structure

There are popular algorithms in graph theory problems which you can use on the graph
structure. Depending on the algorithm, you can expect differently formated output from
each of the functions. Below are several examples:

* BFS

??

* DFS

??

* MinPath

??

* Graph Properties

??

##### Exporting the graph to view in [Graphviz](http://www.graphviz.org/)

??

![](graph.png)
![](graphHighlights.png)

### License

The library is licensed under the [MIT License](LICENSE).