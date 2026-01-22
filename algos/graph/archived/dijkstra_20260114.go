// Generated on 2026-01-14 08:18:27
// Daily practice file: dijkstra_20260114.go (dijkstra template)

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
}

func (q PQ) Len() int            { return len(q) }
func (q PQ) Less(i, j int) bool  { return q[i].Cost < q[j].Cost }
func (q PQ) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *PQ) Push(i interface{}) { *q = append(*q, i.(*Elem)) }
func (q *PQ) Pop() interface{} {
	old := *q
	n := q.Len()
	e := old[n-1]
	*q = old[:n-1]
	return e
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	if ct, ok := costs[to]; !ok || ct == math.MaxInt {
		return 0, false
	}
	return costs[to], true
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]map[int]int)}
}
