package graph

import (
	"math"
	"reflect"
	"testing"
)

type dijkstraInterface interface {
	AddEdge(from, to int)
	AddWeightedEdge(from, to, weight int)
	ShortestPath(start int) (distances map[int]int)
	ShortestPathBetween(from, to int) (distance int, found bool)
}

var dijkstraAddEdgeTestCases = []struct {
	name  string
	edges [][]int
}{
	{"empty graph", [][]int{}},
	{"single edge", [][]int{{0, 1}}},
	{"linear chain", [][]int{{0, 1}, {1, 2}, {2, 3}}},
	{"star structure", [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}},
	{"cycle", [][]int{{0, 1}, {1, 2}, {2, 0}}},
	{"bidirectional", [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}}},
	{"complex graph", [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {2, 4}, {3, 5}, {4, 5}, {5, 0}}},
}

var dijkstraWeightedEdgeTestCases = []struct {
	name  string
	edges [][]int // [from, to, weight]
}{
	{"empty graph", [][]int{}},
	{"single edge", [][]int{{0, 1, 5}}},
	{"multiple edges", [][]int{{0, 1, 4}, {1, 2, 3}, {0, 2, 8}}},
	{"complex weights", [][]int{{0, 1, 2}, {0, 2, 6}, {1, 2, 3}, {1, 3, 1}, {2, 3, 1}}},
	{"cyclic with weights", [][]int{{0, 1, 4}, {1, 2, 3}, {2, 0, 2}}},
}

var dijkstraShortestPathTestCases = []struct {
	name     string
	edges    [][]int // [from, to, weight]
	start    int
	expected map[int]int
}{
	{"empty graph", [][]int{}, 0, map[int]int{}},
	{"single edge", [][]int{{0, 1, 5}}, 0, map[int]int{0: 0, 1: 5}},
	{"linear path", [][]int{{0, 1, 4}, {1, 2, 3}}, 0, map[int]int{0: 0, 1: 4, 2: 7}},
	{"multiple paths", [][]int{{0, 1, 2}, {0, 2, 8}, {1, 2, 3}}, 0, map[int]int{0: 0, 1: 2, 2: 5}},
	{"triangle", [][]int{{0, 1, 4}, {0, 2, 2}, {1, 2, 1}}, 0, map[int]int{0: 0, 1: 4, 2: 2}},
	{"diamond pattern", [][]int{{0, 1, 3}, {0, 2, 8}, {1, 3, 2}, {2, 3, 2}}, 0, map[int]int{0: 0, 1: 3, 2: 8, 3: 5}},
	{"complex graph", [][]int{{0, 1, 4}, {0, 2, 2}, {1, 2, 1}, {1, 3, 5}, {2, 3, 8}, {2, 4, 3}, {3, 4, 1}}, 0, map[int]int{0: 0, 1: 4, 2: 2, 3: 9, 4: 5}},
	{"disconnected component", [][]int{{0, 1, 3}, {2, 3, 2}}, 0, map[int]int{0: 0, 1: 3, 2: math.MaxInt, 3: math.MaxInt}},
	{"start from middle", [][]int{{0, 1, 2}, {1, 2, 3}, {1, 3, 1}}, 1, map[int]int{0: math.MaxInt, 1: 0, 2: 3, 3: 1}},
	{"unreachable nodes", [][]int{{0, 1, 3}, {2, 3, 2}, {4, 5, 1}}, 0, map[int]int{0: 0, 1: 3, 2: math.MaxInt, 3: math.MaxInt, 4: math.MaxInt, 5: math.MaxInt}},
	{"cycle with better path", [][]int{{0, 1, 10}, {0, 2, 5}, {2, 1, 3}}, 0, map[int]int{0: 0, 1: 8, 2: 5}},
	{"bidirectional edges", [][]int{{0, 1, 4}, {1, 0, 2}, {1, 2, 3}, {2, 1, 1}}, 0, map[int]int{0: 0, 1: 4, 2: 7}},
}

var dijkstraShortestPathBetweenTestCases = []struct {
	name     string
	edges    [][]int // [from, to, weight]
	from     int
	to       int
	expected int
	found    bool
}{
	{"empty graph", [][]int{}, 0, 1, 0, false},
	{"single edge - exists", [][]int{{0, 1, 5}}, 0, 1, 5, true},
	{"single edge - reverse", [][]int{{0, 1, 5}}, 1, 0, 0, false},
	{"same node", [][]int{{0, 1, 5}}, 0, 0, 0, true},
	{"linear path", [][]int{{0, 1, 4}, {1, 2, 3}}, 0, 2, 7, true},
	{"multiple paths", [][]int{{0, 1, 2}, {0, 2, 8}, {1, 2, 3}}, 0, 2, 5, true},
	{"no path exists", [][]int{{0, 1, 5}, {2, 3, 3}}, 0, 3, 0, false},
	{"complex shortest path", [][]int{{0, 1, 4}, {0, 2, 2}, {1, 3, 1}, {2, 3, 5}, {2, 4, 3}, {3, 4, 1}}, 0, 4, 5, true},
}

