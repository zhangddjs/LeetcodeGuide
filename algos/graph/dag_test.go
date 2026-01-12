package graph

import (
	"math"
	"reflect"
	"testing"
)

type dagInterface interface {
	AddEdge(from, to int)
	AddWeightedEdge(from, to, weight int)
	TopologicalSort() []int    //https://www.geeksforgeeks.org/dsa/topological-sort-using-dfs/
	TopologicalSortBFS() []int // https://www.geeksforgeeks.org/dsa/topological-sorting-indegree-based-solution/
	IsDAG() bool
	ShortestPath(start int) (distances map[int]int)
}

var dagAddEdgeTestCases = []struct {
	name  string
	edges [][]int
}{
	{"empty DAG", [][]int{}},
	{"single edge", [][]int{{0, 1}}},
	{"linear chain", [][]int{{0, 1}, {1, 2}, {2, 3}}},
	{"tree structure", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}}},
	{"diamond pattern", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
	{"multiple sources", [][]int{{0, 2}, {1, 2}, {2, 3}}},
	{"complex DAG", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {4, 5}}},
}

var topSortTestCases = []struct {
	name     string
	edges    [][]int
	expected [][]int // Multiple valid orderings
}{
	{"empty DAG", [][]int{}, [][]int{{}}},
	{"single node", [][]int{}, [][]int{{0}}},
	{"single edge", [][]int{{0, 1}}, [][]int{{0, 1}}},
	{"linear chain", [][]int{{0, 1}, {1, 2}, {2, 3}}, [][]int{{0, 1, 2, 3}}},
	{"tree structure", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}},
		[][]int{{0, 1, 2, 3, 4}, {0, 2, 1, 4, 3}, {0, 1, 3, 2, 4}, {0, 2, 4, 1, 3}}},
	{"diamond pattern", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}},
		[][]int{{0, 1, 2, 3}, {0, 2, 1, 3}}},
	{"multiple sources", [][]int{{0, 2}, {1, 2}, {2, 3}},
		[][]int{{0, 1, 2, 3}, {1, 0, 2, 3}}},
}

var isDAGTestCases = []struct {
	name     string
	edges    [][]int
	expected bool
}{
	{"empty graph", [][]int{}, true},
	{"single edge", [][]int{{0, 1}}, true},
	{"linear chain", [][]int{{0, 1}, {1, 2}, {2, 3}}, true},
	{"tree structure", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 4}}, true},
	{"diamond pattern", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, true},
	{"simple cycle", [][]int{{0, 1}, {1, 2}, {2, 0}}, false},
	{"self loop", [][]int{{0, 0}}, false},
	{"complex cycle", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 1}}, false},
}

var weightedEdgeTestCases = []struct {
	name  string
	edges [][]int // [from, to, weight]
}{
	{"empty graph", [][]int{}},
	{"single edge", [][]int{{0, 1, 5}}},
	{"multiple edges", [][]int{{0, 1, 4}, {1, 2, 3}, {0, 2, 8}}},
	{"complex weights", [][]int{{0, 1, 2}, {0, 2, 6}, {1, 2, 3}, {1, 3, 1}, {2, 3, 1}}},
}

var dagshortestPathTestCases = []struct {
	name     string
	edges    [][]int // [from, to, weight]
	start    int
	expected map[int]int
}{
	{"empty graph", [][]int{}, 0, map[int]int{}},
	{"empty graph", [][]int{}, 1, map[int]int{}},
	{"single edge", [][]int{{0, 1, 5}}, 0, map[int]int{0: 0, 1: 5}},
	{"linear path", [][]int{{0, 1, 4}, {1, 2, 3}}, 0, map[int]int{0: 0, 1: 4, 2: 7}},
	{"multiple paths", [][]int{{0, 1, 2}, {0, 2, 8}, {1, 2, 3}}, 0, map[int]int{0: 0, 1: 2, 2: 5}},
	{"diamond pattern", [][]int{{0, 1, 2}, {0, 2, 6}, {1, 3, 1}, {2, 3, 1}}, 0, map[int]int{0: 0, 1: 2, 2: 6, 3: 3}},
	{"complex DAG", [][]int{{0, 1, 4}, {0, 2, 2}, {1, 3, 1}, {2, 3, 5}, {2, 4, 3}, {3, 4, 1}}, 0, map[int]int{0: 0, 1: 4, 2: 2, 3: 5, 4: 5}},
	{"disconnected component", [][]int{{0, 1, 3}, {2, 3, 2}}, 0, map[int]int{0: 0, 1: 3, 2: math.MaxInt, 3: math.MaxInt}},
	{"start from middle", [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 1}}, 1, map[int]int{0: math.MaxInt, 1: 0, 2: 3, 3: 1}},
	{"unreachable nodes", [][]int{{0, 1, 3}, {2, 3, 2}, {4, 5, 1}}, 0, map[int]int{0: 0, 1: 3, 2: math.MaxInt, 3: math.MaxInt, 4: math.MaxInt, 5: math.MaxInt}},
	{"negative weights", [][]int{{0, 1, -2}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2}, {2, 3, -5}}, 0, map[int]int{0: 0, 1: -2, 2: 1, 3: -4}},
}

func testDAGAddEdge(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range dagAddEdgeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				dag.AddEdge(edge[0], edge[1])
			}
		})
	}
}

func testDAGAddWeightedEdge(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range weightedEdgeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				if len(edge) >= 3 {
					dag.AddWeightedEdge(edge[0], edge[1], edge[2])
				}
			}
		})
	}
}

