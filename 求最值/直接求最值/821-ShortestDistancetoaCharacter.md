# [#821 Shortest Distance to a Character](https://leetcode.com/problems/shortest-distance-to-a-character/)

![Easy](/figures/Easy.svg)

## 关键词

字符串、遍历、排序、哨兵、三指针、两次遍历、缓存、批量求最值

## 题目

Given a string `S` and a character `C`, return an array of integers representing the shortest distance from the character `C` in the string.

## 描述

**输入：** 字符串；目标字符

**输出：** 字符串中每个字符和目标字符最短的距离

**Notes：**

+ 1 <= 字符串长度 <= 10000.
+ 目标字符保证出现在字符串中。
+ 字符串字符和目标字符均为小写。

## 思路

这道题考察如何寻找一个或多个目标元素和当前元素的最短距离。

我们知道，要得到多个目标元素和当前元素的最短距离，往往需要遍历所有的距离并排序或用哨兵取得最短的距离。对于这道题，要求一个最短距离的数组，我们可以把整个问题进行拆解，对每个元素，都可以用遍历目标元素与该元素的距离来求得最短距离，并记录到输出数组中。时间复杂度为$O(n^2\log(n))$或$O(n^2)$。------方法1

可以注意到，根据字符串的性质，对于每个字符，距离其最近的目标字符必然为其左边或右边第一个出现的目标字符，因此对于方法1，我们不需要遍历每个目标元素，而只需要从当前字符向左右出发，找到左边和右边第一个目标元素，求出最短距离。对于连续的字符来说，如果它们中间没有目标字符，那么它们左边和右边的第一个目标字符的位置都是相同的，因此使用两个指针来缓存目标字符位置，并适时更新。这种方法需要注意一些边界情况，就是左边或者右边没有目标字符的这两种情况。时间复杂度为$O(n)$。------方法2

双指针缓存法的核心其实也是遍历了字符串两次，我们还可以对空间再进行一个优化，通过两次遍历的方法来获取最小值，第一次遍历缓存当前字符和它左边第一个出现的目标字符的距离，第二次遍历比较当前字符和它右边第一个出现的目标字符的距离与当前记录的距离从而得出最短距离。时间复杂度为$O(n)$。------方法3[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-暴力解

遍历字符串，对于每个当前字符通过遍历得到每个目标字符和当前字符的距离，排序或用哨兵求出最短距离并记录。(关键词：遍历，排序，哨兵)

时间复杂度：$O(n^2\log(n))$或$O(n^2)$

空间复杂度：$O(n)$

### 方法2-多指针法

遍历字符串，用两个指针记录当前字符左右两边第一个目标字符，求出最短距离并记录。(关键词：遍历，三指针)

时间复杂度：$O(n)$   ---100%

空间复杂度：$O(n)$   ---92%

``` java
class Solution {
    public int[] shortestToChar(String S, char C) {
        int j = S.indexOf(C), i = -j;
        int [] res = new int[S.length()];
        for(int k = 0; k < S.length(); k++){
            res[k] = Math.min(j - k, k - i);
            if(k == j){
                i = j;
                j = S.indexOf(C, j + 1);
                j = j < 0 ? Integer.MAX_VALUE : j;
            }
        }
        return res;
    }
}
```

### 方法3-多次遍历法[$^{[1]}$](#refer-anchor-1)

两次遍历字符串，第一次记录当前字符与左边第一个目标字符的距离，第二次得到右边的距离并于当前记录的比较，更新最短距离。(关键词：两次遍历，缓存)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
class Solution {
    public int[] shortestToChar(String S, char C) {
        int N = S.length();
        int[] ans = new int[N];
        int prev = Integer.MIN_VALUE / 2;

        for (int i = 0; i < N; ++i) {
            if (S.charAt(i) == C) prev = i;
            ans[i] = i - prev;
        }

        prev = Integer.MAX_VALUE / 2;
        for (int i = N-1; i >= 0; --i) {
            if (S.charAt(i) == C) prev = i;
            ans[i] = Math.min(ans[i], prev - i);
        }

        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 821-Solution](https://leetcode.com/problems/shortest-distance-to-a-character/solution/)
