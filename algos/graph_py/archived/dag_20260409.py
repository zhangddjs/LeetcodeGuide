# Generated on 2026-04-09 10:18:34
# Daily practice file: dag_20260409.py (dag template)

from typing import List, Dict, Optional


class Dag:
    pass

    def add_edge(self, from_node: int, to_node: int) -> None:
        pass

    def add_weighted_edge(self, from_node: int, to_node: int, weight: int) -> None:
        pass

    def topological_sort(self) -> List[int]:
        """
        https://www.geeksforgeeks.org/dsa/topological-sort-using-dfs/
        """
        return []

    def topological_sort_bfs(self) -> List[int]:
        """
        https://www.geeksforgeeks.org/dsa/topological-sorting-indegree-based-solution/
        """
        return []

    def is_dag(self) -> bool:
        return False

    def shortest_path(self, start: int) -> Dict[int, int]:
        return {}

    def __init__(self):
        pass
