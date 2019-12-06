package day06

type Edge struct {
	Source      string
	Destination string
}

type Graph struct {
	graph map[string][]string
}

func NewGraph() Graph {
	var g Graph
	g.graph = make(map[string][]string)
	return g
}

func (g *Graph) AddEdge(e Edge) {
	g.ensureNode(e.Source)
	g.ensureNode(e.Destination)
	g.graph[e.Source] = append(g.graph[e.Source], e.Destination)
}

func (g *Graph) NodeExists(node string) bool {
	_, ok := g.graph[node]
	return ok
}

func (g *Graph) Neighbors(node string) []string {
	var res []string
	if g.NodeExists(node) {
		for _, neighbor := range g.graph[node] {
			res = append(res, neighbor)
		}
	}
	return res
}

func (g *Graph) Nodes() []string {
	var res []string
	for k := range g.graph {
		res = append(res, k)
	}
	return res
}

func (g *Graph) Reachable(node string) int {
	total := 0

	neighbors := g.Neighbors(node)
	total += len(neighbors)
	for _, neighbor := range neighbors {
		total += g.Reachable(neighbor)
	}

	return total
}

func (g *Graph) ensureNode(node string) {
	if !g.NodeExists(node) {
		g.graph[node] = make([]string, 0)
	}
}
