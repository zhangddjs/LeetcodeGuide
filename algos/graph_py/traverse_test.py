import unittest
import sys
import os
import importlib.util
from typing import Protocol, List


class GraphProtocol(Protocol):
    def add_edge(self, from_node: int, to_node: int) -> None: ...
    def dfs(self, start: int) -> List[int]: ...
    def dfs_whole(self) -> List[int]: ...
    def has_path(self, from_node: int, to_node: int) -> bool: ...
    def shortest_path(self, from_node: int, to_node: int) -> List[int]: ...


def load_traverse_module():
    """Dynamically load the most recent traverse_*.py file"""
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    files = [f for f in os.listdir(parent_dir) if f.startswith('traverse_') and f.endswith('.py')]

    if not files:
        raise FileNotFoundError("No traverse_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(parent_dir, latest_file)

    spec = importlib.util.spec_from_file_location("traverse_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module.Graph


class TestGraph(unittest.TestCase):

    def setUp(self):
        self.Graph = load_traverse_module()

    def test_add_edge_empty_graph(self):
        g = self.Graph()
        # Just verify it doesn't crash

    def test_add_edge_single_edge(self):
        g = self.Graph()
        g.add_edge(0, 1)

    def test_add_edge_multiple_edges(self):
        g = self.Graph()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])

    def test_dfs_empty_graph(self):
        g = self.Graph()
        result = g.dfs(0)
        self.assertEqual(result, [])

    def test_dfs_linear_path(self):
        g = self.Graph()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.dfs(0)
        self.assertEqual(result, [0, 1, 2, 3])

    def test_dfs_tree_structure(self):
        g = self.Graph()
        edges = [[0, 1], [0, 2], [1, 3], [1, 4]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.dfs(0)
        self.assertEqual(result, [0, 1, 3, 4, 2])

    def test_dfs_cycle_graph(self):
        g = self.Graph()
        edges = [[0, 1], [1, 2], [2, 3], [3, 0]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.dfs(0)
        self.assertEqual(result, [0, 1, 2, 3])

    def test_dfs_whole_empty_graph(self):
        g = self.Graph()
        result = g.dfs_whole()
        self.assertEqual(result, [])

    def test_dfs_whole_linear_path(self):
        g = self.Graph()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.dfs_whole()
        self.assertEqual(result, [0, 1, 2, 3])

    def test_dfs_whole_disconnected_components(self):
        g = self.Graph()
        edges = [[0, 1], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.dfs_whole()
        self.assertEqual(result, [0, 1, 2, 3])

    def test_has_path_empty_graph(self):
        g = self.Graph()
        self.assertFalse(g.has_path(0, 1))

    def test_has_path_same_node(self):
        g = self.Graph()
        g.add_edge(0, 1)
        self.assertTrue(g.has_path(0, 0))

    def test_has_path_direct_connection(self):
        g = self.Graph()
        g.add_edge(0, 1)
        self.assertTrue(g.has_path(0, 1))

    def test_has_path_no_connection(self):
        g = self.Graph()
        g.add_edge(0, 1)
        self.assertFalse(g.has_path(0, 2))

    def test_has_path_indirect_connection(self):
        g = self.Graph()
        edges = [[0, 1], [1, 2]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        self.assertTrue(g.has_path(0, 2))

    def test_has_path_disconnected_components(self):
        g = self.Graph()
        edges = [[0, 1], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        self.assertFalse(g.has_path(0, 2))

    def test_shortest_path_same_node(self):
        g = self.Graph()
        g.add_edge(0, 1)
        result = g.shortest_path(0, 0)
        self.assertEqual(result, [0])

    def test_shortest_path_direct_connection(self):
        g = self.Graph()
        g.add_edge(0, 1)
        result = g.shortest_path(0, 1)
        self.assertEqual(result, [0, 1])

    def test_shortest_path_no_connection(self):
        g = self.Graph()
        g.add_edge(0, 1)
        result = g.shortest_path(0, 2)
        self.assertEqual(result, [])

    def test_shortest_path_multiple_paths(self):
        g = self.Graph()
        edges = [[0, 1], [0, 2], [1, 3], [2, 3]]
        for edge in edges:
            g.add_edge(edge[0], edge[1])
        result = g.shortest_path(0, 3)
        # Both [0, 1, 3] and [0, 2, 3] are valid shortest paths
        self.assertIn(result, [[0, 1, 3], [0, 2, 3]])


if __name__ == '__main__':
    unittest.main()
