// Generated on 2026-01-09 08:23:06
// Daily practice file: traverse_20260109.go (traverse template)

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
		nf = &GraphNode{make(map[int]*GraphNode)}
		g.Nodes[from] = nf
	}
	nt, ok := g.Nodes[to]
	if !ok {
		nt = &GraphNode{make(map[int]*GraphNode)}
		g.Nodes[to] = nt
	}
	nf.Adjs[to] = nt
	nt.Adjs[from] = nf
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) != 0
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

func (g *Graph) ShortestPath(from, to int) []int {
	if g.Nodes[from] == nil || g.Nodes[to] == nil {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	parents := make(map[int]int, len(g.Nodes))
	visited := make(map[int]bool, len(g.Nodes))
	res := make([]int, 0, len(g.Nodes))
	q := make([]int, 0, len(g.Nodes))
	q = append(q, from)
	visited[from] = true
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for i := range g.Nodes[n].Adjs {
			if visited[i] {
				continue
			}
			visited[i] = true
			parents[i] = n
			if i == to {
				q = make([]int, 0)
				break
			}
			q = append(q, i)
		}
	}
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
		res[i], res[len(res)-i-1] = res[len(res)-1-i], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int]*GraphNode)}
}
