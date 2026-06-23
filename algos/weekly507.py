import heapq
from typing import List

class SolutionI:
    def maxDistance(self, moves: str) -> int:
        dir = {
            'U': [1, 0],
            'D': [-1, 0],
            'L': [0, -1],
            'R': [0, 1],
            '_': [0, 0]
        }
        dist, row, col = 0, 0, 0
        for m in moves:
            row += dir[m][0]
            col += dir[m][1]
            dist += 1 if m == '_' else 0
        return abs(row)+abs(col)+dist

class SolutionII:
    def isvalid(self, num: int, x: int, div: int) -> tuple[ int, bool ]:
        last, first = -1, num // div
        last = num % 10
        if last != x:
            return div, False
        num //= div
        while num // 10 > 0:
            first = num // 10
            num //= 10
            div *= 10
        return div, first == x
    def countValidSubarrays(self, nums: list[int], x: int) -> int:
        cursum, n, cnt = 0, len(nums), 0
        for i in range(n):
            cursum, div = 0, 1
            for j in range(i, n):
                cursum += nums[j]
                div, valid = self.isvalid(cursum, x, div)
                cnt += 1 if valid else 0
        return cnt


class Solution:
    def buildgraph(self, n: int, edges: List[List[int]], labels: str) -> dict:
        graph = {i: {} for i in range(n)}
        for e in edges:
            u, v, w = e[0], e[1], e[2]
            if v not in graph[u] or w < graph[u][v]:
                graph[u][v] = w
        return graph

    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        graph = self.buildgraph(n, edges, labels)
        dist = {}
        dist[(0, 1)] = 0
        pq = [(0, 0, 1)]
        while pq:
            d, u, conu = heapq.heappop(pq)
            if (u, conu) in dist and d > dist[(u, conu)]:
                continue
            if u == n-1:
                return d
            for v, w in graph[u].items():
                if labels[v] == labels[u]:
                    conv = conu + 1
                    if conv > k:
                        continue
                else:
                    conv = 1
                if (v, conv) not in dist or d + w < dist[(v, conv)]:
                    dist[(v, conv)] = d + w
                    heapq.heappush(pq, (d+w, v, conv))
        return -1
