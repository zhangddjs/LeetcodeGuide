import unittest
import sys
import os
import importlib.util
from typing import Callable, List


def load_sortalgo_module():
    """Dynamically load the most recent sortalgo_*.py file"""
    current_dir = os.path.dirname(os.path.abspath(__file__))
    files = [f for f in os.listdir(current_dir) if f.startswith('sortalgo_') and f.endswith('.py') and not f.endswith('_test.py')]

    if not files:
        raise FileNotFoundError("No sortalgo_*.py file found")

    # Get the most recent file
    latest_file = sorted(files)[-1]
    module_path = os.path.join(current_dir, latest_file)

    spec = importlib.util.spec_from_file_location("sortalgo_impl", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)

    return module


def generate_sorted(n: int) -> List[int]:
    return [i + 1 for i in range(n)]


def generate_reverse(n: int) -> List[int]:
    return [n - i for i in range(n)]


def generate_duplicates(n: int) -> List[int]:
    return [(i % 10) + 1 for i in range(n)]


def generate_duplicates_sorted(n: int) -> List[int]:
    return sorted(generate_duplicates(n))


class TestSortAlgorithms(unittest.TestCase):

    def setUp(self):
        self.module = load_sortalgo_module()

    def run_sort_test(self, sort_func: Callable[[List[int]], None], test_name: str, input_arr: List[int], expected: List[int]):
        arr = input_arr.copy()
        sort_func(arr)
        self.assertEqual(arr, expected, f"{test_name} failed")

    def test_all_sorts(self):
        test_cases = [
            ("empty array", [], []),
            ("single element", [42], [42]),
            ("already sorted", [1, 2, 3, 4, 5], [1, 2, 3, 4, 5]),
            ("reverse sorted", [5, 4, 3, 2, 1], [1, 2, 3, 4, 5]),
            ("random order", [3, 1, 4, 1, 5, 9, 2, 6], [1, 1, 2, 3, 4, 5, 6, 9]),
            ("duplicates", [3, 3, 3, 1, 1, 2, 2], [1, 1, 2, 2, 3, 3, 3]),
            ("negative numbers", [-1, -5, 3, 0, -2], [-5, -2, -1, 0, 3]),
            ("large sorted", generate_sorted(1000), generate_sorted(1000)),
            ("large reverse", generate_reverse(1000), generate_sorted(1000)),
            ("large duplicates", generate_duplicates(500), generate_duplicates_sorted(500)),
        ]

        sort_algorithms = [
            ("heapsort", self.module.heapsort),
            ("quicksort", self.module.quicksort),
            ("mergesort", self.module.mergesort),
            ("insertsort", self.module.insertsort),
            ("bubblesort", self.module.bubblesort),
        ]

        for algo_name, sort_func in sort_algorithms:
            with self.subTest(algorithm=algo_name):
                for test_name, input_arr, expected in test_cases:
                    with self.subTest(test_case=test_name):
                        self.run_sort_test(sort_func, f"{algo_name} - {test_name}", input_arr, expected)


class TestHeapSort(unittest.TestCase):
    def setUp(self):
        self.module = load_sortalgo_module()

    def test_heapsort(self):
        arr = [3, 1, 4, 1, 5, 9, 2, 6]
        self.module.heapsort(arr)
        self.assertEqual(arr, [1, 1, 2, 3, 4, 5, 6, 9])


class TestQuickSort(unittest.TestCase):
    def setUp(self):
        self.module = load_sortalgo_module()

    def test_quicksort(self):
        arr = [3, 1, 4, 1, 5, 9, 2, 6]
        self.module.quicksort(arr)
        self.assertEqual(arr, [1, 1, 2, 3, 4, 5, 6, 9])


class TestMergeSort(unittest.TestCase):
    def setUp(self):
        self.module = load_sortalgo_module()

    def test_mergesort(self):
        arr = [3, 1, 4, 1, 5, 9, 2, 6]
        self.module.mergesort(arr)
        self.assertEqual(arr, [1, 1, 2, 3, 4, 5, 6, 9])


class TestInsertSort(unittest.TestCase):
    def setUp(self):
        self.module = load_sortalgo_module()

    def test_insertsort(self):
        arr = [3, 1, 4, 1, 5, 9, 2, 6]
        self.module.insertsort(arr)
        self.assertEqual(arr, [1, 1, 2, 3, 4, 5, 6, 9])


class TestBubbleSort(unittest.TestCase):
    def setUp(self):
        self.module = load_sortalgo_module()

    def test_bubblesort(self):
        arr = [3, 1, 4, 1, 5, 9, 2, 6]
        self.module.bubblesort(arr)
        self.assertEqual(arr, [1, 1, 2, 3, 4, 5, 6, 9])


def test_heapsort():
    """Quick test for heapsort in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestHeapSort
    suite = unittest.TestLoader().loadTestsFromTestCase(TestHeapSort)
    return unittest.TextTestRunner(verbosity=2).run(suite)
# test_heapsort()


def test_quicksort():
    """Quick test for quicksort in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestQuickSort
    suite = unittest.TestLoader().loadTestsFromTestCase(TestQuickSort)
    return unittest.TextTestRunner(verbosity=2).run(suite)
# test_quicksort()


def test_mergesort():
    """Quick test for mergesort in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestMergeSort
    suite = unittest.TestLoader().loadTestsFromTestCase(TestMergeSort)
    return unittest.TextTestRunner(verbosity=2).run(suite)
# test_mergesort()


def test_insertsort():
    """Quick test for insertsort in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestInsertSort
    suite = unittest.TestLoader().loadTestsFromTestCase(TestInsertSort)
    return unittest.TextTestRunner(verbosity=2).run(suite)
#test_insertsort()


def test_bubblesort():
    """Quick test for bubblesort in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestBubbleSort
    suite = unittest.TestLoader().loadTestsFromTestCase(TestBubbleSort)
    return unittest.TextTestRunner(verbosity=2).run(suite)
#test_bubblesort()


def test_all_sorts():
    """Quick test for all sorting algorithms in REPL"""
    import unittest
    from sortalgo_py.sortalgo_test import TestSortAlgorithms
    suite = unittest.TestLoader().loadTestsFromTestCase(TestSortAlgorithms)
    return unittest.TextTestRunner(verbosity=2).run(suite)


if __name__ == '__main__':
    unittest.main()
