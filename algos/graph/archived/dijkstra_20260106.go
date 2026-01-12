package graph

import "math"

type DirectedGraph struct {
	Nodes map[int]*DNode
}

type DNode struct {
	Val     int
	Adjs    map[int]*DNode
	Weights map[int]int
}

func (d *DirectedGraph) AddEdge(from, to int) {
	d.AddWeightedEdge(from, to, 1)
}

func (d *DirectedGraph) AddWeightedEdge(from, to, weight int) {
	nf, ok := d.Nodes[from]
	if !ok {
		nf = &DNode{from, make(map[int]*DNode), make(map[int]int)}
		d.Nodes[from] = nf
	}
	nt, ok := d.Nodes[to]
	if !ok {
		nt = &DNode{to, make(map[int]*DNode), make(map[int]int)}
		d.Nodes[to] = nt
	}
	nf.Adjs[to] = nt
	nf.Weights[to] = weight
}

func (d *DirectedGraph) ShortestPath(start int) map[int]int {
	costs := d.initialCost(start)
	pq := NewPriorityQueue()
	for i := range costs {
		pq.Push(i, costs[i])
	}
	for !pq.IsEmpty() {
		cur := pq.Pop()
		for i := range d.Nodes[cur].Adjs {
			co := costs[i]
			ci := d.relax(cur, i, d.Nodes[cur].Weights[i], costs)
			if ci < co {
				pq.Decrease(i, ci)
			}
		}
	}
	return costs
}

func (d *DirectedGraph) relax(i, j, cost int, costs map[int]int) int {
	ci, ok := costs[i]
	if !ok || ci == math.MaxInt {
		return math.MaxInt
	}
	cj := costs[j]
	if ci+cost < cj {
		costs[j] = ci + cost
	}
	return costs[j]
}

func (d *DirectedGraph) initialCost(start int) map[int]int {
	res := make(map[int]int)
	for i := range d.Nodes {
		res[i] = math.MaxInt
		if i == start {
			res[i] = 0
		}
	}
	return res
}

func (d *DirectedGraph) ShortestPathBetween(from, to int) (int, bool) {
	costs := d.ShortestPath(from)
	cost, ok := costs[to]
	if !ok || cost == math.MaxInt {
		return 0, false
	}
	return cost, true
}

func NewDijkstraGraph() *DirectedGraph {
	return &DirectedGraph{make(map[int]*DNode)}
}

type PriorityQueue struct {
	Elems []int
	Costs map[int]int
	Idx   map[int]int
}

func (p *PriorityQueue) Push(e, c int) {
	p.Elems = append(p.Elems, e)
	p.Costs[e] = c
	p.swiftUp(len(p.Elems) - 1)
}

func (p *PriorityQueue) Pop() int {
	e := p.Elems[0]
	p.Elems = p.Elems[1:]
	delete(p.Costs, e)
	return e
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.Elems) == 0
}

func (p *PriorityQueue) Decrease(e, c int) {
	i, ok := p.Idx[e]
	if !ok {
		return
	}
	pc := p.Costs[e]
	if pc <= c {
		return
	}
	p.Costs[e] = c
	p.swiftUp(i)
}

func (p *PriorityQueue) swiftUp(n int) {
	for n > 0 && p.Costs[p.Elems[n]] < p.Costs[p.Elems[parent(n)]] {
		p.Elems[parent(n)], p.Elems[n] = p.Elems[n], p.Elems[parent(n)]
		n = parent(n)
	}
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{make([]int, 0), make(map[int]int), make(map[int]int)}
}

func parent(n int) int {
	return (n - 1) / 2
}
