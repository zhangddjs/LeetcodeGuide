package graph

type Graph struct {
	Vertx map[int]*GraphNode
}

type GraphNode struct {
	Val    int
	AdjMap map[int]*GraphNode
}

func (g *Graph) AddEdge(from, to int) {
	nf, oknf := g.Vertx[from]
	if !oknf {
		nf = &GraphNode{from, make(map[int]*GraphNode)}
		g.Vertx[from] = nf
	}
	if from == to {
		return
	}
	nt, oknt := g.Vertx[to]
	if !oknt {
		nt = &GraphNode{to, make(map[int]*GraphNode)}
		g.Vertx[to] = nt
	}
	nf.AdjMap[to] = nt
	nt.AdjMap[from] = nf
}

func (g *Graph) HasPath(from, to int) bool {
	path := g.ShortestPath(from, to)
	return len(path) > 0
}

func (g *Graph) BFS(start int) []int {
	if g.Vertx[start] == nil {
		return []int{}
	}
	visited := make(map[int]bool)
	q := make([]int, 0)
	q = append(q, start)
	for len(q) > 0 {
		cur := q[0]
		vertx := g.Vertx[cur]
		q = q[1:]
		visited[cur] = true
		for k := range vertx.AdjMap {
			if visited[k] {
				continue
			}
			q = append(q, k)
		}
	}

	res := make([]int, 0, len(visited))
	for v := range visited {
		res = append(res, v)
	}
	return res
}

func (g *Graph) DFS(start int) []int {
	if g.Vertx[start] == nil {
		return []int{}
	}
	visited := make(map[int]bool)
	g.dfs(start, visited)
	res := make([]int, 0, len(visited))
	for k := range visited {
		res = append(res, k)
	}
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	for i := range g.Vertx[start].AdjMap {
		g.dfs(i, visited)
	}
}

// assume only one path
func (g *Graph) ShortestPath(from, to int) []int {
	if g.Vertx[from] == nil || g.Vertx[to] == nil || from == to {
		return []int{}
	}
	q := make([]int, 0, len(g.Vertx))
	visited := make(map[int]bool)
	res := make([]int, 0, len(g.Vertx))
	q = append(q, from)
	size, i := 1, 0
	res = append(res, from)
	for len(q) > 0 {
		cur := q[0]
		size--
		res[i] = cur
		visited[cur] = true
		q = q[1:]
		vert := g.Vertx[cur]
		if vert.AdjMap[to] != nil {
			return append(res, to)
		}
		for k := range vert.AdjMap {
			if visited[k] {
				continue
			}
			q = append(q, k)
		}
		if size == 0 {
			size = len(q)
			i++
			res = append(res, cur)
		}
	}
	return []int{}
}

func NewGraph() *Graph {
	return &Graph{make(map[int]*GraphNode)}
}
