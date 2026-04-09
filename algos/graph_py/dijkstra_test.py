import unittest
import sys
import os
import importlib.util
from typing import Protocol, Dict, Tuple


class DijkstraProtocol(Protocol):
    def add_edge(self, from_node: int, to_node: int) -> None: ...
    def add_weighted_edge(self, from_node: int, to_node: int, weight: int) -> None: ...
    def shortest_path(self, start: int) -> Dict[int, int]: ...
    def shortest_path_between(self, from_node: int, to_node: int) -> Tuple[int, bool]: ...


def load_dijkstra_module():
    """Dynamically load the most recent dijkstra_*.py file"""
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    files = [f for f in os.listdir(parent_dir) if f.startswith('dijkstra_') and f.endswith('.py')]

    if not files:
        raise FileNotFoundError("No dijkstra_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(parent_dir, latest_file)

    spec = importlib.util.spec_from_file_location("dijkstra_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module.DirectedGraph


class TestDijkstra(unittest.TestCase):

    def setUp(self):
        self.DirectedGraph = load_dijkstra_module()

    def test_add_edge_empty_graph(self):
        g = self.DirectedGraph()
        # Just verify it doesn't crash

    def test_add_edge_single_edge(self):
        g = self.DirectedGraph()
        g.add_edge(0, 1)

    def test_add_edge_linear_chain(self):
        g = self.DirectedGraph()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])

    def test_add_weighted_edge_single_edge(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)

    def test_add_weighted_edge_multiple_edges(self):
        g = self.DirectedGraph()
        edges = [[0, 1, 4], [1, 2, 3], [0, 2, 8]]
        for edge in edges:
            g.add_weighted_edge(edge[0], edge[1], edge[2])

    def test_shortest_path_empty_graph(self):
        g = self.DirectedGraph()
        result = g.shortest_path(0)
        self.assertEqual(result, {})

    def test_shortest_path_single_edge(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 5})

    def test_shortest_path_linear_path(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 4)
        g.add_weighted_edge(1, 2, 3)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 4, 2: 7})

    def test_shortest_path_multiple_paths(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 2)
        g.add_weighted_edge(0, 2, 8)
        g.add_weighted_edge(1, 2, 3)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 2, 2: 5})

    def test_shortest_path_triangle(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 4)
        g.add_weighted_edge(0, 2, 2)
        g.add_weighted_edge(1, 2, 1)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 4, 2: 2})

    def test_shortest_path_diamond_pattern(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 3)
        g.add_weighted_edge(0, 2, 8)
        g.add_weighted_edge(1, 3, 2)
        g.add_weighted_edge(2, 3, 2)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 3, 2: 8, 3: 5})

    def test_shortest_path_complex_graph(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 4)
        g.add_weighted_edge(0, 2, 2)
        g.add_weighted_edge(1, 2, 1)
        g.add_weighted_edge(1, 3, 5)
        g.add_weighted_edge(2, 3, 8)
        g.add_weighted_edge(2, 4, 3)
        g.add_weighted_edge(3, 4, 1)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 4, 2: 2, 3: 9, 4: 5})

    def test_shortest_path_disconnected_component(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 3)
        g.add_weighted_edge(2, 3, 2)
        result = g.shortest_path(0)
        expected = {0: 0, 1: 3, 2: float('inf'), 3: float('inf')}
        self.assertEqual(result, expected)

    def test_shortest_path_cycle_with_better_path(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 10)
        g.add_weighted_edge(0, 2, 5)
        g.add_weighted_edge(2, 1, 3)
        result = g.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 8, 2: 5})

    def test_shortest_path_between_empty_graph(self):
        g = self.DirectedGraph()
        distance, found = g.shortest_path_between(0, 1)
        self.assertFalse(found)

    def test_shortest_path_between_single_edge_exists(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)
        distance, found = g.shortest_path_between(0, 1)
        self.assertTrue(found)
        self.assertEqual(distance, 5)

    def test_shortest_path_between_single_edge_reverse(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)
        distance, found = g.shortest_path_between(1, 0)
        self.assertFalse(found)

    def test_shortest_path_between_same_node(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)
        distance, found = g.shortest_path_between(0, 0)
        self.assertTrue(found)
        self.assertEqual(distance, 0)

    def test_shortest_path_between_linear_path(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 4)
        g.add_weighted_edge(1, 2, 3)
        distance, found = g.shortest_path_between(0, 2)
        self.assertTrue(found)
        self.assertEqual(distance, 7)

    def test_shortest_path_between_multiple_paths(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 2)
        g.add_weighted_edge(0, 2, 8)
        g.add_weighted_edge(1, 2, 3)
        distance, found = g.shortest_path_between(0, 2)
        self.assertTrue(found)
        self.assertEqual(distance, 5)

    def test_shortest_path_between_no_path_exists(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 5)
        g.add_weighted_edge(2, 3, 3)
        distance, found = g.shortest_path_between(0, 3)
        self.assertFalse(found)

    def test_shortest_path_between_complex_shortest_path(self):
        g = self.DirectedGraph()
        g.add_weighted_edge(0, 1, 4)
        g.add_weighted_edge(0, 2, 2)
        g.add_weighted_edge(1, 3, 1)
        g.add_weighted_edge(2, 3, 5)
        g.add_weighted_edge(2, 4, 3)
        g.add_weighted_edge(3, 4, 1)
        distance, found = g.shortest_path_between(0, 4)
        self.assertTrue(found)
        self.assertEqual(distance, 5)


if __name__ == '__main__':
    unittest.main()
