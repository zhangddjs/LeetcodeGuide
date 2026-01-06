package graph

import (
	"reflect"
	"sort"
	"testing"
)

type graphInterface interface {
	AddEdge(from, to int)
	DFS(start int) []int
	HasPath(from, to int) bool
	ShortestPath(from, to int) []int
}

var addEdgeTestCases = []struct {
	name  string
	edges [][]int
}{
	{"empty graph", [][]int{}},
	{"single edge", [][]int{{0, 1}}},
	{"multiple edges", [][]int{{0, 1}, {1, 2}, {2, 3}}},
	{"disconnected components", [][]int{{0, 1}, {2, 3}}},
	{"cycle", [][]int{{0, 1}, {1, 2}, {2, 0}}},
	{"self loop", [][]int{{0, 0}}},
	{"bidirectional edges", [][]int{{0, 1}, {1, 0}}},
	{"star pattern", [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}},
}

var dfsTestCases = []struct {
	name     string
	edges    [][]int
	start    int
	expected []int
}{
	{"empty graph", [][]int{}, 0, []int{}},
	{"single node", [][]int{{0, 0}}, 0, []int{0}},
	{"linear path", [][]int{{0, 1}, {1, 2}, {2, 3}}, 0, []int{0, 1, 2, 3}},
	{"tree structure", [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}, 0, []int{0, 1, 3, 4, 2}},
	{"disconnected start", [][]int{{0, 0}, {1, 2}, {2, 3}}, 0, []int{0}},
	{"cycle graph", [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, 0, []int{0, 1, 2, 3}},
	{"star from center", [][]int{{0, 1}, {0, 2}, {0, 3}}, 0, []int{0, 1, 2, 3}},
	{"star from leaf", [][]int{{0, 1}, {0, 2}, {0, 3}}, 1, []int{1, 0, 2, 3}},
}

var hasPathTestCases = []struct {
	name     string
	edges    [][]int
	from     int
	to       int
	expected bool
}{
	{"empty graph", [][]int{}, 0, 1, false},
	{"same node", [][]int{{0, 1}}, 0, 0, true},
	{"direct connection", [][]int{{0, 1}}, 0, 1, true},
	{"no connection", [][]int{{0, 1}}, 0, 2, false},
	{"indirect connection", [][]int{{0, 1}, {1, 2}}, 0, 2, true},
	{"disconnected components", [][]int{{0, 1}, {2, 3}}, 0, 2, false},
	{"cycle reachable", [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2, true},
	{"reverse path", [][]int{{0, 1}, {1, 2}}, 2, 0, true},
}

var shortestPathTestCases = []struct {
	name     string
	edges    [][]int
	from     int
	to       int
	expected []int
}{
	{"same node", [][]int{{0, 1}}, 0, 0, []int{0}},
	{"direct connection", [][]int{{0, 1}}, 0, 1, []int{0, 1}},
	{"no path", [][]int{{0, 1}}, 0, 2, []int{}},
	{"two hops", [][]int{{0, 1}, {1, 2}}, 0, 2, []int{0, 1, 2}},
	{"two paths", [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}}, 0, 3, []int{0, 1, 3}},
	// {"multiple paths", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, 0, 3, []int{0, 1, 3}},
	{"longer path", [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 0, 4, []int{0, 1, 2, 3, 4}},
	{"disconnected", [][]int{{0, 1}, {2, 3}}, 0, 3, []int{}},
}

func testGraphAddEdge(t *testing.T, newGraph func() graphInterface) {
	for _, tc := range addEdgeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			g := newGraph()
			for _, edge := range tc.edges {
				g.AddEdge(edge[0], edge[1])
			}
		})
	}
}

func testGraphDFS(t *testing.T, newGraph func() graphInterface) {
	for _, tc := range dfsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			g := newGraph()
			for _, edge := range tc.edges {
				g.AddEdge(edge[0], edge[1])
			}
			result := g.DFS(tc.start)

			// Sort both expected and result for comparison (DFS order can vary)
			sortedResult := make([]int, len(result))
			copy(sortedResult, result)
			sort.Ints(sortedResult)

			sortedExpected := make([]int, len(tc.expected))
			copy(sortedExpected, tc.expected)
			sort.Ints(sortedExpected)

			if !reflect.DeepEqual(sortedResult, sortedExpected) {
				t.Errorf("DFS(%d) = %v, want %v", tc.start, result, tc.expected)
			}
		})
	}
}