func testDAGTopologicalSort(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range topSortTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				dag.AddEdge(edge[0], edge[1])
			}
			result := dag.TopologicalSort()

			// Check if result matches any of the expected valid orderings
			valid := false
			for _, expected := range tc.expected {
				if reflect.DeepEqual(result, expected) {
					valid = true
					break
				}
			}

			if !valid && len(tc.expected) > 0 {
				// Additional check: verify topological ordering property
				if isValidTopologicalOrder(result, tc.edges) {
					valid = true
				}
			}

			if !valid {
				t.Errorf("TopologicalSort() = %v, want one of %v", result, tc.expected)
			}
		})
	}
}

func testDAGTopologicalSortBFS(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range topSortTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				dag.AddEdge(edge[0], edge[1])
			}
			result := dag.TopologicalSortBFS()

			// Check if result matches any of the expected valid orderings
			valid := false
			for _, expected := range tc.expected {
				if reflect.DeepEqual(result, expected) {
					valid = true
					break
				}
			}

			if !valid && len(tc.expected) > 0 {
				// Additional check: verify topological ordering property
				if isValidTopologicalOrder(result, tc.edges) {
					valid = true
				}
			}

			if !valid {
				t.Errorf("TopologicalSortBFS() = %v, want one of %v", result, tc.expected)
			}
		})
	}
}

func testDAGIsDAG(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range isDAGTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				dag.AddEdge(edge[0], edge[1])
			}
			result := dag.IsDAG()
			if result != tc.expected {
				t.Errorf("IsDAG() = %v, want %v", result, tc.expected)
			}
		})
	}
}

func testDAGShortestPath(t *testing.T, newDAG func() dagInterface) {
	for _, tc := range dagshortestPathTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dag := newDAG()
			for _, edge := range tc.edges {
				if len(edge) >= 3 {
					dag.AddWeightedEdge(edge[0], edge[1], edge[2])
				}
			}

			// if len(tc.edges) == 0 && tc.start == 0 {
			// 	dag.AddEdge(0, 0)
			// }

			result := dag.ShortestPath(tc.start)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ShortestPath(%d) = %v, want %v", tc.start, result, tc.expected)
			}
		})
	}
}

func testDAGComprehensive(t *testing.T, newDAG func() dagInterface) {
	dag := newDAG()

	// Build a complex DAG
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {4, 5}}
	for _, edge := range edges {
		dag.AddEdge(edge[0], edge[1])
	}

	// Verify it's a DAG
	if !dag.IsDAG() {
		t.Errorf("IsDAG() = false, want true for test DAG")
	}

	// Test topological sort exists
	topOrder := dag.TopologicalSort()
	if len(topOrder) == 0 {
		t.Errorf("TopologicalSort() returned empty, want valid ordering")
	}

	// Verify topological ordering property
	if !isValidTopologicalOrder(topOrder, edges) {
		t.Errorf("TopologicalSort() = %v, not a valid topological ordering", topOrder)
	}
}

func testDAGSequentialOperations(t *testing.T, newDAG func() dagInterface) {
	dag := newDAG()

	// Initial state - empty DAG
	if !dag.IsDAG() {
		t.Errorf("IsDAG() = false for empty graph, want true")
	}

	// Add first edge
	dag.AddWeightedEdge(0, 1, 5)

	// Verify still DAG
	if !dag.IsDAG() {
		t.Errorf("IsDAG() = false after adding edge, want true")
	}

	// Test topological sort
	topOrder := dag.TopologicalSort()
	if len(topOrder) < 2 {
		t.Errorf("TopologicalSort() = %v, want at least 2 nodes", topOrder)
	}

	// Add more edges to create diamond pattern
	dag.AddWeightedEdge(0, 2, 3)
	dag.AddWeightedEdge(1, 3, 2)
	dag.AddWeightedEdge(2, 3, 1)

	// Test final topological sort
	finalTopOrder := dag.TopologicalSort()
	if len(finalTopOrder) < 4 {
		t.Errorf("TopologicalSort() after adding edges = %v, want at least 4 nodes", finalTopOrder)
	}

	// Test that adding a cycle would make it not a DAG
	// Note: This test assumes your implementation can detect cycles when adding edges
	// If your implementation allows cycles to be added, you might need to modify this test
}

// Helper function to validate topological ordering
func isValidTopologicalOrder(order []int, edges [][]int) bool {
	position := make(map[int]int)
	for i, node := range order {
		position[node] = i
	}

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if position[from] >= position[to] {
			return false
		}
	}
	return true
}

// Uncomment and implement when DAG implementation is available
func TestDAG(t *testing.T) {
	t.Run("AddEdge", func(t *testing.T) {
		testDAGAddEdge(t, func() dagInterface { return NewDAG() })
	})
	t.Run("AddWeightedEdge", func(t *testing.T) {
		testDAGAddWeightedEdge(t, func() dagInterface { return NewDAG() })
	})
	t.Run("TopologicalSort", func(t *testing.T) {
		testDAGTopologicalSort(t, func() dagInterface { return NewDAG() })
	})
	t.Run("TopologicalSortBFS", func(t *testing.T) {
		testDAGTopologicalSortBFS(t, func() dagInterface { return NewDAG() })
	})
	t.Run("IsDAG", func(t *testing.T) {
		testDAGIsDAG(t, func() dagInterface { return NewDAG() })
	})
	t.Run("ShortestPath", func(t *testing.T) {
		testDAGShortestPath(t, func() dagInterface { return NewDAG() })
	})
	t.Run("Comprehensive", func(t *testing.T) {
		testDAGComprehensive(t, func() dagInterface { return NewDAG() })
	})
	t.Run("SequentialOperations", func(t *testing.T) {
		testDAGSequentialOperations(t, func() dagInterface { return NewDAG() })
	})
}
