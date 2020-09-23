# [#694 Number of Distinct Islands](https://leetcode.com/problems/number-of-distinct-islands/)

![Medium](/figures/Medium.svg)

## 关键词

分类、统计、矩阵、遍历、两两对比、去重、深度优先遍历、HashSet、坐标转索引、散列相对坐标、散列路径签名

## 题目

Given a non-empty 2D array `grid` of 0's and 1's, an **island** is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Count the number of **distinct** islands. An island is considered to be the same as another if and only if one island can be translated (and not rotated or reflected) to equal the other.

## 简述

**输入：** 包含水面和岛屿的二维数组

**输出：** 不同岛屿的数量

**Notes：**

+ 输入数组每个维度长不超过50
+ 0代表水面，1代表岛屿

## 思路

本题考察分类，看到分类题，往往可以联想到用HashMap或HashSet来存储分类信息或分类的统计信息。

本题只需要求类别数量，不需要求某个类下的元素数量，因此可以尝试用**HashSet**的方法来存储每种情况的岛屿，最后返回HashSet的长度即可。但是岛屿的情况复杂，用怎样的数据结构保存岛屿的信息是个问题。因此可以尝试一种**暴力方法**，也就是**先**遍历出所有岛屿，存放到一个集合中，记录总数，**然后**将每个岛屿和其它岛屿两两对比，如果当前岛屿和其它某岛屿相同，代表类别相同，总数减一，**最后**返回总数即可知道类别数。

**那么如何比较两座岛屿是否相等呢**？可以知道，遍历时是从左上到右下遍历，接触到某个岛屿时第一个必然接触到岛屿的最上最左的那个点，也就是**岛屿的入口**，随后采用**深度优先遍历**的方法遍历岛屿的每个位置。此时假设有两座岛屿完全相同，那么深度遍历它们每个位置的**顺序**是一致的，只要记录下遍历的每个位置的位置和顺序信息，就可以通过计算两座岛每个位置的**偏移量**来很方便地判断两座岛屿是否相同了(可以想象成两座相同岛屿的关系是岛屿A由岛屿B**平移**得到)。------方法1

