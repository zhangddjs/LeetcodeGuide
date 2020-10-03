# [#200 Number of Islands](https://leetcode.com/problems/number-of-islands)

![Medium](/figures/Medium.svg)

## 关键词

统计、矩阵、岛屿问题、遍历、访问标记、DFS、并查集

## 题目

Given a 2d grid map of `'1'`s (land) and `'0'`s (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

## 简述

**输入：** 0-1矩阵

**输出：** 岛屿数量

## 思路

本题考察岛屿数量统计，为岛屿问题中最基本的一题。

既然要统计岛屿数量，首先得知道哪些情况属于岛屿，哪些不属于。对于任何一个岛屿来说，岛屿内部任意一个点都可以在不跨越水面的条件下通过一定路径到达另一个点，于是对于岛屿内所有的点，我们可以通过DFS的方式实现遍历访问，访问结束后岛屿总数+1。

据此，我们通过访问标记来表示某个点(或岛屿)是否访问过，然后从左上到右下遍历整个矩阵，遇到`1`则代表岛屿中的某个点，判断它是否访问过，如果访问过，则代表整个岛屿已经加入计数，不必再访问，否则使用DFS方法从该点出发遍历所有岛屿内的点。最后返回计数。------方法1

## 解决方案

### 方法1-DFS法

左上至右下遍历矩阵，遇到未统计的岛屿，用DFS访问所有点，计入统计。(关键词：遍历、访问标记、DFS)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---99%

``` java
class Solution {
    public int numIslands(char[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) return 0;
        int count = 0;
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                if (grid[i][j] == '1') {
                    count++;
                    helper(grid, i, j);
                }
        return count;
    }

    public void helper(char[][] grid, int i, int j) {
        if (i < 0 || i >= grid.length || j < 0 || j >= grid[0].length || grid[i][j] != '1') return;
        grid[i][j] = '2';
        helper(grid, i + 1, j);
        helper(grid, i - 1, j);
        helper(grid, i, j + 1);
        helper(grid, i, j - 1);
    }
}
```

> 深度优先遍历的空间复杂度较大，可以用BFS来优化空间复杂度。

## 扩展

### 扩展方法-并查集[$^{[1]}$](#refer-anchor-1)

初始化并查集数据结构，令每个陆地的parent指向自己，初始时岛屿数为矩阵中`1`的个数(陆地的个数)。遍历矩阵，对每一个陆地，连接上下左右四个方向上的相邻陆地，并标记已访问，令岛屿数减去相邻陆地数，最后返回的岛屿数即为要求的岛屿数。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
  class UnionFind {
    int count; // # of connected components
    int[] parent;
    int[] rank;

    public UnionFind(char[][] grid) { // for problem 200
      count = 0;
      int m = grid.length;
      int n = grid[0].length;
      parent = new int[m * n];
      rank = new int[m * n];
      for (int i = 0; i < m; ++i) {
        for (int j = 0; j < n; ++j) {
          if (grid[i][j] == '1') {
            parent[i * n + j] = i * n + j;
            ++count;
          }
          rank[i * n + j] = 0;
        }
      }
    }

    public int find(int i) { // path compression
      if (parent[i] != i) parent[i] = find(parent[i]);
      return parent[i];
    }

    public void union(int x, int y) { // union with rank
      int rootx = find(x);
      int rooty = find(y);
      if (rootx != rooty) {
        if (rank[rootx] > rank[rooty]) {
          parent[rooty] = rootx;
        } else if (rank[rootx] < rank[rooty]) {
          parent[rootx] = rooty;
        } else {
          parent[rooty] = rootx; rank[rootx] += 1;
        }
        --count;
      }
    }

    public int getCount() {
      return count;
    }
  }

  public int numIslands(char[][] grid) {
    if (grid == null || grid.length == 0) {
      return 0;
    }

    int nr = grid.length;
    int nc = grid[0].length;
    int num_islands = 0;
    UnionFind uf = new UnionFind(grid);
    for (int r = 0; r < nr; ++r) {
      for (int c = 0; c < nc; ++c) {
        if (grid[r][c] == '1') {
          grid[r][c] = '0';
          if (r - 1 >= 0 && grid[r-1][c] == '1') {
            uf.union(r * nc + c, (r-1) * nc + c);
          }
          if (r + 1 < nr && grid[r+1][c] == '1') {
            uf.union(r * nc + c, (r+1) * nc + c);
          }
          if (c - 1 >= 0 && grid[r][c-1] == '1') {
            uf.union(r * nc + c, r * nc + c - 1);
          }
          if (c + 1 < nc && grid[r][c+1] == '1') {
            uf.union(r * nc + c, r * nc + c + 1);
          }
        }
      }
    }

    return uf.getCount();
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 200-Solution](https://leetcode.com/problems/number-of-islands/solution/)
