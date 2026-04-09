// Generated on 2026-01-07 08:33:23
// Daily practice file: dag_20260107.go (dag template)

package graph

import "math"

type Dag struct {
	Nodes map[int]*DagNode
}

type DagNode struct {
	Adjs    map[int]*DagNode
	Weights map[int]int
}

func (d *Dag) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *Dag) AddWeightedEdge(from, to, weight int) {
	nf, ok := d.Nodes[from]
	if !ok {
		nf = &DagNode{make(map[int]*DagNode), make(map[int]int)}
		d.Nodes[from] = nf
	}
	nt, ok := d.Nodes[to]
	if !ok {
		nt = &DagNode{make(map[int]*DagNode), make(map[int]int)}
		d.Nodes[to] = nt
	}
	nf.Adjs[to] = nt
	nf.Weights[to] = weight
}

func (d *Dag) TopologicalSort() []int {
	res := make([]int, 0, len(d.Nodes))
	visited := make(map[int]bool, len(d.Nodes))
	for i := range d.Nodes {
		d.toposort(i, visited, &res)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func (d *Dag) toposort(start int, visited map[int]bool, res *[]int) {
	if visited[start] {
		return
	}
	visited[start] = true
	for i := range d.Nodes[start].Adjs {
		d.toposort(i, visited, res)
	}
	*res = append(*res, start)
}

func (d *Dag) TopologicalSortBFS() []int {
	indegree := d.initialIndegree()
	q := make([]int, 0, len(d.Nodes))
	for k, v := range indegree {
		if v == 0 {
			q = append(q, k)
		}
	}
	res := make([]int, 0, len(d.Nodes))
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		res = append(res, n)
		for i := range d.Nodes[n].Adjs {
			indegree[i]--
			if indegree[i] == 0 {
				q = append(q, i)
			}
		}
	}
	return res
}

func (d *Dag) initialIndegree() map[int]int {
	res := make(map[int]int, len(d.Nodes))
	for i, n := range d.Nodes {
		res[i] = res[i]
		for j := range n.Adjs {
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

func (d *Dag) hasCycle(n int, colors map[int]int) bool {
	if colors[n] == 1 {
		return true
	} else if colors[n] == 2 {
		return false
	}
	colors[n] = 1
	for i := range d.Nodes[n].Adjs {
		if d.hasCycle(i, colors) {
			return true
		}
	}
	colors[n] = 2
	return false
}

func (d *Dag) ShortestPath(start int) map[int]int {
	nodes := d.TopologicalSort()
	costs := d.initialCosts(start)
	for i := range nodes {
		for j := range d.Nodes[i].Adjs {
			cost := d.Nodes[i].Weights[j]
			d.relax(i, j, cost, costs)
		}
	}
	return costs
}

func (d *Dag) relax(i, j, cost int, costs map[int]int) {
	ci, ok := costs[i]
	if !ok || ci == math.MaxInt {
		return
	}
	costs[j] = min(ci+cost, costs[j])
}

func (d *Dag) initialCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
		if i == start {
			res[i] = 0
		}
	}
	return res
}

func NewDAG() *Dag {
	return &Dag{make(map[int]*DagNode)}
}
