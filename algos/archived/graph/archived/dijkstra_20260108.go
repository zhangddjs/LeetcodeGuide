// Generated on 2026-01-08 09:14:38
// Daily practice file: dijkstra_20260108.go (dijkstra template)

package graph

import "math"

type DirectedGraph struct {
	Nodes map[int]*DNode
}

type DNode struct {
	Adjs    map[int]*DNode
	Weights map[int]int
}

func (d *DirectedGraph) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *DirectedGraph) AddWeightedEdge(from, to, weight int) {
	nf, ok := d.Nodes[from]
	if !ok {
		nf = &DNode{make(map[int]*DNode), make(map[int]int)}
		d.Nodes[from] = nf
	}
	nt, ok := d.Nodes[to]
	if !ok {
		nt = &DNode{make(map[int]*DNode), make(map[int]int)}
		d.Nodes[to] = nt
	}
	nf.Adjs[to] = nt
	nf.Weights[to] = weight
}

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	costs := d.initialCosts(start)
	q := NewPQ()
	for k, v := range costs {
		q.Push(k, v)
	}
	for !q.IsEmpty() {
		i := q.Pop()
		for j := range d.Nodes[i].Adjs {
			w := d.Nodes[i].Weights[j]
			d.relax(i, j, w, costs)
			q.Decrease(j, costs[j])
		}
	}
	return costs
}

func (d *DirectedGraph) relax(i, j, w int, costs map[int]int) {
	ci, ok := costs[i]
	if !ok || ci == math.MaxInt {
		return
	}
	if costs[j] > ci+w {
		costs[j] = ci + w
	}
}

func (d *DirectedGraph) initialCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
		if i == start {
			res[i] = 0
		}
	}
	return res
}

type PQ struct {
	Elems []int
	Costs map[int]int
	Idxs  map[int]int
}

func (q *PQ) Push(e, c int) {
	q.Elems = append(q.Elems, e)
	q.Costs[e] = c
	q.swiftup(len(q.Elems) - 1)
}

func (q *PQ) Pop() int {
	x := q.Elems[0]
	q.Elems = q.Elems[1:]
	delete(q.Costs, x)
	delete(q.Idxs, x)
	for i, e := range q.Elems {
		q.Idxs[e] = i
	}
	return x
}

func (q *PQ) IsEmpty() bool {
	return len(q.Elems) == 0
}

func (q *PQ) swiftup(n int) {
	for i := parent(n); n > 0 && q.Costs[q.Elems[i]] > q.Costs[q.Elems[n]]; {
		q.Elems[i], q.Elems[n] = q.Elems[n], q.Elems[i]
		q.Idxs[q.Elems[i]] = n
		q.Idxs[q.Elems[n]] = i
		n = i
		i = parent(i)
	}
}

func (q *PQ) Decrease(e, c int) {
	co := q.Costs[e]
	if co >= c {
		return
	}
	q.Costs[e] = c
	q.swiftup(q.Idxs[e])
}

func parent(i int) int {
	return (i - 1) / 2
}

func NewPQ() *PQ {
	return &PQ{make([]int, 0), make(map[int]int), make(map[int]int)}
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	ct, ok := costs[to]
	if !ok || ct == math.MaxInt {
		return 0, false
	}
	return costs[to], true
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]*DNode)}
}
