# [#695 Max Area of Island](https://leetcode.com/problems/max-area-of-island/)

![Medium](/figures/Medium.svg)

## 关键词

求最值、矩阵、遍历、深度优先遍历、统计、哨兵

## 题目

Given a non-empty 2D array `grid` of 0's and 1's, an **island** is a group of `1`'s (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

## 简述

**输入：** 包含水面和岛屿的二维数组

**输出：** 最大岛屿的面积

**Notes：**

+ 输入数组每个维度长不超过50
+ 0代表水面，1代表岛屿

## 思路

本题考察求最值。

解这道题只需要在深度优先遍历岛屿时增加统计和求最值的方法即可。我们知道，求最值常用的一个方法是哨兵记录最大值，因此这道题在统计岛屿大小后和哨兵比较并更新哨兵，最后返回哨兵值。------方法1

## 解决方案

### 方法1-暴力法

遍历数组，深度优先遍历每个岛屿，统计岛屿大小，哨兵记录最大值。(关键词：遍历、深度优先遍历、统计、哨兵)

时间复杂度：$O(n)$ ---68%

空间复杂度：$O(1)$ ---99%

``` java
class Solution {
    public int maxAreaOfIsland(int[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) return 0;
        int res = 0;
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                res = Math.max(dfs(grid, i, j), res);
        return res;
    }

    public int dfs(int [][] grid, int i, int j) {
        if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] != 1) return 0;
        int tmp = 1;
        grid[i][j] = 2;
        tmp += dfs(grid, i - 1, j);
        tmp += dfs(grid, i + 1, j);
        tmp += dfs(grid, i, j - 1);
        tmp += dfs(grid, i, j + 1);
        return tmp;
    }
}
```
