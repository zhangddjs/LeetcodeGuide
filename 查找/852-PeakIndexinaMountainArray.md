# [#852 Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)

![Easy](/figures/Easy.svg)

## 关键词

数组、查找、求最值、遍历、哨兵、数组有序、二分法

## 题目

Let's call an array `A` a _mountain_ if the following properties hold:

+ `A.length >= 3`

+ There exists some `0 < i < A.length - 1` such that `A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]`
Given an array that is definitely a mountain, return any `i` such that `A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]`.

## 简述

**输入：** _山峰_ 数组

**输出：** 峰顶索引

**Notes：**

+ 3 <= 山峰数组长度 <= 10000
+ 0 <= 山峰高度 <= $10^6$

## 思路

本题考察求数组最大值索引，需要知道如何找到最大值。

首先可以对数组进行遍历，用哨兵实时记录当前最大值和最大值的索引，遍历完毕后返回哨兵记录的索引。------方法1

根据数组的特殊性质，遍历时不用全部遍历完也可以得到最大值的索引，即在遍历时，当前元素大于左右元素时，就是最大元素，直接返回其索引即可。------方法2

仔细观察可以发现输入数组是有序的，既然是有序的，就可以立刻想到二分法来解决此问题。------方法3

## 解决方案

### 方法1-暴力法

遍历，哨兵记录当前最大值和其索引，遍历完毕后返回哨兵记录的索引。(关键词：遍历、哨兵)

时间复杂度：$O(n)$ ---20%

空间复杂度：$O(1)$ ---85%

``` java
class Solution {
    public int peakIndexInMountainArray(int[] A) {
        int [] temp = new int[2];
        for (int i = 0; i < A.length; ++i) {
            if (A[i] > temp[0]) {
                temp[0] = A[i];
                temp[1] = i;
            }
        }
        return temp[1];
    }
}
```

### 方法2-优化的暴力法

遍历，当前元素大于左右时，返回其索引(关键词：遍历)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---57%

``` java
class Solution {
    public int peakIndexInMountainArray(int[] A) {
        for (int i = 1; i < A.length - 1; ++i)
            if (A[i] > A[i + 1]) return i;
        return -1;
    }
}
```

### 方法3-二分法

二分查找，当前元素大于左右时，返回其索引(关键词：数组有序、二分查找)

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(1)$ ---83%

``` java
class Solution {
    public int peakIndexInMountainArray(int[] A) {
        int low = 1, high = A.length - 1, mid;
        while (low <= high) {
            mid = low + (high - low) / 2;
            if (A[mid] > A[mid - 1] && A[mid] > A[mid + 1]) return mid;
            if (A[mid] < A[mid + 1]) low = mid + 1;
            else high = mid - 1;
        }
        return -1;
    }
}
```

## 扩展

### 扩展方法1-黄金分割搜素[$^{[1]}$](#refer-anchor-1)

黄金分割搜素是运用在单峰(unimodal)问题的一个高效算法，在性能上有如下两个特征[$^{[2]}$](#refer-anchor-2)：

+ 相对二分法可以减少50%的索引值的计算量，如果每个索引值需要用复杂函数计算，可以很大程度提高效率。

+ 黄金分割法有多个迭代，平均时间复杂度为$O(\log_{1.618}(n))$，为二分法的1.44倍。

如果索引值计算复杂，黄金分割法的效率提升很明显，时间为二分法的`0.5 * 1.44 = 0.72`倍。

(关键词：单峰问题、黄金分割查找)

``` python
def peakIndexInMountainArray(self, A):
    def gold1(l, r):
        return l + int(round((r - l) * 0.382))

    def gold2(l, r):
        return l + int(round((r - l) * 0.618))
    l, r = 0, len(A) - 1
    x1, x2 = gold1(l, r), gold2(l, r)
    while x1 < x2:
        if A[x1] < A[x2]:
            l = x1
            x1 = x2
            x2 = gold1(x1, r)
        else:
            r = x2
            x2 = x1
            x1 = gold2(l, x2)
    return A.index(max(A[l:r + 1]), l)
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 852-Discuss](https://leetcode.com/problems/peak-index-in-a-mountain-array/discuss/139848/C++JavaPython-Better-than-Binary-Search)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 852-Comment](https://leetcode.com/problems/peak-index-in-a-mountain-array/discuss/139848/C++JavaPython-Better-than-Binary-Search/245390)