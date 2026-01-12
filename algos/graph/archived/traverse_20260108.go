// Generated on 2026-01-08 09:14:38
// Daily practice file: traverse_20260108.go (traverse template)

package graph

type Graph struct {
	Nodes map[int]*GraphNode
}

type GraphNode struct {
	Adjs map[int]*GraphNode
}

func (g *Graph) AddEdge(from, to int) {
	nf, ok := g.Nodes[from]
	if !ok {
		g.Nodes[from] = &GraphNode{make(map[int]*GraphNode)}
		nf = g.Nodes[from]
	}
	nt, ok := g.Nodes[to]
	if !ok {
		g.Nodes[to] = &GraphNode{make(map[int]*GraphNode)}
		nt = g.Nodes[to]
	}
	nf.Adjs[to] = nt
	nt.Adjs[from] = nf
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) > 0
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	res := make([]int, 0, len(g.Nodes))
	g.dfs(start, visited, &res)
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool, res *[]int) {
	if visited[start] || g.Nodes[start] == nil {
		return
	}
	visited[start] = true
	*res = append(*res, start)
	for i := range g.Nodes[start].Adjs {
		g.dfs(i, visited, res)
	}
}

func (g *Graph) DFSWhole() []int {
	visited := make(map[int]bool)
	res := make([]int, 0, len(g.Nodes))
	for i := range g.Nodes {
		g.dfs(i, visited, &res)
	}
	return res
}

func (g *Graph) BFSWhole() []int { return g.DFSWhole() }

func (g *Graph) ShortestPath(from, to int) []int {
	if g.Nodes[from] == nil || g.Nodes[to] == nil {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	parents := make(map[int]int, len(g.Nodes))
	visited := make(map[int]bool, len(g.Nodes))
	q := make([]int, 0, len(g.Nodes))
	q = append(q, from)
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for i := range g.Nodes[n].Adjs {
			if visited[i] {
				continue
			}
			visited[i] = true
			q = append(q, i)
			parents[i] = n
			if i == to {
				break
			}
		}
	}
	res := make([]int, 0, len(g.Nodes))
	p, ok := parents[to]
	if !ok {
		return []int{}
	}
	res = append(res, to)
	for p != from {
		res = append(res, p)
		p = parents[p]
	}
	res = append(res, from)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int]*GraphNode)}
}
