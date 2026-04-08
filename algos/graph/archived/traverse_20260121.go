// Generated on 2026-01-21 09:00:20
// Daily practice file: traverse_20260121.go (traverse template)

package graph

type Graph struct {
	Nodes map[int][]int
}

func (g *Graph) AddEdge(from, to int) {
	if _, ok := g.Nodes[from]; !ok {
		g.Nodes[from] = make([]int, 0)
	}
	if _, ok := g.Nodes[to]; !ok {
		g.Nodes[to] = make([]int, 0)
	}
	g.Nodes[from] = append(g.Nodes[from], to)
	g.Nodes[to] = append(g.Nodes[to], from)
}

func (g *Graph) HasPath(from, to int) bool {
	return len(g.ShortestPath(from, to)) > 0
}

func (g *Graph) DFS(start int) []int {
	if g.Nodes[start] == nil {
		return []int{}
	}
	visited := make(map[int]bool)
	res := make([]int, 0, len(g.Nodes))
	g.dfs(start, visited, &res)
	return res
}

func (g *Graph) dfs(start int, visited map[int]bool, res *[]int) {
	if visited[start] {
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
	res := make([]int, 0)
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
	q := make([]int, 0)
	q = append(q, from)
	parents[from] = from
	_, ok := parents[to]
	for len(q) > 0 && !ok {
		n := q[0]
		q = q[1:]
		for _, j := range g.Nodes[n] {
			_, ok := parents[j]
			if ok {
				continue
			}
			parents[j] = n
			if j == to {
				return g.buildPath(parents, to)
			}
			q = append(q, j)
		}
	}
	return []int{}
}

func (g *Graph) buildPath(parents map[int]int, to int) []int {
	res := make([]int, 0)
	res = append(res, to)
	p, c := parents[to], to
	for p != c {
		res = append(res, p)
		p, c = parents[p], p
	}
	reverse(res)
	return res
}

func reverse(res []int) {
	for i := 0; i < len(res)/2; i++ {
		n := len(res) - i - 1
		res[i], res[n] = res[n], res[i]
	}
}

func NewGraph() *Graph {
	return &Graph{make(map[int][]int)}
}