func testGraphHasPath(t *testing.T, newGraph func() graphInterface) {
	for _, tc := range hasPathTestCases {
		t.Run(tc.name, func(t *testing.T) {
			g := newGraph()
			for _, edge := range tc.edges {
				g.AddEdge(edge[0], edge[1])
			}
			result := g.HasPath(tc.from, tc.to)
			if result != tc.expected {
				t.Errorf("HasPath(%d, %d) = %v, want %v", tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

func testGraphShortestPath(t *testing.T, newGraph func() graphInterface) {
	for _, tc := range shortestPathTestCases {
		t.Run(tc.name, func(t *testing.T) {
			g := newGraph()
			for _, edge := range tc.edges {
				g.AddEdge(edge[0], edge[1])
			}
			result := g.ShortestPath(tc.from, tc.to)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ShortestPath(%d, %d) = %v, want %v", tc.from, tc.to, result, tc.expected)
			}
		})
	}
}

func testGraphComprehensive(t *testing.T, newGraph func() graphInterface) {
	g := newGraph()

	// Build a more complex graph
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 5}, {5, 6}}
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	// Test BFS from different starting points
	bfsTests := map[int]int{
		0: 7, // Should visit all 7 nodes
		3: 7, // Should visit nodes 3, 1, 0, 4, 2, 5, 6, and possibly others
		6: 7, // Should only visit node 6
	}

	for start, expectedCount := range bfsTests {
		result := g.DFS(start)
		if len(result) != expectedCount {
			t.Errorf("DFS(%d) visited %d nodes, want %d", start, len(result), expectedCount)
		}
	}

	// Test path existence
	pathTests := map[string]bool{
		"0->6": true,
		"0->3": true,
		"6->0": true,
		"3->2": true,
		"1->2": true,
		"2->3": true,
	}

	pathMapping := map[string][]int{
		"0->6": {0, 6},
		"0->3": {0, 3},
		"6->0": {6, 0},
		"3->2": {3, 2},
		"1->2": {1, 2},
		"2->3": {2, 3},
	}

	for pathStr, expected := range pathTests {
		path := pathMapping[pathStr]
		if result := g.HasPath(path[0], path[1]); result != expected {
			t.Errorf("HasPath(%d, %d) = %v, want %v", path[0], path[1], result, expected)
		}
	}
}

func testGraphSequentialOperations(t *testing.T, newGraph func() graphInterface) {
	g := newGraph()

	// Initial BFS on empty graph
	if result := g.DFS(0); len(result) != 0 {
		t.Errorf("BFS(0) on empty graph = %v, want empty", result)
	}

	// Add first edge
	g.AddEdge(0, 1)

	// Test after first edge
	result := g.DFS(0)
	expectedNodes := []int{0, 1}
	if len(result) != len(expectedNodes) {
		t.Errorf("DFS(0) after first edge = %v, want %d nodes", result, len(expectedNodes))
	}

	// Test path existence
	if !g.HasPath(0, 1) {
		t.Errorf("HasPath(0, 1) = false, want true")
	}
	if !g.HasPath(1, 0) {
		t.Errorf("HasPath(1, 0) = false, want true (assuming directed)")
	}

	// Add more edges
	g.AddEdge(1, 2)
	g.AddEdge(0, 3)

	// Test BFS covers all reachable nodes
	result = g.DFS(0)
	if len(result) < 3 {
		t.Errorf("DFS(0) after adding more edges = %v, should reach at least 3 nodes", result)
	}

	// Test shortest path
	path := g.ShortestPath(0, 2)
	if len(path) != 3 { // 0 -> 1 -> 2
		t.Errorf("ShortestPath(0, 2) = %v, want path of length 3", path)
	}
}

// Uncomment and implement when graph implementation is available
func TestGraph(t *testing.T) {
	t.Run("AddEdge", func(t *testing.T) {
		testGraphAddEdge(t, func() graphInterface { return NewGraph() })
	})
	t.Run("DFS", func(t *testing.T) {
		testGraphDFS(t, func() graphInterface { return NewGraph() })
	})
	t.Run("HasPath", func(t *testing.T) {
		testGraphHasPath(t, func() graphInterface { return NewGraph() })
	})
	t.Run("ShortestPath", func(t *testing.T) {
		testGraphShortestPath(t, func() graphInterface { return NewGraph() })
	})
	t.Run("Comprehensive", func(t *testing.T) {
		testGraphComprehensive(t, func() graphInterface { return NewGraph() })
	})
	t.Run("SequentialOperations", func(t *testing.T) {
		testGraphSequentialOperations(t, func() graphInterface { return NewGraph() })
	})
}