func testDijkstraAddEdge(t *testing.T, newGraph func() dijkstraInterface) {
	for _, tc := range dijkstraAddEdgeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			graph := newGraph()
			for _, edge := range tc.edges {
				graph.AddEdge(edge[0], edge[1])
			}
		})
	}
}

func testDijkstraAddWeightedEdge(t *testing.T, newGraph func() dijkstraInterface) {
	for _, tc := range dijkstraWeightedEdgeTestCases {
		t.Run(tc.name, func(t *testing.T) {
			graph := newGraph()
			for _, edge := range tc.edges {
				if len(edge) >= 3 {
					graph.AddWeightedEdge(edge[0], edge[1], edge[2])
				}
			}
		})
	}
}

func testDijkstraShortestPath(t *testing.T, newGraph func() dijkstraInterface) {
	for _, tc := range dijkstraShortestPathTestCases {
		t.Run(tc.name, func(t *testing.T) {
			graph := newGraph()
			for _, edge := range tc.edges {
				if len(edge) >= 3 {
					graph.AddWeightedEdge(edge[0], edge[1], edge[2])
				}
			}

			result := graph.ShortestPath(tc.start)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ShortestPath(%d) = %v, want %v", tc.start, result, tc.expected)
			}
		})
	}
}

func testDijkstraShortestPathBetween(t *testing.T, newGraph func() dijkstraInterface) {
	for _, tc := range dijkstraShortestPathBetweenTestCases {
		t.Run(tc.name, func(t *testing.T) {
			graph := newGraph()
			for _, edge := range tc.edges {
				if len(edge) >= 3 {
					graph.AddWeightedEdge(edge[0], edge[1], edge[2])
				}
			}

			distance, found := graph.ShortestPathBetween(tc.from, tc.to)
			if distance != tc.expected || found != tc.found {
				t.Errorf("ShortestPathBetween(%d, %d) = (%d, %v), want (%d, %v)",
					tc.from, tc.to, distance, found, tc.expected, tc.found)
			}
		})
	}
}

func testDijkstraComprehensive(t *testing.T, newGraph func() dijkstraInterface) {
	graph := newGraph()

	// Build a complex weighted graph
	edges := [][]int{{0, 1, 4}, {0, 2, 2}, {1, 2, 1}, {1, 3, 5}, {2, 3, 8}, {2, 4, 3}, {3, 4, 1}}
	for _, edge := range edges {
		graph.AddWeightedEdge(edge[0], edge[1], edge[2])
	}

	// Test shortest paths from node 0
	distances := graph.ShortestPath(0)
	expected := map[int]int{0: 0, 1: 4, 2: 2, 3: 9, 4: 5}
	if !reflect.DeepEqual(distances, expected) {
		t.Errorf("ShortestPath(0) = %v, want %v", distances, expected)
	}

	// Test specific path
	distance, found := graph.ShortestPathBetween(0, 4)
	if !found || distance != 5 {
		t.Errorf("ShortestPathBetween(0, 4) = (%d, %v), want (5, true)", distance, found)
	}
}

func testDijkstraSequentialOperations(t *testing.T, newGraph func() dijkstraInterface) {
	graph := newGraph()

	// Start with empty graph
	emptyDistances := graph.ShortestPath(0)
	if len(emptyDistances) != 0 {
		t.Errorf("ShortestPath(0) on empty graph = %v, want empty map", emptyDistances)
	}

	// Add first edge
	graph.AddWeightedEdge(0, 1, 5)
	distances := graph.ShortestPath(0)
	expected := map[int]int{0: 0, 1: 5}
	if !reflect.DeepEqual(distances, expected) {
		t.Errorf("After adding first edge, ShortestPath(0) = %v, want %v", distances, expected)
	}

	// Add more edges to create triangle
	graph.AddWeightedEdge(0, 2, 3)
	graph.AddWeightedEdge(1, 2, 1)
	finalDistances := graph.ShortestPath(0)
	finalExpected := map[int]int{0: 0, 1: 5, 2: 3}
	if !reflect.DeepEqual(finalDistances, finalExpected) {
		t.Errorf("After adding triangle, ShortestPath(0) = %v, want %v", finalDistances, finalExpected)
	}
}

// Uncomment and implement when Dijkstra implementation is available
func TestDijkstra(t *testing.T) {
	t.Run("AddEdge", func(t *testing.T) {
		testDijkstraAddEdge(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
	t.Run("AddWeightedEdge", func(t *testing.T) {
		testDijkstraAddWeightedEdge(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
	t.Run("ShortestPath", func(t *testing.T) {
		testDijkstraShortestPath(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
	t.Run("ShortestPathBetween", func(t *testing.T) {
		testDijkstraShortestPathBetween(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
	t.Run("Comprehensive", func(t *testing.T) {
		testDijkstraComprehensive(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
	t.Run("SequentialOperations", func(t *testing.T) {
		testDijkstraSequentialOperations(t, func() dijkstraInterface { return NewDijkstraGraph() })
	})
}
