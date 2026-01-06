// Generated on 2026-01-02 15:01:37
// Daily practice file: unionfind_20260102.go

package unionfind

type UnionFind struct {
	Parents []int
	Ranks   []int
}

// --------------------------------------------------------------

func (u *UnionFind) Find(x int) int {
	root := u.Parents[x]
	if x != root {
		root = u.Find(root)
		u.Parents[x] = root
	}
	return root
}

// --------------------------------------------------------------

func (u *UnionFind) Union(x, y int) {
	p, q := u.Find(x), u.Find(y)
	if p == q {
		return
	}
	if u.Ranks[p] < u.Ranks[q] {
		u.Parents[p] = q
	} else if u.Ranks[p] > u.Ranks[q] {
		u.Parents[q] = p
	} else {
		u.Parents[p] = q
		u.Ranks[q]++
	}
}

// --------------------------------------------------------------

func (u *UnionFind) Connected(x, y int) bool {
	p, q := u.Find(x), u.Find(y)
	return p == q
}

// --------------------------------------------------------------

func (u *UnionFind) Count() int {
	roots := make(map[int]bool)
	for i := range u.Parents {
		roots[u.Parents[i]] = true
	}
	return len(roots)
}

// --------------------------------------------------------------

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := range n {
		parents[i] = i
	}
	return &UnionFind{parents, make([]int, n)}
}
