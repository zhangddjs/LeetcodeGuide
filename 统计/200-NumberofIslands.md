# [#200 Number of Islands](https://leetcode.com/problems/number-of-islands)

![Medium](/figures/Medium.svg)

## 关键词

统计、矩阵、岛屿问题、

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

空间复杂度：$O(1)$ ---99%

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

## 扩展

### TODO扩展一些知识和方法[$^{[1]}$](#refer-anchor-1)

内容

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
//Extension Solution
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 0-Solution]()
