package unionfind

type UnionFind struct {
	Parents []int
	Rank    []int
}

func (u *UnionFind) Find(x int) int {
	p := u.Parents[x]
	if p != x {
		p = u.Find(p)
		u.Parents[x] = p
	}
	return p
}

func (u *UnionFind) Union(x, y int) {
	p, q := u.Find(x), u.Find(y)
	rankp, rankq := u.Rank[p], u.Rank[q]
	if p == q {
		return
	}
	if rankp < rankq {
		u.Parents[p] = q
	} else if rankp > rankq {
		u.Parents[q] = p
	} else {
		u.Parents[p] = q
		u.Rank[q]++
	}
}

func (u *UnionFind) Connected(x, y int) bool {
	p, q := u.Find(x), u.Find(y)
	return p == q
}

func (u *UnionFind) Count() int {
	ds := make(map[int]bool)
	for i := range u.Parents {
		ds[u.Parents[i]] = true
	}
	return len(ds)
}

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{make([]int, n), make([]int, n)}
	for i := range n {
		u.Parents[i] = i
	}
	return u
}
