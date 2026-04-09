import unittest
import sys
import os
import importlib.util
from typing import Protocol, List, Dict


class DagProtocol(Protocol):
    def add_edge(self, from_node: int, to_node: int) -> None: ...
    def add_weighted_edge(self, from_node: int, to_node: int, weight: int) -> None: ...
    def topological_sort(self) -> List[int]: ...
    def topological_sort_bfs(self) -> List[int]: ...
    def is_dag(self) -> bool: ...
    def shortest_path(self, start: int) -> Dict[int, int]: ...


def load_dag_module():
    """Dynamically load the most recent dag_*.py file"""
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    files = [f for f in os.listdir(parent_dir) if f.startswith('dag_') and f.endswith('.py')]

    if not files:
        raise FileNotFoundError("No dag_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(parent_dir, latest_file)

    spec = importlib.util.spec_from_file_location("dag_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module.Dag


class TestDag(unittest.TestCase):

    def setUp(self):
        self.Dag = load_dag_module()

    def test_add_edge_empty_dag(self):
        dag = self.Dag()
        # Just verify it doesn't crash

    def test_add_edge_single_edge(self):
        dag = self.Dag()
        dag.add_edge(0, 1)

    def test_add_edge_linear_chain(self):
        dag = self.Dag()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])

    def test_topological_sort_empty_dag(self):
        dag = self.Dag()
        result = dag.topological_sort()
        self.assertEqual(result, [])

    def test_topological_sort_single_edge(self):
        dag = self.Dag()
        dag.add_edge(0, 1)
        result = dag.topological_sort()
        self.assertEqual(result, [0, 1])

    def test_topological_sort_linear_chain(self):
        dag = self.Dag()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])
        result = dag.topological_sort()
        self.assertEqual(result, [0, 1, 2, 3])

    def test_topological_sort_diamond_pattern(self):
        dag = self.Dag()
        edges = [[0, 1], [0, 2], [1, 3], [2, 3]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])
        result = dag.topological_sort()
        # Both [0, 1, 2, 3] and [0, 2, 1, 3] are valid
        self.assertIn(result, [[0, 1, 2, 3], [0, 2, 1, 3]])

    def test_topological_sort_bfs_empty_dag(self):
        dag = self.Dag()
        result = dag.topological_sort_bfs()
        self.assertEqual(result, [])

    def test_topological_sort_bfs_single_edge(self):
        dag = self.Dag()
        dag.add_edge(0, 1)
        result = dag.topological_sort_bfs()
        self.assertEqual(result, [0, 1])

    def test_topological_sort_bfs_linear_chain(self):
        dag = self.Dag()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])
        result = dag.topological_sort_bfs()
        self.assertEqual(result, [0, 1, 2, 3])

    def test_is_dag_empty_graph(self):
        dag = self.Dag()
        self.assertTrue(dag.is_dag())

    def test_is_dag_single_edge(self):
        dag = self.Dag()
        dag.add_edge(0, 1)
        self.assertTrue(dag.is_dag())

    def test_is_dag_linear_chain(self):
        dag = self.Dag()
        edges = [[0, 1], [1, 2], [2, 3]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])
        self.assertTrue(dag.is_dag())

    def test_is_dag_simple_cycle(self):
        dag = self.Dag()
        edges = [[0, 1], [1, 2], [2, 0]]
        for edge in edges:
            dag.add_edge(edge[0], edge[1])
        self.assertFalse(dag.is_dag())

    def test_is_dag_self_loop(self):
        dag = self.Dag()
        dag.add_edge(0, 0)
        self.assertFalse(dag.is_dag())

    def test_add_weighted_edge_single_edge(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, 5)

    def test_add_weighted_edge_multiple_edges(self):
        dag = self.Dag()
        edges = [[0, 1, 4], [1, 2, 3], [0, 2, 8]]
        for edge in edges:
            dag.add_weighted_edge(edge[0], edge[1], edge[2])

    def test_shortest_path_empty_graph(self):
        dag = self.Dag()
        result = dag.shortest_path(0)
        self.assertEqual(result, {})

    def test_shortest_path_single_edge(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, 5)
        result = dag.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 5})

    def test_shortest_path_linear_path(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, 4)
        dag.add_weighted_edge(1, 2, 3)
        result = dag.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 4, 2: 7})

    def test_shortest_path_multiple_paths(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, 2)
        dag.add_weighted_edge(0, 2, 8)
        dag.add_weighted_edge(1, 2, 3)
        result = dag.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 2, 2: 5})

    def test_shortest_path_diamond_pattern(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, 2)
        dag.add_weighted_edge(0, 2, 6)
        dag.add_weighted_edge(1, 3, 1)
        dag.add_weighted_edge(2, 3, 1)
        result = dag.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: 2, 2: 6, 3: 3})

    def test_shortest_path_negative_weights(self):
        dag = self.Dag()
        dag.add_weighted_edge(0, 1, -2)
        dag.add_weighted_edge(0, 2, 4)
        dag.add_weighted_edge(1, 2, 3)
        dag.add_weighted_edge(1, 3, 2)
        dag.add_weighted_edge(2, 3, -5)
        result = dag.shortest_path(0)
        self.assertEqual(result, {0: 0, 1: -2, 2: 1, 3: -4})


if __name__ == '__main__':
    unittest.main()
