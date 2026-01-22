// Generated on 2026-01-13 08:29:45
// Daily practice file: traverse_20260113.go (traverse template)

package graph

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) AddEdge(from, to int) {
	_, ok := g.Nodes[from]
	if !ok {
		g.Nodes[from] = make([]int, 0)
	}
	_, ok = g.Nodes[to]
	if !ok {
		g.Nodes[to] = make([]int, 0)
	}
	g.Nodes[from] = append(g.Nodes[from], to)
	g.Nodes[to] = append(g.Nodes[to], from)
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) > 0
}

func (g *Graph) DFS(start int) []int {
	visited := make(map[int]bool)
	res := make([]int, 0)
	g.dfs(start, visited, &res)
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool, res *[]int) {
	n, ok := g.Nodes[start]
	if visited[start] || !ok {
		return
	}
	visited[start] = true
	*res = append(*res, start)
	for _, i := range n {
		g.dfs(i, visited, res)
	}
}

func (g *Graph) DFSWhole() []int {
	visited := make(map[int]bool)
	res := make([]int, 0)
	for i := range g.Nodes {
		g.dfs(i, visited, &res)
	}
	return res
}

func (g *Graph) ShortestPath(from, to int) []int {
	if from == to {
		return []int{from}
	}
	if len(g.Nodes[from]) == 0 || len(g.Nodes[to]) == 0 {
		return []int{}
	}
	parents := make(map[int]int)
	q := make([]int, 0, len(g.Nodes))
	q = append(q, from)
	parents[from] = from
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, j := range g.Nodes[n] {
			_, ok := parents[j]
			if ok {
				continue
			}
			parents[j] = n
			if j == to {
				return buildPath(parents, to)
			}
			q = append(q, j)
		}
	}
	return []int{}
}

func buildPath(parents map[int]int, to int) []int {
	res := make([]int, 0, len(parents))
	res = append(res, to)
	p := parents[to]
	n := to
	for p != n {
		res = append(res, p)
		n = p
		p = parents[p]
	}

	for i := 0; i < len(res)/2; i++ {
		n := len(res) - 1 - i
		res[i], res[n] = res[n], res[i]
	}
	return res
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]int)}
}
