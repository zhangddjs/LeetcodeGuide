// Generated on 2026-01-07 08:33:23
// Daily practice file: traverse_20260107.go (traverse template)

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

func (g *Graph) ShortestPath(from, to int) []int {
	if g.Nodes[from] == nil || g.Nodes[to] == nil {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	parents := make(map[int]int)
	visited := make(map[int]bool)
	q := make([]int, 0, len(g.Nodes))
	q = append(q, from)
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		for i := range g.Nodes[e].Adjs {
			if visited[i] {
				continue
			}
			parents[i] = e
			visited[i] = true
			if i == to {
				break
			}
			q = append(q, i)
		}
	}
	p, ok := parents[to]
	if !ok {
		return []int{}
	}
	res := make([]int, 0, len(g.Nodes))
	res = append(res, to)
	for ok && p != from {
		res = append(res, p)
		p, ok = parents[p]
	}
	res = append(res, from)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int]*GraphNode)}
}
