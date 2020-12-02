package day06_2019

// Edge represents an edge in a adj
type Edge struct {
	Source      string
	Destination string
}

// Graph represents a directed adj where the nodes are strings
type Graph struct {
	// Map where the keys are the nodes and the values are the list of nodes
	// that are neighbors of the key node
	adj map[string][]string
}

// NewGraph creates and initializes a new adj
func NewGraph() Graph {
	var g Graph
	g.adj = make(map[string][]string)
	return g
}

// AddEdge adds an edge to the adj
func (g *Graph) AddEdge(e Edge) {
	g.ensureNode(e.Source)
	g.ensureNode(e.Destination)
	g.adj[e.Source] = append(g.adj[e.Source], e.Destination)
}

// NodeExists tells if a node exists in the adj
func (g *Graph) NodeExists(node string) bool {
	_, ok := g.adj[node]
	return ok
}

// Neighbors returns a list of neighbors for a given node
func (g *Graph) Neighbors(node string) []string {
	var res []string
	if g.NodeExists(node) {
		for _, neighbor := range g.adj[node] {
			res = append(res, neighbor)
		}
	}
	return res
}

// Nodes returns the list of nodes the the adj
func (g *Graph) Nodes() []string {
	var res []string
	for k := range g.adj {
		res = append(res, k)
	}
	return res
}

// Reachable returns how many nodes are reachable from a given one.
// The current implementation naively follows neighbors without checking for cycles.
// This is ok for the problem being solved, where we know we don't have cycles.
func (g *Graph) Reachable(node string) int {
	total := 0

	neighbors := g.Neighbors(node)
	total += len(neighbors)
	for _, neighbor := range neighbors {
		total += g.Reachable(neighbor)
	}

	return total
}

// Path returns list of nodes in the path from startNode to endNode.
// The current implementation naively traverses the adj by only checking
// the first neighbor of each node.
// This is ok for the problem being solved, where we know the adj is a tree
// and we can safely assume there's always a direct path between start and end
func (g *Graph) Path(startNode string, endNode string) []string {
	var res []string

	currentNode := startNode
	neighbors := g.Neighbors(currentNode)

	for len(neighbors) != 0 && currentNode != endNode {
		currentNode = neighbors[0]
		neighbors = g.Neighbors(currentNode)
		res = append(res, currentNode)
	}

	return res
}

// ensureNode ensures there's an entry for the node in the adjacency map
func (g *Graph) ensureNode(node string) {
	if !g.NodeExists(node) {
		g.adj[node] = make([]string, 0)
	}
}
