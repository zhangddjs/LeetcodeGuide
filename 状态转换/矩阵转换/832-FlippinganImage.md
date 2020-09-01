# [#832 Positions of Large Groups](https://leetcode.com/problems/flipping-an-image/)

![Easy](/figures/Easy.svg)

## 关键词

矩阵转换、遍历、反转数组、逆置(异或)、特征处理

## 题目

Given a binary matrix `A`, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping `[1, 1, 0]` horizontally results in `[0, 1, 1]`.

To invert an image means that each `0` is replaced by `1`, and each `1` is replaced by `0`. For example, inverting `[0, 1, 1]` results in `[1, 0, 0]`.

## 简述

**输入：** 二维矩阵数组

**输出：** 水平翻转(reverse)并逆置(invert)后的二维矩阵数组

**Notes：**

+ 1 \* 1 <= 矩阵大小 <= 20 \* 20.
+ 矩阵中每个元素值为0或1
+ 矩阵长宽相等

## 思路

本题考察矩阵转换，懂一些线性代数的知识会很有利。

常规方法是按行-列遍历矩阵，对每一行进行反转操作，在反转的同时对对调后的元素进行取反，就可以在一次遍历下完成所有要求的操作。------方法1

根据线性代数知识，我们还可以通过构造反转和逆置矩阵与初始矩阵做矩阵乘法，不过作者暂时并没有找到线性代数中反转矩阵的实现，外加实现起来会消耗大量的额外空间和计算量，所以暂时不实现。

## 解决方案

### 方法1-暴力法

按行遍历矩阵，反转每一行的元素，并对所有元素取反。(关键词：遍历(行-列)、反转数组、逆置)

时间复杂度：$O(n)$ ---51% $n$为所有元素个数

空间复杂度：$O(1)$ ---56%

``` java
class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        for (int i = 0; i < A.length; ++i) {
            for (int j = 0; j < A[i].length / 2; ++j){
                A[i][j] = A[i][j] + A[i][A[i].length - j - 1];
                A[i][A[i].length - j - 1] = A[i][j] - A[i][A[i].length - j - 1];
                A[i][j] -= A[i][A[i].length - j - 1];
                A[i][j] = 1 - A[i][j];
                A[i][A[i].length - j - 1] = 1 - A[i][A[i].length - j - 1];
            }
            if (A[i].length % 2 == 1) A[i][A[i].length / 2] = 1 - A[i][A[i].length / 2];
        }
        return A;
    }
}
```

## 扩展

### 扩展方法1-更简洁的暴力法

该官方方法[$^{[1]}$](#refer-anchor-1)在遍历行元素时做了长度+1处理，不用在循环结束后再单独判断中间元素，同时使用异或的方法进行了逆置，更加高效。(关键词：逆置(异或))

``` java
class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        int C = A[0].length;
        for (int[] row: A)
            for (int i = 0; i < (C + 1) / 2; ++i) {
                int tmp = row[i] ^ 1;
                row[i] = row[C - 1 - i] ^ 1;
                row[C - 1 - i] = tmp;
            }
        return A;
    }
}
```

该最高赞方法[$^{[2]}$](#refer-anchor-2)也在边界上进行了处理，用`i * 2 < n`作为判断条件，从而也不会错过中间元素。该方法还简化了反转和逆置操作，根据性质添加了判断条件减少了计算量，没有借助额外空间，两行代码完成，非常巧妙。(关键词：特征处理)

``` java
class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        int n = A.length;
        for (int[] row : A)
            for (int i = 0; i * 2 < n; i++)
                if (row[i] == row[n - i - 1])
                    row[i] = row[n - i - 1] ^= 1;
        return A;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 832-Solution](https://leetcode.com/problems/flipping-an-image/solution/)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 832-Discuss1](https://leetcode.com/problems/flipping-an-image/discuss/130590/C++JavaPython-Reverse-and-Toggle)

