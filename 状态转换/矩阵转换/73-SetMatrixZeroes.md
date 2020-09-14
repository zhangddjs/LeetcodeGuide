# [#73 Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、矩阵、遍历、缓存、动态规划、投影

## 题目

Given an `m x n` matrix. If an element is **0**, set its entire row and column to **0**. Do it **in-place**.

**Follow up:**

+ A straight forward solution using $O(mn)$ space is probably a bad idea.
+ A simple improvement uses $O(m + n)$ space, but still not the best solution.
+ Could you devise a constant space solution?

## 简述

**输入：** 矩阵

**输出：** 就地转换后的矩阵

**Notes：**

+ 1 <= 矩阵长、宽 <= 200
+ -2$^{31}$ <= 矩阵元素 <= 2$^{31}$ - 1

## 思路

本题考察矩阵转换。

本题要求将矩阵中有0的元素所在的整行和整列置为0，很容易想到一种暴力法，就是遍历原始矩阵，每当遇到一个0，就遍历其行列置0，结果则存入新矩阵中。不过这样做时间复杂度为$O((mn)^2)$，空间复杂度为$O(mn)$，并且会出现大量重复运算。

我们可以对每个行列设置标记来避免重复运算，在遍历时遇到0时，先判断该0所在行或列有没有置0，如果没有则进行置0操作，并标记该行或列已经置0，置0结果则存入新矩阵中。这样做时间复杂度可以减少到$O(mn)$，但空间复杂度不变。

继续分析可以发现，对于一个`m x n`矩阵的每一行或每一列，如果有0，则表示该行/列需要被置0，否则不需要置0。因此我们可以扫描矩阵，将每个0所在的行和列用两个HashSet或数组记录，最后再根据记录的行和列对矩阵进行置0操作即可。------方法1

不过方法1的空间复杂度为$O(m + n)$，仍然不符合题目要求的就地。因此需要进一步优化。

如果不借助缓存的话，那么我们在遍历矩阵时就必须将置0操作执行，可如果按照常规的将整行或整列直接置0的话，会给后面遍历造成困扰，因为一行或一列上有多个0，遍历时无法判断该元素是被置0的还是原本就是0。
所以，我们是不可以操作未遍历的区域的。

那么我们从已遍历的区域来入手，对于当前元素来说，已遍历的区域为它的上面所有行以及左边所有列，如果当前元素为0，我们只需将当前行和当前列已遍历的每个元素置0即可，如果当前元素非0，则判断它左边和上边这两个相邻的元素是否存在0，如果存在，则说明当前元素处在0所在的行或列，需要被置0。

这样做看上去可以，但是会很快进入一个误区，就是在遍历下一行时，会被上一行已置0的元素影响，所以需要进行一个改进操作。

我们可以发现，遍历时是按行进行遍历的，如果这一行有一个0，那么一整行元素都可以置0，于是对于非0元素，我们可以不用判断左边是不是0，对于0元素，我们也不用急着去把它左边已遍历的元素置0。我们只需添加一个标记来记录当前行需不需要置0，同时为了不影响到下一层，我们在下一层遍历完后进行置0操作。------方法2

## 解决方案

### 方法1-改进的暴力法

遍历找到所有0，记录要置0的行和列，再根据记录将矩阵置0。(关键词：遍历、缓存)

时间复杂度：$O(mn)$ ---8%

空间复杂度：$O(m + n)$ ---18%

``` java
class Solution {
    public void setZeroes(int[][] matrix) {
        Set<Integer> set1 = new HashSet<>();
        Set<Integer> set2 = new HashSet<>();
        for (int i = 0; i < matrix.length; ++i)
            for (int j = 0; j < matrix[0].length; ++j)
                if (matrix[i][j] == 0) {
                    set1.add(i);
                    set2.add(j);
                }
        for (int row : set1)
            Arrays.fill(matrix[row], 0);
        for (int col : set2)
            for (int i = 0; i < matrix.length; ++i) matrix[i][col] = 0;
    }
}
```

### 方法2-动态规划法

按行遍历矩阵，用两个标记记录当前行和前一行是否需要置0，遇到0元素时，当前行需要置0，同时对当前列置0；非0元素取决于当前列上一行元素是否为0，为0则置0，当遍历完当前行后，将上一行根据标记进行处理。(关键词：动态规划)

时间复杂度：$O(mn)$ ---97%

空间复杂度：$O(1)$ ---87%

``` java
class Solution {
    public void setZeroes(int[][] matrix) {
        boolean pre = false, cur = false;
        for (int i = 0; i < matrix.length; ++i) {
            pre = cur;
            cur = false;
            for (int j = 0; j < matrix[0].length; ++j) {
                if (matrix[i][j] == 0) {
                    cur = true;
                    for (int k = i - 1; k >= 0; --k) matrix[k][j] = 0;
                } else if (i > 0 && matrix[i - 1][j] == 0) matrix[i][j] = 0;
            }
            if (pre) for (int k = 0; k < matrix[0].length; ++k) matrix[i - 1][k] = 0;
        }
        if (cur) for (int k = 0; k < matrix[0].length; ++k) matrix[matrix.length - 1][k] = 0;
    }
}
```

## 扩展

### 扩展方法-投影法[$^{[1]}$](#refer-anchor-1)

用矩阵第一行和第一列来记录该行或该列是否需要置0。每当遇到0元素时，将该元素所在行和列的第一个元素置为0，代表该行和该列需要置0。

``` java
class Solution {
  public void setZeroes(int[][] matrix) {
    Boolean isCol = false;
    int R = matrix.length;
    int C = matrix[0].length;

    for (int i = 0; i < R; i++) {

      // Since first cell for both first row and first column is the same i.e. matrix[0][0]
      // We can use an additional variable for either the first row/column.
      // For this solution we are using an additional variable for the first column
      // and using matrix[0][0] for the first row.
      if (matrix[i][0] == 0) {
        isCol = true;
      }

      for (int j = 1; j < C; j++) {
        // If an element is zero, we set the first element of the corresponding row and column to 0
        if (matrix[i][j] == 0) {
          matrix[0][j] = 0;
          matrix[i][0] = 0;
        }
      }
    }

    // Iterate over the array once again and using the first row and first column, update the elements.
    for (int i = 1; i < R; i++) {
      for (int j = 1; j < C; j++) {
        if (matrix[i][0] == 0 || matrix[0][j] == 0) {
          matrix[i][j] = 0;
        }
      }
    }

    // See if the first row needs to be set to zero as well
    if (matrix[0][0] == 0) {
      for (int j = 0; j < C; j++) {
        matrix[0][j] = 0;
      }
    }

    // See if the first column needs to be set to zero as well
    if (isCol) {
      for (int i = 0; i < R; i++) {
        matrix[i][0] = 0;
      }
    }
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 73-Solution](https://leetcode.com/problems/set-matrix-zeroes/solution/)
