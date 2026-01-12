package graph

import "math"

type Dag struct {
	Nodes map[int]*DagNode
}

type DagNode struct {
	Val     int
	Adjs    map[int]*DagNode
	Weights map[int]int
}

func (d *Dag) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 0)
}

func (d *Dag) AddWeightedEdge(from, to, weight int) {
	df, ok := d.Nodes[from]
	if !ok {
		df = &DagNode{from, make(map[int]*DagNode), make(map[int]int)}
		d.Nodes[from] = df
	}
	dt, ok := d.Nodes[to]
	if !ok {
		dt = &DagNode{to, make(map[int]*DagNode), make(map[int]int)}
		d.Nodes[to] = dt
	}
	df.Adjs[to] = dt
	df.Weights[to] = weight
}

func (d *Dag) TopologicalSort() []int {
	visited := make(map[int]bool, len(d.Nodes))
	res := make([]int, 0, len(d.Nodes))
	for _, n := range d.Nodes {
		if !visited[n.Val] {
			d.toposort(n.Val, visited, &res)
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func (d *Dag) toposort(v int, visited map[int]bool, res *[]int) {
	visited[v] = true
	for k := range d.Nodes[v].Adjs {
		if !visited[k] {
			d.toposort(k, visited, res)
		}
	}
	*res = append(*res, v)
}

func (d *Dag) TopologicalSortBFS() []int {
	indegree := make(map[int]int, len(d.Nodes))
	for i, n := range d.Nodes {
		indegree[i] = indegree[i]
		for j := range n.Adjs {
			indegree[j]++
		}
	}

	res := make([]int, 0, len(d.Nodes))
	Q := make([]int, 0, len(d.Nodes))
	for i, d := range indegree {
		if d == 0 {
			Q = append(Q, i)
		}
	}

	for len(Q) > 0 {
		n := Q[0]
		Q = Q[1:]
		res = append(res, n)
		for i := range d.Nodes[n].Adjs {
			indegree[i]--
			if indegree[i] == 0 {
				Q = append(Q, i)
			}
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
	dist := d.initialDis(start)

	nodes := d.TopologicalSort()

	for i := range nodes {
		for j := range d.Nodes[i].Adjs {
			w := d.Nodes[i].Weights[j]
			d.relax(i, j, w, dist)
		}
	}

	return dist
}

func (d *Dag) relax(i, j, w int, dist map[int]int) {
	costi, costj := dist[i], dist[j]
	if costi == math.MaxInt {
		return
	}
	if costi+w < costj {
		dist[j] = costi + w
	}
}

func (d *Dag) initialDis(start int) map[int]int {
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
