# [#849 Maximize Distance to Closest Person](https://leetcode.com/problems/maximize-distance-to-closest-person/)

![Easy](/figures/Easy.svg)

## 关键词

数组、求线性最大值、遍历、两次遍历、双指针、三指针、排序、哨兵、规律、分组、压缩、局部推全局、注意边界

## 题目

In a row of `seats`, `1` represents a person sitting in that seat, and `0` represents that the seat is empty.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return that maximum distance to closest person.

## 简述

**输入：** 部分坐了人的座位数组

**输出：** 每个空座位和最近的人的距离中的最大值

**Notes：**

+ 2 <= 座位长度 <= 20000
+ 至少有一个空座位和一个已坐人的座位

## 思路

本题考察如何寻找一个或多个目标元素和当前元素的最短距离。

这道题和[[引用]821题](821-ShortestDistancetoaCharacter.md)非常相似，多了一步求最短距离集合中的最大值的步骤，因此求最短距离集合的方法沿用821题方法，在求出最小值集合后排序或者在遍历过程中使用哨兵获取最大值。------方法1、方法2

本题和821题不同的一点是本题只需要求得最短距离的最大值，不需要知道细节，虽然通过多指针法可以使空间最优，但时间上最好也要遍历两次，也就是$o(2n)$。这个时间复杂度也是可以优化的，根据观察发现，对于两个人中间的连续K个空座位，最大距离是`(K+1) / 2`，对于一个人和边界之间的连续K个空座位，最大距离是`K`。这两种情况中的最大值为本题的答案，最好的情况只需要$o(n)$。------方法3[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-多次遍历法

两次遍历数组，第一次从左往右记录每个空位置和左侧最近的人的距离数组，第二次从右往左记录右侧并比较取较小值，最后排序数组取得最大值。(关键词：两次遍历、双指针、缓存、排序)

时间复杂度：$O(n\log(n))$ ---87%

空间复杂度：$O(n)$ ---66%

``` java
class Solution {
    public int maxDistToClosest(int[] seats) {
        int [] shortest = new int[seats.length];
        int prev = -seats.length;
        for (int i = 0; i < seats.length; ++i) {
            if (seats[i] == 0) shortest[i] = i - prev;
            else prev = i;
        }
        prev = 2 * seats.length;
        for (int i = seats.length - 1; i >= 0; --i) {
            if (seats[i] == 0) shortest[i] = Math.min(shortest[i], prev - i);
            else prev = i;
        }
        Arrays.sort(shortest);
        return shortest[shortest.length - 1];
    }
}
```

### 方法2-多指针法

遍历数组，用两个指针记录两侧距离最近的人的位置，用哨兵记录最近距离的最大值。(关键词：遍历、三指针、哨兵)

时间复杂度：$O(n)$ ---87%

空间复杂度：$O(1)$ ---60%

``` java
class Solution {
    public int maxDistToClosest(int[] seats) {
        int i = -seats.length, j, res = 0;
        for (j = 0; j < seats.length && seats[j] != 1; ++j);
        for (int k = 0; k < seats.length; ++k) {
            if (k != j) res = Math.max(Math.min(j - k, k - i), res);
            else {
                i = j;
                for (j = j + 1; j < seats.length && seats[j] != 1; ++j);
                j = j == seats.length ? 2 * j : j;
            }
        }
        return res;
    }
}
```

### 方法3-分组法[$^{[1]}$](#refer-anchor-1)

遍历数组，人与人之间或人与边界之间的所有空座位具有相同的局部最大值，按空座位分组，从每组局部最大值可推出全局最大值。(关键词：遍历、规律、分组、压缩、局部推全局)

时间复杂度：$O(n)$

空间复杂度：$O(1)$

``` java
class Solution {
    public int maxDistToClosest(int[] seats) {
        int N = seats.length;
        int K = 0; //current longest group of empty seats
        int ans = 0;

        for (int i = 0; i < N; ++i) {
            if (seats[i] == 1) {
                K = 0;
            } else {
                K++;
                ans = Math.max(ans, (K + 1) / 2);
            }
        }

        for (int i = 0; i < N; ++i)  if (seats[i] == 1) {
            ans = Math.max(ans, i);
            break;
        }

        for (int i = N-1; i >= 0; --i)  if (seats[i] == 1) {
            ans = Math.max(ans, N - 1 - i);
            break;
        }

        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 849-Solution](https://leetcode.com/problems/maximize-distance-to-closest-person/solution/)
