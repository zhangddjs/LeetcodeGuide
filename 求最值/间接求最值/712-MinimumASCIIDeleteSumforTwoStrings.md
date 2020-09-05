# [#712 Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/)

![Medium](/figures/Medium.svg)

## 关键词

求最值、枚举、统计、比较、求交集、状态转换、双指针、哨兵、回溯法、拆分子问题、拆分条件、递归、备忘录、动态规划法

## 题目

Given two strings `s1`, `s2`, find the lowest ASCII sum of deleted characters to make two strings equal.

## 简述

**输入：** 两个字符串

**输出：** 删除字符使得两字符串相等的最小删除ASCII总和

**Notes：**

+ 0 < 两字符串长度 < 1000
+ 所有字符小写，ASCII码范围是`[97, 122]`

## 思路

本题是一道综合题，考察了多个方面的知识，最终输出为最值，因此尝试使用求最值框架来解这道题。

很容易想到本题的暴力解法，就是枚举所有删除字符的可能组合，统计删除字符的ASCII总和，用哨兵记录最小的，当遍历完毕后返回哨兵。

难点在于如何枚举所有情况，可以知道枚举常用方法为回溯法，本题自然也可以尝试使用回溯法+双指针法来枚举所有情况。我们又知道，回溯法的核心是递归，递归需要base case和recursion case，那么接下来我们对几种需要回溯的case进行分析：

+ base case: 两个指针都指向字符串末尾->返回
+ recursion case1: 字符相等->都右移
+ recursion case2: 字符不等
  + 如果字符串A未到末尾，删除字符串A中的当前字符并统计
  + 如果字符串B未到末尾，删除字符串B中的当前字符并统计

至此回溯法的分析完毕。我们可以根据case方便地实现递归操作。------方法1

方法1虽然很简洁，但会触发TLE，所以需要进行一个优化操作。首先我们分析一下其中的重复计算，假设两字符串前两个字符都不相等，那么递归时会出现4条分支：

1. 删除字符串A的前两个字符
2. 删除字符串B的前两个字符
3. 先删字符串A的第一个字符，再删字符串B的第一个字符
4. 先删字符串B的第一个字符，再删字符串A的第一个字符

那么对于第三和第四种情况，它们后续的计算就都是重复的了。于是可以引入带备忘录的递归方式，通过二维数组，记录所有指针位置组合的最小值计算结果。------方法2

带备忘录的递归法得到后，我们可以方便地想到自底向上的动态规划法来对当前解法做更进一步的优化。------方法3

## 解决方案

### 方法1-暴力法(递归)

使用回溯法遍历所有情况，拆分递归条件和子问题，返回最大值。(关键词：回溯法、拆分子问题、递归、双指针、哨兵)

时间复杂度：$O(2^n)$ ---TLE

空间复杂度：$O(n)$ ---TLE

``` java
class Solution {
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty())) return 0;
        if (s1 == null || s1.isEmpty()) return (int)s2.charAt(0) + minimumDeleteSum(s1, s2.substring(1));
        else if (s2 == null || s2.isEmpty())return (int)s1.charAt(0) + minimumDeleteSum(s1.substring(1), s2);
        if (s1.charAt(0) == s2.charAt(0)) return minimumDeleteSum(s1.substring(1), s2.substring(1));
        return Math.min((int)s2.charAt(0) + minimumDeleteSum(s1, s2.substring(1)), (int)s1.charAt(0) + minimumDeleteSum(s1.substring(1), s2));
    }
}
```

### 方法2-优化的暴力法(带备忘录的递归)

使用回溯法遍历所有情况，拆分递归条件和子问题，用备忘录减少重复计算，返回最大值。(关键词：回溯法、拆分子问题、递归、备忘录)

时间复杂度：$O(n^2)$ ---43%

空间复杂度：$O(n^2)$ ---58%

``` java
class Solution {
    int [][] mem;
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty()) || s1.equals(s2)) return 0;
        mem = new int[s1.length() + 1][s2.length() + 1];
        for (int[] elm : mem) Arrays.fill(elm, -1);
        return dfs(s1, s2, 0, 0);
    }

    public int dfs(String s1, String s2, int p1, int p2) {
        if (p1 == s1.length() && p2 == s2.length()) return 0;
        if (p1 == s1.length())
            mem[p1][p2] = mem[p1][p2] == -1 ? 
                (int)s2.charAt(p2) + dfs(s1, s2, p1, p2 + 1) : mem[p1][p2];
        else if (p2 == s2.length())
            mem[p1][p2] = mem[p1][p2] == -1 ? 
                (int)s1.charAt(p1) + dfs(s1, s2, p1 + 1, p2) : mem[p1][p2];
        else if (s1.charAt(p1) == s2.charAt(p2))
            mem[p1][p2] = mem[p1][p2] == -1 ?
                dfs(s1, s2, p1 + 1, p2 + 1) : mem[p1][p2];
        if (mem[p1][p2] == -1)
            mem[p1][p2] = Math.min((int)s2.charAt(p2) + dfs(s1, s2, p1, p2 + 1),
                        (int)s1.charAt(p1) + dfs(s1, s2, p1 + 1, p2));
        return mem[p1][p2];
    }
}
```

### 方法3-动态规划法

将方法2进行改进，使用动态规划法自底向上遍历所有情况，返回最大值。(关键词：动态规划法)

时间复杂度：$O(n^2)$ ---88%

空间复杂度：$O(n^2)$ ---91%

``` java
class Solution {
    int [][] mem;
    public int minimumDeleteSum(String s1, String s2) {
        if ((s1 == null || s1.isEmpty()) && (s2 == null || s2.isEmpty()) || s1.equals(s2)) return 0;
        int M = s1.length(), N = s2.length();
        mem = new int[M + 1][N + 1];
        //init
        for (int[] elm : mem) Arrays.fill(elm, -1);
        mem[M][N] = 0;
        for (int i = M - 1; i >= 0; --i) mem[i][N] = mem[i + 1][N] + (int)s1.charAt(i);
        for (int j = N - 1; j >= 0; --j) mem[M][j] = mem[M][j + 1] + (int)s2.charAt(j);
        //iterate
        for (int i = M - 1; i >= 0; --i) {
            for (int j = N - 1; j >= 0; --j) {
                if (s1.charAt(i) == s2.charAt(j)) mem[i][j] = mem[i + 1][j + 1];
                else mem[i][j] = Math.min(s1.charAt(i) + mem[i + 1][j], s2.charAt(j) + mem[i][j + 1]);
            }
        }
        return mem[0][0];
    }
}
```
