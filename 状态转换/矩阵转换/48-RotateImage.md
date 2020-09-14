# [#48 Rotate Image](https://leetcode.com/problems/rotate-image/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、矩阵、就地、由外向内遍历、链式替换、逆置、转置

## 题目

You are given an _n x n_ 2D `matrix` representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image **in-place**, which means you have to modify the input 2D matrix directly. **DO NOT** allocate another 2D matrix and do the rotation.

![eg](https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg)

## 简述

**输入：** 矩阵

**输出：** 顺时针旋转90°的矩阵

**Notes：**

+ 1 <= 矩阵长 == 矩阵宽 <= 20
+ -1000 <= 矩阵元素 <= 1000

## 思路

本题考察矩阵就地转换，往往只能修改遍历过的元素，不能影响到未遍历元素，同时可以用少量标记位辅助。

通过观察可以发现，矩阵旋转后的一个明显特征就是原先的第一行变成了最后一列，最后一列变成了最后一行，最后一行变成了第一列，第一列变成了第一行，从而构成了最外层一圈的旋转。我们将最外层这两行两列去掉后，内层也是相同的变化。

因此我们可以由外向内进行矩阵的旋转，由于每个元素的目标位置都已经明确，所以我们对当前一圈任意一条边的每个元素进行链式替换，用临时变量来存储当前被替换的元素。------方法1

## 解决方案

### 方法1-暴力法(链式替换)

由外向内进行旋转，对每一圈中某一条边的每个元素链式替换其他边上目标元素，用临时变量记录替换的值。(关键词：由外向内遍历、链式替换)

时间复杂度：$O(n)$ ---7%

空间复杂度：$O(1)$ ---99%

``` java
//顺时针链式替换
class Solution {
    public void rotate(int[][] matrix) {
        int N = matrix.length;
        for (int i = 0; i < N / 2; ++i) {
            for(int j = i; j < N - i - 1; ++j) {
                int next = matrix[j][N - 1 - i], cur = matrix[i][j];
                matrix[j][N - 1 - i] = cur;
                cur = next;
                next = matrix[N - 1 - i][N - 1 - j];
                matrix[N - 1 - i][N - 1 - j] = cur;
                cur = next;
                next = matrix[N - 1 - j][i];
                matrix[N - 1 - j][i] = cur;
                cur = next;
                next = matrix[i][j];
                matrix[i][j] = cur;
                cur = next;
            }
        }
    }
}

//逆时针链式替换
class Solution {
    public void rotate(int[][] matrix) {
        int N = matrix.length;
        for (int i = 0; i < N / 2; ++i) {
            for(int j = i; j < N - i - 1; ++j) {
                int pre = matrix[i][j];
                matrix[i][j] = matrix[N - 1 - j][i];
                matrix[N - 1 - j][i] = matrix[N - 1 - i][N - 1 - j];
                matrix[N - 1 - i][N - 1 - j] = matrix[j][N - 1 - i];
                matrix[j][N - 1 - i] = pre;
            }
        }
    }
}
```

## 扩展

### 扩展方法1-旋转4个矩形[$^{[1]}$](#refer-anchor-1)

方法1的变种，每一圈可以拆分成如图4个矩形。

![fore-rect](https://leetcode.com/problems/rotate-image/Figures/48/48_rectangles.png)

存储4个位置的新值并统一进行赋值操作。

``` java
class Solution {
  public void rotate(int[][] matrix) {
    int n = matrix.length;
    for (int i = 0; i < n / 2 + n % 2; i++) {
      for (int j = 0; j < n / 2; j++) {
        int[] tmp = new int[4];
        int row = i;
        int col = j;
        for (int k = 0; k < 4; k++) {
          tmp[k] = matrix[row][col];
          int x = row;
          row = col;
          col = n - 1 - x;
        }
        for (int k = 0; k < 4; k++) {
          matrix[row][col] = tmp[(k + 3) % 4];
          int x = row;
          row = col;
          col = n - 1 - x;
        }
      }
    }
  }
}
```

### 扩展方法2-通用方法(逆置与转置)[$^{[2]}$](#refer-anchor-2)

先将矩阵进行逆置，再进行转置。反过来也可以。(关键词：逆置、转置)

``` c++
/*
 * clockwise rotate
 * first reverse up to down, then swap the symmetry 
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
*/
void rotate(vector<vector<int> > &matrix) {
    reverse(matrix.begin(), matrix.end());
    for (int i = 0; i < matrix.size(); ++i) {
        for (int j = i + 1; j < matrix[i].size(); ++j)
            swap(matrix[i][j], matrix[j][i]);
    }
}

/*
 * anticlockwise rotate
 * first reverse left to right, then swap the symmetry
 * 1 2 3     3 2 1     3 6 9
 * 4 5 6  => 6 5 4  => 2 5 8
 * 7 8 9     9 8 7     1 4 7
*/
void anti_rotate(vector<vector<int> > &matrix) {
    for (auto vi : matrix) reverse(vi.begin(), vi.end());
    for (int i = 0; i < matrix.size(); ++i) {
        for (int j = i + 1; j < matrix[i].size(); ++j)
            swap(matrix[i][j], matrix[j][i]);
    }
}
```

### 扩展方法3-更多的简洁方法[$^{[3]}$](#refer-anchor-3)

有大神总结了很多Python写法，具体可见[${[3]}$](#refer-anchor-3)

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 48-Solution](https://leetcode.com/problems/rotate-image/solution/)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 48-Discuss](https://leetcode.com/problems/rotate-image/discuss/18872/A-common-method-to-rotate-the-image)

<div id="refer-anchor-3"></div>

+ [3] [Leetcode. 48-Discuss2](https://leetcode.com/problems/rotate-image/discuss/18884/Seven-Short-Solutions-(1-to-7-lines))
