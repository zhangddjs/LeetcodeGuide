// Generated on 2026-01-11 09:15:39
// Daily practice file: traverse_20260111.go (traverse template)

package graph

type Graph struct {
	Nodes map[int]map[int]bool
}

func (g *Graph) AddEdge(from, to int) {
	nf, ok := g.Nodes[from]
	if !ok {
		nf = make(map[int]bool)
		g.Nodes[from] = nf
	}
	nt, ok := g.Nodes[to]
	if !ok {
		nt = make(map[int]bool)
		g.Nodes[to] = nt
	}
	nf[to] = true
	nt[from] = true
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) > 0
}

func (g *Graph) DFS(start int) []int {
	res := make([]int, 0, len(g.Nodes))
	visited := make(map[int]bool)
	g.dfs(start, visited, &res)
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool, res *[]int) {
	if visited[start] || g.Nodes[start] == nil {
		return
	}
	visited[start] = true
	*res = append(*res, start)
	for i := range g.Nodes[start] {
		g.dfs(i, visited, res)
	}
}

func (g *Graph) DFSWhole() []int {
	res := make([]int, 0, len(g.Nodes))
	visited := make(map[int]bool)
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
	parents := make(map[int]int)
	visited := make(map[int]bool)
	res := make([]int, 0, len(g.Nodes))
	q := make([]int, 0, len(g.Nodes))
	q = append(q, from)
	ok := false
	visited[from] = true
	for len(q) > 0 && !ok {
		n := q[0]
		q = q[1:]
		for i := range g.Nodes[n] {
			if visited[i] {
				continue
			}
			visited[i] = true
			parents[i] = n
			q = append(q, i)
			if i == to {
				ok = true
				break
			}
		}
	}
	if !ok {
		return []int{}
	}
	res = append(res, to)
	p, ok := parents[to]
	for ok {
		res = append(res, p)
		p, ok = parents[p]
	}
	for i := 0; i < len(res)/2; i++ {
		n := len(res) - 1 - i
		res[i], res[n] = res[n], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int]map[int]bool)}
}
