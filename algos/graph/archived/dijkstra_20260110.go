// Generated on 2026-01-10 08:24:38
// Daily practice file: dijkstra_20260110.go (dijkstra template)

package graph

import "math"

type DirectedGraph struct {
	Nodes map[int]*DNode
}

type DNode struct {
	Adjs    map[int]*DNode
	Weights map[int]int
}

func NewDNode() *DNode {
	return &DNode{make(map[int]*DNode), make(map[int]int)}
}

func (d *DirectedGraph) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *DirectedGraph) AddWeightedEdge(from, to, weight int) {
	nf, ok := d.Nodes[from]
	if !ok {
		nf = NewDNode()
		d.Nodes[from] = nf
	}
	nt, ok := d.Nodes[to]
	if !ok {
		nt = NewDNode()
		d.Nodes[to] = nt
	}
	nf.Adjs[to] = nt
	nf.Weights[to] = weight
}

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	costs := d.initialCosts(start)
	q := NewPQ()
	for i, c := range costs {
		q.Push(i, c)
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
	costs[j] = min(costs[j], ci+w)
}

func (d *DirectedGraph) initialCosts(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
		if i == start {
			res[start] = 0
		}
	}
	return res
}

type PQ struct {
	Elems []*Elem
	Idxs  map[int]int
}

func NewPQ() *PQ {
	return &PQ{make([]*Elem, 0), make(map[int]int)}
}

type Elem struct {
	Val  int
	Cost int
}

func (p *PQ) Decrease(e, c int) {
	i := p.Idxs[e]
	if i >= len(p.Elems) { // panic record here, remember to handle back edge
		return
	}
	ele := p.Elems[i]
	if ele.Cost <= c {
		return
	}
	ele.Cost = c
	p.swiftup(i)
}

func (p *PQ) Push(e, c int) {
	p.Elems = append(p.Elems, &Elem{e, c})
	p.Idxs[e] = len(p.Elems) - 1
	p.swiftup(len(p.Elems) - 1)
}

func (p *PQ) Pop() int {
	e, l := p.Elems[0], p.Elems[len(p.Elems)-1]
	p.Elems[0], p.Elems[len(p.Elems)-1] = p.Elems[len(p.Elems)-1], p.Elems[0]
	p.Idxs[e.Val], p.Idxs[l.Val] = p.Idxs[l.Val], p.Idxs[e.Val]
	p.Elems = p.Elems[:len(p.Elems)-1]
	delete(p.Idxs, e.Val)
	p.swiftdown(0)
	return e.Val
}

func (p *PQ) IsEmpty() bool {
	return len(p.Elems) == 0
}

func (p *PQ) swiftup(i int) {
	for n := parent(i); i > 0 && p.Elems[n].Cost > p.Elems[i].Cost; i, n = n, parent(n) {
		p.Elems[i], p.Elems[n] = p.Elems[n], p.Elems[i]
		p.Idxs[p.Elems[i].Val], p.Idxs[p.Elems[n].Val] = p.Idxs[p.Elems[n].Val], p.Idxs[p.Elems[i].Val]
	}
}

func (p *PQ) swiftdown(i int) {
	n := len(p.Elems)
	for true {
		l, r := children(i)
		minIdx := i
		if l < n && p.Elems[l].Cost < p.Elems[minIdx].Cost {
			minIdx = l
		}
		if r < n && p.Elems[r].Cost < p.Elems[minIdx].Cost {
			minIdx = r
		}
		if i == minIdx {
			break
		}
		p.Elems[i], p.Elems[minIdx] = p.Elems[minIdx], p.Elems[i]
		p.Idxs[p.Elems[i].Val], p.Idxs[p.Elems[minIdx].Val] = p.Idxs[p.Elems[minIdx].Val], p.Idxs[p.Elems[i].Val]
		i = minIdx
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
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
	return &DirectedGraph{make(map[int]*DNode)}
}
