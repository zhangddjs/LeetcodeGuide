import unittest
import sys
import os
import importlib.util
from typing import Protocol


class UnionFindProtocol(Protocol):
    def find(self, x: int) -> int: ...
    def union(self, x: int, y: int) -> None: ...
    def connected(self, x: int, y: int) -> bool: ...
    def count(self) -> int: ...


def load_unionfind_module():
    """Dynamically load the most recent unionfind_*.py file"""
    parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    files = [f for f in os.listdir(parent_dir) if f.startswith('unionfind_') and f.endswith('.py')]

    if not files:
        raise FileNotFoundError("No unionfind_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(parent_dir, latest_file)

    spec = importlib.util.spec_from_file_location("unionfind_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module.UnionFind


class TestUnionFind(unittest.TestCase):

    def setUp(self):
        self.UnionFind = load_unionfind_module()

    def test_single_element(self):
        uf = self.UnionFind(1)
        self.assertEqual(uf.find(0), 0)
        self.assertEqual(uf.count(), 1)

    def test_two_separate_elements(self):
        uf = self.UnionFind(2)
        self.assertFalse(uf.connected(0, 1))
        self.assertEqual(uf.count(), 2)

    def test_basic_union(self):
        uf = self.UnionFind(3)
        uf.union(0, 1)
        self.assertTrue(uf.connected(0, 1))
        self.assertFalse(uf.connected(1, 2))
        self.assertEqual(uf.count(), 2)

    def test_chain_unions(self):
        uf = self.UnionFind(4)
        uf.union(0, 1)
        uf.union(1, 2)
        self.assertTrue(uf.connected(0, 2))
        self.assertFalse(uf.connected(0, 3))
        self.assertEqual(uf.count(), 2)

    def test_complex_operations(self):
        uf = self.UnionFind(6)
        uf.union(0, 1)
        uf.union(2, 3)
        uf.union(4, 5)
        self.assertEqual(uf.count(), 3)

        uf.union(1, 3)
        self.assertTrue(uf.connected(0, 2))
        self.assertFalse(uf.connected(0, 4))
        self.assertEqual(uf.count(), 2)

    def test_path_compression(self):
        uf = self.UnionFind(10)

        # Create a long chain: 0-1-2-3-4-5-6-7-8-9
        for i in range(9):
            uf.union(i, i + 1)

        # All should be connected
        for i in range(10):
            for j in range(i + 1, 10):
                self.assertTrue(uf.connected(i, j),
                              f"Expected {i} and {j} to be connected")

        # Should have only 1 component
        self.assertEqual(uf.count(), 1)

    def test_disjoint_sets(self):
        uf = self.UnionFind(8)

        # Create two disjoint sets: {0,1,2,3} and {4,5,6,7}
        uf.union(0, 1)
        uf.union(1, 2)
        uf.union(2, 3)

        uf.union(4, 5)
        uf.union(5, 6)
        uf.union(6, 7)

        # Test within-set connections
        test_connections = [
            (0, 1, True), (0, 2, True), (0, 3, True),
            (1, 2, True), (1, 3, True), (2, 3, True),
            (4, 5, True), (4, 6, True), (4, 7, True),
            (5, 6, True), (5, 7, True), (6, 7, True),
            (0, 4, False), (1, 5, False), (2, 6, False), (3, 7, False),
        ]

        for x, y, expected in test_connections:
            result = uf.connected(x, y)
            self.assertEqual(result, expected,
                           f"Connected({x}, {y}) = {result}, want {expected}")

        self.assertEqual(uf.count(), 2)

    def test_sequential_operations(self):
        uf = self.UnionFind(5)

        # Initial state - all separate
        self.assertEqual(uf.count(), 5, "Initial count should be 5")

        # Check not connected
        self.assertFalse(uf.connected(0, 1), "0 and 1 should not be connected initially")

        # Union first pair
        uf.union(0, 1)

        # Check now connected
        self.assertTrue(uf.connected(0, 1), "0 and 1 should be connected after union")

        # Count should decrease
        self.assertEqual(uf.count(), 4, "Count should be 4 after first union")

        # Check Find returns same root for connected elements
        root0 = uf.find(0)
        root1 = uf.find(1)
        self.assertEqual(root0, root1, "Find(0) and Find(1) should return same root")

        # Union another pair
        uf.union(2, 3)

        # Count should decrease again
        self.assertEqual(uf.count(), 3, "Count should be 3 after second union")

        # Elements 0,1 should not be connected to 2,3
        self.assertFalse(uf.connected(0, 2), "0 and 2 should not be connected")

    def test_idempotent(self):
        uf = self.UnionFind(3)

        # Union the same pair multiple times
        uf.union(0, 1)
        count1 = uf.count()

        uf.union(0, 1)  # Should have no effect
        count2 = uf.count()

        uf.union(1, 0)  # Reverse order, should have no effect
        count3 = uf.count()

        self.assertEqual(count1, count2, "Count should remain same after repeated union")
        self.assertEqual(count2, count3, "Count should remain same after reversed union")
        self.assertEqual(count1, 2, "Count should be 2")


if __name__ == '__main__':
    unittest.main()
