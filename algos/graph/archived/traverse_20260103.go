// Generated on 2026-01-03 14:34:22
// Daily practice file: traverse_20260103.go

package graph

type Graph struct {
	Vertices map[int]*GraphNode
}

type GraphNode struct {
	Adjs map[int]*GraphNode
}

func (g *Graph) AddEdge(from, to int) {
	nf, ok := g.Vertices[from]
	if !ok {
		g.Vertices[from] = &GraphNode{make(map[int]*GraphNode)}
		nf = g.Vertices[from]
	}
	nt, ok := g.Vertices[to]
	if !ok {
		g.Vertices[to] = &GraphNode{make(map[int]*GraphNode)}
		nt = g.Vertices[to]
	}
	nf.Adjs[to] = nt
	nt.Adjs[from] = nf
}

func (g *Graph) HasPath(from, to int) bool {
	shortestPath := g.ShortestPath(from, to)
	return len(shortestPath) > 0
}

func (g *Graph) DFS(start int) []int {
	_, ok := g.Vertices[start]
	if !ok {
		return []int{}
	}
	visited := make(map[int]bool, len(g.Vertices))
	g.dfs(start, visited)
	res := make([]int, 0, len(g.Vertices))
	for i := range visited {
		res = append(res, i)
	}
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool) {
	ns, ok := g.Vertices[start]
	if !ok || visited[start] {
		return
	}
	visited[start] = true
	for n := range ns.Adjs {
		g.dfs(n, visited)
	}
}

func (g *Graph) ShortestPath(from, to int) []int {
	if g.Vertices[from] == nil || g.Vertices[to] == nil {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	parents := make(map[int]int, len(g.Vertices))
	visited := make(map[int]bool, len(g.Vertices))
	Q := make([]int, 0, len(g.Vertices))
	Q = append(Q, from)
	for len(Q) > 0 {
		n := Q[0]
		Q = Q[1:]
		visited[n] = true
		if g.Vertices[n].Adjs[to] != nil {
			parents[to] = n
			break
		}
		for i := range g.Vertices[n].Adjs {
			if visited[i] {
				continue
			}
			Q = append(Q, i)
			parents[i] = n
		}
	}
	res := make([]int, 0, len(g.Vertices))
	if _, ok := parents[to]; !ok {
		return res
	}
	res = append(res, to)
	p, ok := parents[to]
	for ok {
		res = append(res, p)
		p, ok = parents[p]
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int]*GraphNode)}
}