由于一个岛屿平移后得到另一个相同岛屿的特点，令每个岛屿左上角**坐标**为(0, 0)，记录该岛屿每个点和左上角点的**相对坐标**，从而得到岛屿的形状。此处做一个很巧妙的转换，就是将**相对坐标转换成一维索引**。也就是一个整数，添加到HashSet中保存，于是每个岛屿的形状都可以转化成HashSet。最后再用一个HashSet保存所有形状对应的HashSet，这样一来重复的形状就被去掉了，返回其size即可得到结果。------方法2[$^{[1]}$](#refer-anchor-1)(散列相对坐标-Hash By Local Coordinates)

## 解决方案

### 方法1-暴力法

遍历数组，深度优先遍历提取所有岛屿，记录路径信息，两两对比每个岛屿，去掉重复岛屿。(关键词：遍历、两两对比、去重、深度优先遍历、记录遍历路径)

时间复杂度：$O(n)$ ---63%

空间复杂度：$O(n)$ ---79%

``` java
class Solution {
    public int numDistinctIslands(int[][] grid) {
        if (grid.length == 0 || grid[0].length == 0) return 0;
        List<List<Integer[]>> islands = new ArrayList<>();
        for (int i = 0; i < grid.length; ++i)
            for (int j = 0; j < grid[0].length; ++j)
                if (grid[i][j] == 1)
                    islands.add(dfs(grid, i, j, new ArrayList<Integer[]>()));
        int res = islands.size();
        for (int i = 0; i < islands.size(); ++i)
            for (int j = i + 1; j < islands.size(); ++j)
                if (judgeTwoIslands(grid, islands.get(i), islands.get(j))){
                    res--;
                    break;
                }
        return res;
    }

    public List dfs(int[][] grid, int i, int j, List<Integer[]> island) {
        if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length || grid[i][j] != 1) return island;
        grid[i][j] = 2;
        island.add(new Integer[]{i, j});
        dfs(grid, i - 1, j, island);
        dfs(grid, i, j - 1, island);
        dfs(grid, i, j + 1, island);
        return dfs(grid, i + 1, j, island);
    }

    public boolean judgeTwoIslands(int[][] grid, List<Integer[]> A, List<Integer[]> B) {
        if (A.size() != B.size()) return false;
        int [] offset = new int[]{B.get(0)[0] - A.get(0)[0], B.get(0)[1] - A.get(0)[1]};
        for (int i = 1; i < A.size(); ++i)
            if (B.get(i)[0] - A.get(i)[0] != offset[0] || 
                B.get(i)[1] - A.get(i)[1] != offset[1]) return false;
        return true;
    }
}
```

### 方法2-优化的暴力法(散列形状)[$^{[1]}$](#refer-anchor-1)

遍历数组，深度优先遍历提取所有岛屿，用HashSet记录岛屿形状，再用一个HashSet记录所有不同的形状。(关键词：遍历、深度优先遍历、HashSet、相对坐标转化整形)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    int[][] grid;
    boolean[][] seen;
    Set<Integer> shape;

    public void explore(int r, int c, int r0, int c0) {
        if (0 <= r && r < grid.length && 0 <= c && c < grid[0].length &&
                grid[r][c] == 1 && !seen[r][c]) {
            seen[r][c] = true;
            shape.add((r - r0) * 2 * grid[0].length + (c - c0));
            explore(r+1, c, r0, c0);
            explore(r-1, c, r0, c0);
            explore(r, c+1, r0, c0);
            explore(r, c-1, r0, c0);
        }
    }
    public int numDistinctIslands(int[][] grid) {
        this.grid = grid;
        seen = new boolean[grid.length][grid[0].length];
        Set shapes = new HashSet<HashSet<Integer>>();

        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[0].length; c++) {
                shape = new HashSet<Integer>();
                explore(r, c, r, c);
                if (!shape.isEmpty()) {
                    shapes.add(shape);
                }
            }
        }

        return shapes.size();
    }
}
```

## 扩展

### 扩展方法-用路径签名来优化方法1(散列路径签名-Hash By Path Signature)[$^{[1]}$](#refer-anchor-1)

对于岛屿中的一个点，从它开始或者进入它的遍历方向分为上下左右四个方向，分别记为1-4，那么每遍历到一个点，就把进入它的方向记录到集合中，然后携带方向签名继续遍历。其中初始点可以记为0。遍历结束后整个集合就是一个路径签名，将它添加到HashSet中，所有岛屿遍历结束后返回HashSet的size。(关键词：路径签名)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    int[][] grid;
    boolean[][] seen;
    ArrayList<Integer> shape;

    public void explore(int r, int c, int di) {
        if (0 <= r && r < grid.length && 0 <= c && c < grid[0].length &&
                grid[r][c] == 1 && !seen[r][c]) {
            seen[r][c] = true;
            shape.add(di);
            explore(r+1, c, 1);
            explore(r-1, c, 2);
            explore(r, c+1, 3);
            explore(r, c-1, 4);
            shape.add(0);
        }
    }
    public int numDistinctIslands(int[][] grid) {
        this.grid = grid;
        seen = new boolean[grid.length][grid[0].length];
        Set shapes = new HashSet<ArrayList<Integer>>();

        for (int r = 0; r < grid.length; r++) {
            for (int c = 0; c < grid[0].length; c++) {
                shape = new ArrayList<Integer>();
                explore(r, c, 0);
                if (!shape.isEmpty()) {
                    shapes.add(shape);
                }
            }
        }

        return shapes.size();
    }
}
```

### 相对坐标转索引

以长l高h的二维矩阵grid中某点(i, j)坐标为原点(0, 0)，对于某点(x, y)，其相对坐标为(x - i, y - j)，转化成对应整数为：

```java
int num = (x - i) * 2 * l + (y - j);    //2 * l 代替 l 因为(x - i) 可能为负数。
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 694-Solution](https://leetcode.com/problems/number-of-distinct-islands/solution/)
