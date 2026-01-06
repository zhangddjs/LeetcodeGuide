package unionfind

type UnionFind struct {
	Rank   []int
	Parent []int
}

func (u *UnionFind) Find(x int) int {
	y := u.Parent[x]
	if x != y {
		u.Parent[x] = u.Find(y)
		return u.Parent[x]
	}
	return y
}

func (u *UnionFind) Union(x, y int) {
	px, py := u.Find(x), u.Find(y)
	if px == py {
		return
	}
	if u.Rank[px] < u.Rank[py] {
		u.Parent[px] = py
	} else if u.Rank[px] > u.Rank[py] {
		u.Parent[py] = px
	} else {
		u.Parent[py] = px
		u.Rank[py]++
	}
}

func (u *UnionFind) Connected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

func (u *UnionFind) Count() int {
	parents := make(map[int]int, len(u.Parent))
	for i := range u.Parent {
		parents[u.Find(u.Parent[i])]++
	}
	return len(parents)
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}
	return &UnionFind{make([]int, n), parent}
}
