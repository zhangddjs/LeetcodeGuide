// Generated on 2026-01-15 07:49:40
// Daily practice file: traverse_20260115.go (traverse template)

package graph

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) AddEdge(from, to int) {
	g.Nodes[from] = append(g.Nodes[from], to)
	g.Nodes[to] = append(g.Nodes[to], from)
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) > 0
}

func (g *Graph) DFS(start int) []int {
	if len(g.Nodes[start]) == 0 {
		return []int{}
	}
	visited := make(map[int]bool)
	res := make([]int, 0)
	g.dfs(start, visited, &res)
	return res
}

func (g *Graph) dfs(i int, visited map[int]bool, res *[]int) {
	if visited[i] {
		return
	}
	visited[i] = true
	*res = append(*res, i)
	for _, j := range g.Nodes[i] {
		g.dfs(j, visited, res)
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
	if len(g.Nodes[from]) == 0 || len(g.Nodes[to]) == 0 {
		return []int{}
	}
	if from == to {
		return []int{from}
	}
	parents := make(map[int]int)
	q := make([]int, 0)
	q = append(q, from)
	parents[from] = from
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, j := range g.Nodes[n] {
			if _, ok := parents[j]; ok {
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
	res := make([]int, 0)
	p, c := parents[to], to
	res = append(res, to)
	for p != c {
		res = append(res, p)
		p, c = parents[p], p
	}
	reverse(res)
	return res
}

func reverse(res []int) {
	for i := 0; i < len(res)/2; i++ {
		n := len(res) - 1 - i
		res[i], res[n] = res[n], res[i]
	}
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]int)}
}
