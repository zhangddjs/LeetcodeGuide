// Generated on 2026-01-13 08:29:45
// Daily practice file: dag_20260113.go (dag template)

package graph

import "math"

type Dag struct {
	Nodes map[int]map[int]int
}

func (d *Dag) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *Dag) AddWeightedEdge(from, to, weight int) {
	nf, ok := d.Nodes[from]
	if !ok {
		nf = make(map[int]int)
		d.Nodes[from] = nf
	}
	nt, ok := d.Nodes[to]
	if !ok {
		nt = make(map[int]int)
		d.Nodes[to] = nt
	}
	nf[to] = weight
}

func (d *Dag) TopologicalSort() []int {
	visited := make(map[int]bool)
	res := make([]int, 0, len(d.Nodes))
	for i := range d.Nodes {
		d.topo(i, visited, &res)
	}
	return reverse(res)
}

func (d *Dag) topo(i int, visited map[int]bool, res *[]int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for j := range d.Nodes[i] {
		d.topo(j, visited, res)
	}
	*res = append(*res, i)
}

func reverse(res []int) []int {
	for i := 0; i < len(res)/2; i++ {
		n := len(res) - 1 - i
		res[i], res[n] = res[n], res[i]
	}
	return res
}

func (d *Dag) TopologicalSortBFS() []int {
	indegrees := d.initialIndegrees()
	res := make([]int, 0, len(d.Nodes))
	q := make([]int, 0, len(d.Nodes))
	for i, d := range indegrees {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		res = append(res, n)
		for j := range d.Nodes[n] {
			indegrees[j]--
			if indegrees[j] == 0 {
				q = append(q, j)
			}
		}
	}
	return res
}

func (d *Dag) initialIndegrees() map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = res[i]
		for j := range d.Nodes[i] {
			res[j]++
		}
	}
	return res
}

func (d *Dag) IsDAG() bool {
	colors := make(map[int]int)
	for i := range d.Nodes {
		if d.hasCycle(i, colors) {
			return false
		}
	}
	return true
}

func (d *Dag) hasCycle(i int, colors map[int]int) bool {
	if colors[i] == 1 {
		return true
	} else if colors[i] == 2 {
		return false
	}
	colors[i] = 1
	for j := range d.Nodes[i] {
		if d.hasCycle(j, colors) {
			return true
		}
	}
	colors[i] = 2
	return false
}

func (d *Dag) ShortestPath(start int) map[int]int {
	_, ok := d.Nodes[start]
	if !ok {
		return make(map[int]int)
	}
	nodes := d.TopologicalSort()
	costs := d.initialCosts(start)
	for i := range nodes {
		for j, w := range d.Nodes[i] {
			d.relax(i, j, w, costs)
		}
	}
	return costs
}

func (d *Dag) relax(i, j, w int, costs map[int]int) {
	ci := costs[i]
	if ci != math.MaxInt {
		costs[j] = min(costs[j], ci+w)
	}
}

func (d *Dag) initialCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
	}
	res[start] = 0
	return res
}

func NewDAG() *Dag {
	return &Dag{make(map[int]map[int]int)}
}
