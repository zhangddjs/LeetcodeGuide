// Generated on 2026-01-09 08:23:06
// Daily practice file: dijkstra_20260109.go (dijkstra template)

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
	for e, c := range costs {
		q.Push(e, c)
	}
	for !q.IsEmpty() {
		e := q.Pop()
		for j := range d.Nodes[e].Adjs {
			w := d.Nodes[e].Weights[j]
			d.relax(e, j, w, costs)
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
	costs[j] = min(costs[j], ci+w)
}

func (d *DirectedGraph) initialCosts(start int) map[int]int {
	res := make(map[int]int, len(d.Nodes))
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

func NewPQ() *PQ {
	return &PQ{make([]int, 0), make(map[int]int), make(map[int]int)}
}

func (q *PQ) Push(e, c int) {
	q.Elems = append(q.Elems, e)
	q.Costs[e] = c
	q.Idxs[e] = len(q.Elems) - 1
	q.swiftup(len(q.Elems) - 1)
}

func (q *PQ) Pop() int {
	n := len(q.Elems)
	e, x := q.Elems[0], q.Elems[n-1]
	q.Elems[0], q.Elems[n-1] = q.Elems[n-1], q.Elems[0]
	q.Elems = q.Elems[:n-1]
	q.Idxs[x] = 0
	delete(q.Costs, e)
	delete(q.Idxs, e)
	q.swiftdown(0, n-1)
	return e
}

func (q *PQ) IsEmpty() bool {
	return len(q.Elems) == 0
}

func (q *PQ) swiftup(i int) {
	for p := parent(i); i > 0 && q.Costs[q.Elems[i]] < q.Costs[q.Elems[p]]; i, p = p, parent(p) {
		q.Elems[i], q.Elems[p] = q.Elems[p], q.Elems[i]
		q.Idxs[q.Elems[i]], q.Idxs[q.Elems[p]] = p, i
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func (q *PQ) swiftdown(i, n int) {
	for true {
		l, r := children(i)
		minIdx := i
		if l < n && q.Costs[q.Elems[l]] < q.Costs[q.Elems[minIdx]] {
			minIdx = l
		}
		if r < n && q.Costs[q.Elems[l]] < q.Costs[q.Elems[minIdx]] {
			minIdx = r
		}
		if i == minIdx {
			break
		}
		q.Elems[minIdx], q.Elems[i] = q.Elems[i], q.Elems[minIdx]
		q.Idxs[q.Elems[minIdx]], q.Idxs[q.Elems[i]] = i, minIdx
	}
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

func (q *PQ) Decrease(e, c int) {
	if q.Costs[e] <= c {
		return
	}
	q.Costs[e] = c
	q.swiftup(q.Idxs[e])
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	ct, ok := costs[to]
	if !ok || ct == math.MaxInt {
		return 0, false
	}
	return ct, true
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]*DNode)}
}
