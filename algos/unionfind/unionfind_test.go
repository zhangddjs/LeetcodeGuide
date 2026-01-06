package unionfind

import "testing"

type unionFindInterface interface {
	Find(x int) int
	Union(x, y int)
	Connected(x, y int) bool
	Count() int
}

var basicTestCases = []struct {
	name       string
	size       int
	operations []operation
	expected   expectedResults
}{
	{
		name: "single element",
		size: 1,
		operations: []operation{
			{op: "find", x: 0, expected: 0},
			{op: "count", expected: 1},
		},
	},
	{
		name: "two separate elements",
		size: 2,
		operations: []operation{
			{op: "connected", x: 0, y: 1, expected: false},
			{op: "count", expected: 2},
		},
	},
	{
		name: "basic union",
		size: 3,
		operations: []operation{
			{op: "union", x: 0, y: 1},
			{op: "connected", x: 0, y: 1, expected: true},
			{op: "connected", x: 1, y: 2, expected: false},
			{op: "count", expected: 2},
		},
	},
	{
		name: "chain unions",
		size: 4,
		operations: []operation{
			{op: "union", x: 0, y: 1},
			{op: "union", x: 1, y: 2},
			{op: "connected", x: 0, y: 2, expected: true},
			{op: "connected", x: 0, y: 3, expected: false},
			{op: "count", expected: 2},
		},
	},
	{
		name: "complex operations",
		size: 6,
		operations: []operation{
			{op: "union", x: 0, y: 1},
			{op: "union", x: 2, y: 3},
			{op: "union", x: 4, y: 5},
			{op: "count", expected: 3},
			{op: "union", x: 1, y: 3},
			{op: "connected", x: 0, y: 2, expected: true},
			{op: "connected", x: 0, y: 4, expected: false},
			{op: "count", expected: 2},
		},
	},
}

type operation struct {
	op       string
	x, y     int
	expected interface{}
}

type expectedResults struct {
	finalCount int
	groups     [][]int
}

func testUnionFindBasic(t *testing.T, newUF func(n int) unionFindInterface) {
	for _, tc := range basicTestCases {
		t.Run(tc.name, func(t *testing.T) {
			uf := newUF(tc.size)

			for i, op := range tc.operations {
				switch op.op {
				case "find":
					result := uf.Find(op.x)
					if expected, ok := op.expected.(int); ok && result != expected {
						t.Errorf("operation %d: Find(%d) = %d, want %d", i, op.x, result, expected)
					}
				case "union":
					uf.Union(op.x, op.y)
				case "connected":
					result := uf.Connected(op.x, op.y)
					if expected, ok := op.expected.(bool); ok && result != expected {
						t.Errorf("operation %d: Connected(%d, %d) = %v, want %v", i, op.x, op.y, result, expected)
					}
				case "count":
					result := uf.Count()
					if expected, ok := op.expected.(int); ok && result != expected {
						t.Errorf("operation %d: Count() = %d, want %d", i, result, expected)
					}
				}
			}
		})
	}
}

func testUnionFindPathCompression(t *testing.T, newUF func(n int) unionFindInterface) {
	uf := newUF(10)

	// Create a long chain: 0-1-2-3-4-5-6-7-8-9
	for i := 0; i < 9; i++ {
		uf.Union(i, i+1)
	}

	// All should be connected
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			if !uf.Connected(i, j) {
				t.Errorf("Expected %d and %d to be connected", i, j)
			}
		}
	}

	// Should have only 1 component
	if count := uf.Count(); count != 1 {
		t.Errorf("Count() = %d, want 1", count)
	}
}

func testUnionFindDisjointSets(t *testing.T, newUF func(n int) unionFindInterface) {
	uf := newUF(8)

	// Create two disjoint sets: {0,1,2,3} and {4,5,6,7}
	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(2, 3)

	uf.Union(4, 5)
	uf.Union(5, 6)
	uf.Union(6, 7)

	// Test within-set connections
	testConnections := []struct {
		x, y     int
		expected bool
	}{
		{0, 1, true}, {0, 2, true}, {0, 3, true},
		{1, 2, true}, {1, 3, true}, {2, 3, true},
		{4, 5, true}, {4, 6, true}, {4, 7, true},
		{5, 6, true}, {5, 7, true}, {6, 7, true},
		{0, 4, false}, {1, 5, false}, {2, 6, false}, {3, 7, false},
	}

	for _, tc := range testConnections {
		if result := uf.Connected(tc.x, tc.y); result != tc.expected {
			t.Errorf("Connected(%d, %d) = %v, want %v", tc.x, tc.y, result, tc.expected)
		}
	}

	if count := uf.Count(); count != 2 {
		t.Errorf("Count() = %d, want 2", count)
	}
}

func testUnionFindSequentialOperations(t *testing.T, newUF func(n int) unionFindInterface) {
	uf := newUF(5)

	// Initial state - all separate
	if count := uf.Count(); count != 5 {
		t.Errorf("Initial Count() = %d, want 5", count)
	}

	// Check not connected
	if connected := uf.Connected(0, 1); connected {
		t.Errorf("Connected(0, 1) = %v, want false", connected)
	}

	// Union first pair
	uf.Union(0, 1)

	// Check now connected
	if connected := uf.Connected(0, 1); !connected {
		t.Errorf("Connected(0, 1) after union = %v, want true", connected)
	}

	// Count should decrease
	if count := uf.Count(); count != 4 {
		t.Errorf("Count() after first union = %d, want 4", count)
	}

	// Check Find returns same root for connected elements
	root0 := uf.Find(0)
	root1 := uf.Find(1)
	if root0 != root1 {
		t.Errorf("Find(0) = %d, Find(1) = %d, want same root", root0, root1)
	}

	// Union another pair
	uf.Union(2, 3)

	// Count should decrease again
	if count := uf.Count(); count != 3 {
		t.Errorf("Count() after second union = %d, want 3", count)
	}

	// Elements 0,1 should not be connected to 2,3
	if connected := uf.Connected(0, 2); connected {
		t.Errorf("Connected(0, 2) = %v, want false", connected)
	}
}

func testUnionFindIdempotent(t *testing.T, newUF func(n int) unionFindInterface) {
	uf := newUF(3)

	// Union the same pair multiple times
	uf.Union(0, 1)
	count1 := uf.Count()

	uf.Union(0, 1) // Should have no effect
	count2 := uf.Count()

	uf.Union(1, 0) // Reverse order, should have no effect
	count3 := uf.Count()

	if count1 != count2 || count2 != count3 {
		t.Errorf("Counts after repeated unions: %d, %d, %d - should be same", count1, count2, count3)
	}

	if count1 != 2 {
		t.Errorf("Count() = %d, want 2", count1)
	}
}

// Uncomment and implement when union-find implementation is available
func TestUnionFind(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		testUnionFindBasic(t, func(n int) unionFindInterface { return NewUnionFind(n) })
	})
	t.Run("PathCompression", func(t *testing.T) {
		testUnionFindPathCompression(t, func(n int) unionFindInterface { return NewUnionFind(n) })
	})
	t.Run("DisjointSets", func(t *testing.T) {
		testUnionFindDisjointSets(t, func(n int) unionFindInterface { return NewUnionFind(n) })
	})
	t.Run("SequentialOperations", func(t *testing.T) {
		testUnionFindSequentialOperations(t, func(n int) unionFindInterface { return NewUnionFind(n) })
	})
	t.Run("Idempotent", func(t *testing.T) {
		testUnionFindIdempotent(t, func(n int) unionFindInterface { return NewUnionFind(n) })
	})
}
