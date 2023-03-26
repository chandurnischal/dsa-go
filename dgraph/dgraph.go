package dgraph

type Graph struct {
	N             int
	AdjacencyList map[int][]int
}

func (g *Graph) AddNode(node int) {
	if _, ok := g.AdjacencyList[node]; ok {
		return
	}
	if g.AdjacencyList == nil {
		g.AdjacencyList = make(map[int][]int)
	}
	g.AdjacencyList[node] = []int{}
	g.N++
}

func (g *Graph) AddEdge(node1, node2 int) {
	if _, ok := g.AdjacencyList[node1]; !ok {
		g.AddNode(node1)
	}
	if _, ok := g.AdjacencyList[node2]; !ok {
		g.AddNode(node2)
	}
	g.AdjacencyList[node1] = append(g.AdjacencyList[node1], node2)
}
