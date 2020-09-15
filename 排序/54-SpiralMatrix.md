# [#54 Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)

![Medium](/figures/Medium.svg)

## 关键词

排序、矩阵、螺旋序列、由外向内遍历、螺旋遍历

## 题目

Given a matrix of _m x n_ elements (_m_ rows, _n_ columns), return all elements of the matrix in spiral order.

e.g.

``` text
1 2 3
4 5 6 ==> 1 2 3 6 9 8 7 4 5
7 8 9
```

## 简述

**输入：** 矩阵

**输出：** 矩阵的螺旋序列

## 思路

本题考察排序，需要将输入转化成指定序列进行输出。

很容易想到一种暴力解法，就是按照顺时针螺旋顺序对矩阵进行由外向内的遍历，将遍历到的元素按顺序添加到结果中，最后输出结果即可。------方法1

## 解决方案

### 方法1-暴力法

由外向内螺旋遍历，记录遍历过的元素。(关键词：仿真、由外向内遍历、螺旋遍历)

时间复杂度：$O(mn)$ ---100%

空间复杂度：$O(1)$ ---55% 不计结果集合

``` java
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList<>();
        if (matrix.length == 0 || matrix[0].length == 0) return res;
        int M = matrix.length, N = matrix[0].length;
        for (int i = 0; i <= (Math.min(M, N) + 1) / 2 - 1; ++i) {
            for (int j = i; j <= N - 1 - i; ++j) res.add(matrix[i][j]);
            for (int j = i + 1; j <= M - 1 - i; ++j) res.add(matrix[j][N - 1 - i]);
            for (int j = N - 2 - i; i != M - 1 - i && j > i; --j) res.add(matrix[M - 1 - i][j]);
            for (int j = M - 1 - i; i != N - 1 - i && j > i; --j) res.add(matrix[j][i]);
        }
        return res;
    }
}
```

## 扩展

### 各种1-line Python解法[$^{[1]}$](#refer-anchor-1)

详见[${[1]}$](#refer-anchor-1)

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 54-Discuss](https://leetcode.com/problems/spiral-matrix/discuss/20571/1-liner-in-Python-+-Ruby)
