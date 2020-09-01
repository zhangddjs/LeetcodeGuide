# [#868 Binary Gap](https://leetcode.com/problems/binary-gap/)

![Easy](/figures/Easy.svg)

## 关键词

求最大值(距离)、位操作、进制转换、遍历、排序、哨兵、注意边界

## 题目

Given a positive integer `N`, find and return the longest distance between two consecutive 1's in the binary representation of `N`.

If there aren't two consecutive 1's, return 0.

## 简述

**输入：** 十进制整数

**输出：** 对应二进制中1之间的最大距离值

**Notes：**

+ 1 <= 输入整数 <= 10$^9$

## 思路

本题考察求线性最大值和位操作。

既然要求二进制，需要知道十进制数怎么转二进制，同时得熟悉各种位操作方法。可以将该数的二进制数想象成一个元素位0/1的数组，只要遍历1之间的距离并用数组记录排序(索引法)或哨兵记录最大值即可得到最大值。用哨兵的时候需要注意两端情况。题目和[[引用]849题](849-MaximizeDistancetoClosestPerson.md)类似------方法1

索引法即在遍历时用缓存记录1所在的索引，然后按索引从小到大计算相邻索引的差值，求得最大。此处不具体实现。

## 解决方案

### 方法1-暴力法(多指针法)

遍历每个位，用哨兵记录最大值缓存(第二指针)和最大值。(关键词：遍历、双指针、哨兵)

时间复杂度：$O(\log{N})==O(n)==O(1)$ ---100% (N为输入整数，n为位数，位数<=32)

空间复杂度：$O(1)$ ---81%

``` java
class Solution {
    public int binaryGap(int N) {
        int res = 0, tmp = 1;
        boolean flag = false;
        while (N != 0) {
            if(N % 2 != 0 && !flag) flag = true;    //first '1' found
            else if (N % 2 != 0 && flag) {
                res = Math.max(tmp, res);
                tmp = 1;
            } else if(N % 2 == 0 && flag) tmp++;    //if cur is '0', second pointer + 1.
            N /= 2;
        }
        return res;
    }
}
```

方法中

## 扩展

### 扩展位处理方法[$^{[1]}$](#refer-anchor-1)

上述方法用`N % 2`来判断当前位数是否为1，下面介绍另一种位处理方法，通过位移的方法取位值，并简化了判断条件，缺点是复杂度固定为$O(32)$。

``` java
class Solution {
    public int binaryGap(int N) {
        int last = -1, ans = 0;
        for (int i = 0; i < 32; ++i)
            if (((N >> i) & 1) > 0) {
                if (last >= 0)
                    ans = Math.max(ans, i - last);
                last = i;
            }

        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 868-Solution](https://leetcode.com/problems/binary-gap/solution/)
