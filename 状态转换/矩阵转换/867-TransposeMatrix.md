# [#867 Transpose Matrix](https://leetcode.com/problems/transpose-matrix/)

![Easy](/figures/Easy.svg)

## 关键词

矩阵转换、矩阵转置、遍历(行-列)、插入(列-行)

## 题目

Given a matrix `A`, return the transpose of `A`.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

## 简述

**输入：** 二维矩阵数组

**输出：** 转置后的二维矩阵数组

**Notes：**

+ 1 <= 矩阵行数、列数 <= 1000

## 思路

本题考察矩阵转换，懂一些线性代数的知识会很有利。

常规方法是按行-列遍历，然后再按列-行插如到结果数组中。------方法1

## 解决方案

### 方法1-暴力法

初始化二维结果数组，按行-列遍历输入矩阵，按列-行插入结果数组中。(关键词：遍历(行-列)、插入(列-行))

时间复杂度：$O(n)$ ---46%

空间复杂度：$O(n)$ ---5%

``` java
class Solution {
    public int[][] transpose(int[][] A) {
        int [][] res = new int[A[0].length][A.length];
        for (int i = 0; i < A.length; ++i)
            for (int j = 0; j < A[0].length; ++j)
                res[j][i] = A[i][j];
        return res;
    }
}
```

## 扩展

### 扩展方法1-一行Python法[$^{[1]}$](#refer-anchor-1)

python调用特定函数可以一行实现

``` python
#
# copyright: LeetCode(https://leetcode.com)
# 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
#
return list(zip(*A))
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 867-Solution](https://leetcode.com/problems/transpose-matrix/discuss/146767/Python-1-Liner)
