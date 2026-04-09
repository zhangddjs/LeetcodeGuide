// Generated on 2026-01-12 11:12:21
// Daily practice file: dijkstra_20260112.go (dijkstra template)

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
	_, ok := d.Nodes[start]
	if !ok {
		return make(map[int]int)
	}
	costs := d.initalCosts(start)
	q := make(PQ, 0)
	heap.Init(&q)
	heap.Push(&q, &Elem{start, 0})
	for q.Len() > 0 {
		n := heap.Pop(&q).(*Elem)
		for j, w := range d.Nodes[n.Val] {
			if n.Cost+w < costs[j] {
				costs[j] = n.Cost + w
				heap.Push(&q, &Elem{j, costs[j]})
			}
		}
	}
	return costs
}

func (d *DirectedGraph) initalCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
	}
	res[start] = 0
	return res
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	ct, ok := costs[to]
	if ok && ct != math.MaxInt {
		return ct, true
	}
	return 0, false
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]map[int]int)}
}

type Elem struct {
	Val  int
	Cost int
}
type PQ []*Elem

func (q PQ) Len() int           { return len(q) }
func (q PQ) Less(i, j int) bool { return q[i].Cost < q[j].Cost }
func (q PQ) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *PQ) Push(i interface{}) {
	e := i.(*Elem)
	*q = append(*q, e)
}

func (q *PQ) Pop() interface{} {
	old := *q
	n := q.Len()
	e := old[n-1]
	*q = old[:n-1]
	return e
}
