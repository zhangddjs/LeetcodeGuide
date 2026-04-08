// Generated on 2026-01-26 10:21:29
// Daily practice file: dijkstra_20260126.go (dijkstra template)

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
	if _, ok := d.Nodes[from]; !ok {
		d.Nodes[from] = make(map[int]int)
	}
	if _, ok := d.Nodes[to]; !ok {
		d.Nodes[to] = make(map[int]int)
	}
	d.Nodes[from][to] = weight
}

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	if d.Nodes[start] == nil {
		return make(map[int]int)
	}
	costs := d.initialCosts(start)
	q := make(Q, 0)
	heap.Push(&q, &Elem{start, 0})
	for q.Len() > 0 {
		i := heap.Pop(&q).(*Elem).Val
		for j, w := range d.Nodes[i] {
			if costs[i] != math.MaxInt && costs[i]+w < costs[j] {
				costs[j] = costs[i] + w
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

type Q []*Elem
type Elem struct {
	Val  int
	Cost int
}

func (q Q) Len() int            { return len(q) }
func (q Q) Less(i, j int) bool  { return q[i].Cost < q[j].Cost }
func (q Q) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *Q) Push(i interface{}) { *q = append(*q, i.(*Elem)) }
func (q *Q) Pop() interface{} {
	old := *q
	n := q.Len()
	e := old[n-1]
	*q = old[:n-1]
	return e
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	c, ok := costs[to]
	if !ok || c == math.MaxInt {
		return 0, false
	}
	return c, true
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]map[int]int)}
}
