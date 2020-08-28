# [#840 Magic Squares In Grid]([leetcodelink](https://leetcode.com/problems/magic-squares-in-grid/))

![Easy](/figures/Easy.svg)

## 关键词

枚举、遍历(Top-Down)、优化判断条件、规律、组合数学

## 题目

A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers **from 1 to 9** such that each row, column, and both diagonals all have the same sum.

Given an `grid` of integers, how many 3 x 3 "magic square" subgrids are there?  (Each subgrid is contiguous).

## 简述

**输入：** 二维矩阵

**输出：** 三阶幻方数

**Notes：**

+ 1 \* 1 <= 矩阵大小 <= 10 \* 10
+ 0 <= 矩阵元素值 <= 15

## 思路

本题考察枚举，本题只需按题目要求返回符合条件的情况的数量即可，不需要列举。

拿到题目后，可以立马想到一种暴力解法，就是遍历所有情况，从矩阵左上角从左往右从上到下遍历，每有一个情况符合条件，就让结果加一，最后输出。这里需要注意的是怎样判断是否符合条件，我们可以将 3 x 3 矩阵的每一行、每一列、对角线的和都求出来，然后判断是否相等，并且要确保每个元素在1~9之间且无重复。------方法1

可以发现方法1中在判断是否满足条件时做了大量的重复计算，如果能使用缓存的话可以一定程度上减少计算量。同时还可以在判断时先判断 3 x 3 矩阵中间的元素是否等于5，如果不等于则不可能是幻方，从而进一步减少计算量。------方法2

## 解决方案

### 方法1-暴力法

遍历所有情况，判断每个情况是否满足条件，满足则结果+1。(关键词：遍历(Top-Down))

时间复杂度：$O(n)$ ---56%

空间复杂度：$O(1)$ ---50%

``` java
class Solution {
    public int numMagicSquaresInside(int[][] grid) {
        if(grid.length < 3 || grid[0].length < 3) return 0;
        int res = 0;
        for (int i = 0; i < grid.length - 2; i++)
            for (int j = 0; j < grid[i].length - 2; j++)
                res += isMagic(grid, i, j);
        return res;
    }
    public int isMagic(int[][] grid, int i, int j) {
        int sum = grid[i][j] + grid[i][j + 1] + grid[i][j + 2];
        Set<Integer> set = new HashSet<>();
        int [] diagtemps = new int[2];
        for (int k = 0; k < 3; k++) {
            int rowtemp = 0, coltemp = 0;
            for (int h = 0; h < 3; h++) {
                if (set.remove(grid[i + k][j + h]) ||
                    grid[i + k][j + h] > 9 ||
                    grid[i + k][j + h] < 1) return 0;
                set.add(grid[i + k][j + h]);
                rowtemp += grid[i + k][j + h];
                coltemp += grid[i + h][j + k];
                if(h == k) diagtemps[0] += grid[i + k][j + h];
                if(h + k == 2) diagtemps[1] += grid[i + k][j + h];
            }
            if (rowtemp != sum || coltemp != sum) return 0;
        }
        if(diagtemps[0] != sum || diagtemps[1] != sum) return 0;
        return 1;
    }
}
```

### 方法2-优化的暴力法

遍历所有情况，判断每个情况是否满足条件，满足则结果+1，使用缓存减少判断步骤计算量。(关键词：遍历(Top-Down)、缓存、优化判断条件)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

## 扩展

### 扩展方法1-优化判断算法[$^{[1]}$](#refer-anchor-1)

利用如下3阶幻方的一些特性可以优化判断算法。(关键词：优化判断条件、规律、组合数学)

+ 幻方的中心是5
+ 偶数在角落
+ 奇数在边上
+ 顺时针或逆时针按43816729序列排序

``` java
def numMagicSquaresInside(self, g):
    def isMagic(i, j):
        s = "".join(str(g[i + x / 3][j + x % 3]) for x in [0, 1, 2, 5, 8, 7, 6, 3])
        return g[i][j] % 2 == 0 and (s in "43816729" * 2 or s in "43816729"[::-1] * 2)
    return sum(isMagic(i, j) for i in range(len(g) - 2) for j in range(len(g[0]) - 2) if g[i + 1][j + 1] == 5)
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 840-Discuss](https://leetcode.com/problems/magic-squares-in-grid/discuss/133874/Python-5-and-43816729)
