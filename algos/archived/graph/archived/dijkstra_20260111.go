// Generated on 2026-01-11 09:15:39
// Daily practice file: dijkstra_20260111.go (dijkstra template)

package graph

import (
	"container/heap"
	"math"
)

type DirectedGraph struct {
	Nodes map[int]map[int]int
}

func (d *DirectedGraph) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *DirectedGraph) AddWeightedEdge(from, to, weight int) {
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

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	if _, ok := d.Nodes[start]; !ok {
		return make(map[int]int)
	}
	costs := d.initialCosts(start)
	elmMap := make(map[int]*Elem)
	q := make(PQ, 0)
	for i, c := range costs {
		e := &Elem{i, c, -1}
		elmMap[i] = e
		heap.Push(&q, e)
	}
	for q.Len() > 0 {
		e := heap.Pop(&q).(*Elem)
		i := e.Val
		for j, w := range d.Nodes[i] {
			d.relax(i, j, w, costs)
			elmMap[j].Cost = costs[j]
			heap.Fix(&q, elmMap[j].Idx)
		}
	}

	return costs
}

func (d *DirectedGraph) relax(i, j, w int, costs map[int]int) {
	ci := costs[i]
	if ci != math.MaxInt {
		costs[j] = min(costs[j], ci+w)
	}
}

func (d *DirectedGraph) initialCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
	}
	res[start] = 0
	return res
}

type PQ []*Elem

type Elem struct {
	Val  int
	Cost int
	Idx  int
}

func (q PQ) Len() int { return len(q) }
func (q PQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].Idx, q[j].Idx = q[j].Idx, q[i].Idx
}
func (q PQ) Less(i, j int) bool {
	return q[i].Cost < q[j].Cost
}

func (q *PQ) Push(i interface{}) {
	e := i.(*Elem)
	n := q.Len()
	e.Idx = n
	*q = append(*q, e)
}

func (q *PQ) Pop() interface{} {
	old := *q
	n := q.Len()
	e := old[n-1]
	e.Idx = -1
	*q = old[:n-1]
	return e
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	ct, ok := costs[to] // panic record here, remember to judge the ok
	if ok && ct != math.MaxInt {
		return ct, true
	}
	return 0, false
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]map[int]int)}
}
