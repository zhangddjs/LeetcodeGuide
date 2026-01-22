// Generated on 2026-01-12 11:12:21
// Daily practice file: traverse_20260112.go (traverse template)

package graph

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) AddEdge(from, to int) {
	nf, ok := g.Nodes[from]
	if !ok {
		nf = make([]int, 0)
		g.Nodes[from] = nf
	}
	nt, ok := g.Nodes[to]
	if !ok {
		nt = make([]int, 0)
		g.Nodes[to] = nt
	}
	g.Nodes[from] = append(g.Nodes[from], to)
	g.Nodes[to] = append(g.Nodes[to], from)
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
	if visited[start] || len(g.Nodes[start]) == 0 {
		return
	}
	visited[start] = true
	*res = append(*res, start)
	for _, j := range g.Nodes[start] {
		g.dfs(j, visited, res)
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
	if from == to {
		return []int{from}
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
			q = append(q, j)
			if j == to {
				return buildPath(parents, j)
			}
		}
	}
	return []int{}
}

func buildPath(parents map[int]int, to int) []int {
	res := make([]int, 0)
	res = append(res, to)
	for p, c := parents[to], to; p != c; p, c = parents[p], p {
		res = append(res, p)
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
