// Generated on 2026-01-07 08:33:23
// Daily practice file: dijkstra_20260107.go (dijkstra template)

package graph

import "math"

type DirectedGraph struct {
	Nodes map[int]*DNode
}

type DNode struct {
	Adjs   map[int]*DNode
	Weight map[int]int
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
	nf.Weight[to] = weight
}

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	costs := d.initialCosts(start)
	q := NewPriorityQueue()
	for i, c := range costs {
		q.Push(i, c)
	}
	for !q.IsEmpty() {
		e := q.Pop()
		for j := range d.Nodes[e].Adjs {
			w := d.Nodes[e].Weight[j]
			c := d.relax(e, j, w, costs)
			q.Decrease(j, c)
		}
	}

	return costs
}

func (d *DirectedGraph) relax(i, j, w int, costs map[int]int) int {
	c := math.MaxInt
	ci, ok := costs[i]
	if !ok || ci == c {
		return c
	}
	if ci+w < costs[j] {
		costs[j] = ci + w
	}
	return costs[j]
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

type PriorityQueue struct {
	Elems []int
	Costs map[int]int
	Idxs  map[int]int
}

func (p *PriorityQueue) Push(x, c int) {
	p.Elems = append(p.Elems, x)
	p.Costs[x] = c
	i := p.swiftup(len(p.Elems) - 1)
	p.Idxs[x] = i
}

func (p *PriorityQueue) Pop() int {
	e := p.Elems[0]
	p.Elems = p.Elems[1:]
	delete(p.Costs, e)
	delete(p.Idxs, e)
	// Update indices for remaining elements
	for i, elem := range p.Elems {
		p.Idxs[elem] = i
	}
	return e
}

func (p *PriorityQueue) Decrease(x, c int) {
	oc, ok := p.Costs[x]
	if !ok || c > oc {
		return
	}
	i := p.Idxs[x]
	p.Costs[x] = c
	p.swiftup(i)
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.Elems) == 0
}

func (p *PriorityQueue) swiftup(i int) int {
	for j := parent(i); i > 0 && p.Costs[p.Elems[j]] > p.Costs[p.Elems[i]]; {
		p.Elems[j], p.Elems[i] = p.Elems[i], p.Elems[j]
		i = j
		j = parent(i)
	}
	return i
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{make([]int, 0), make(map[int]int), make(map[int]int)}
}

func parent(i int) int {
	return (i - 1) / 2
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
	return &DirectedGraph{make(map[int]*DNode)}
}
