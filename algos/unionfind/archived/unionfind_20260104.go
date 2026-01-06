// Generated on 2026-01-04 08:38:04
// Daily practice file: unionfind_20260104.go

package unionfind

/*
 * The union-find data structure, also known as disjoint set union (DSU), is used to keep track of a partition of a set into disjoint subsets.
 * It supports two main operations:
 *  Find: This operation determines which subset a particular element belongs to.
 *        In other words, it returns the representative or the root of the set that contains the element.
 *  Union: This operation merges two subsets into a single subset.
 *         When you union two elements, you connect their sets so that they share the same representative.
 */
type UnionFind struct {
	Parents []int
	Ranks   []int
}

// --------------------------------------------------------------

/*
 * The find function in the union-find data structure is responsible for identifying the representative or the root of the set that a particular element belongs to.
 *
 * 1. Start at the Element
 * 2. Traverse Up the Tree:
 *    The function follows parent pointers from the element upwards until it reaches a node that is its own parent.
 *    In other words, you keep moving up until you find the root of the set.
 * 3. Path Compression (Optional but Common):
 *    As you traverse the path from the element to the root, you can perform path compression.
 *    This means that you update each node along the path to point directly to the root.
 *    This flattening of the structure makes future find operations faster by reducing the path length.
 *
 * 4. Return the Root:
 *    Once you reach the root, the find function returns that root element, which serves as the representative of the set.
 */
func (u *UnionFind) Find(x int) int {
	p := u.Parents[x]
	if p != x {
		p = u.Find(p)
		u.Parents[x] = p
	}
	return p
}

// --------------------------------------------------------------

/**
 * The union function in the union-find data structure is responsible for merging two sets into a single set.
 * 1. Find the Roots:
 *    To begin, you first use the find function on both elements you want to union.
 *    This gives you the root of each element’s set.
 * 2. Compare the Roots:
 *    Once you have the roots, you compare them to determine which one should become the parent of the other.
 * 3. Union by Rank or Size:
 *    To keep the structure balanced, the union is typically done by comparing the rank (or the size) of the trees.
 *    The root with the lower rank or smaller size is attached under the root of the root with the higher rank or larger size.
 *    This helps keep the overall tree balanced and prevents it from becoming too deep.
 * 4. Update the Parent:
 *    After determining which root should be the parent,
 *    you update the parent pointer of the root with the lower rank or size to point to the root with the higher rank or size.
 */
func (u *UnionFind) Union(x, y int) {
	p, q := u.Find(x), u.Find(y)
	if p == q {
		return
	}
	rankp, rankq := u.Ranks[p], u.Ranks[q]
	if rankp < rankq {
		u.Parents[p] = q
	} else if rankp > rankq {
		u.Parents[q] = p
	} else {
		u.Parents[p] = q
		u.Ranks[q]++
	}
}

// --------------------------------------------------------------

func (u *UnionFind) Connected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// --------------------------------------------------------------

func (u *UnionFind) Count() int {
	parents := make(map[int]bool)
	for i := range u.Parents {
		parents[u.Find(u.Parents[i])] = true
	}
	return len(parents)
}

// --------------------------------------------------------------

func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{make([]int, n), make([]int, n)}
	for i := range u.Parents {
		u.Parents[i] = i
	}
	return u
}
