# [#702 Search in a Sorted Array of Unknown Size](https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/)

![Medium](/figures/Medium.svg)

## 关键词

查找、数组、有序数组、大小未知、二分查找、滑动窗口优化

## 题目

Given an integer array sorted in ascending order, write a function to search `target` in nums. If target exists, then return its index, otherwise return `-1`. **However, the array size is unknown to you.** You may only access the array using an `ArrayReader` interface, where `ArrayReader.get(k)` returns the element of the array at index `k` (0-indexed).

You may assume all integers in the array are less than 10000, and if you access the array out of bounds, `ArrayReader.get` will return `2147483647`.

## 简述

**输入：** 大小未知的数组; 要查找的目标元素

**输出：** 目标元素索引

**Notes：**

+ 假设数组中每个元素唯一
+ 每个元素的值的范围是[-9999, 9999]
+ 数组大小的范围是[1, 10$^4$]

## 思路

本题考察二分查找，但与普通二分查找的不同是输入数组的大小未知，因此需要进行相关处理。

暴力法很容易解决问题，遍历一次数组即可得到目标元素，时间复杂度是$O(n)$。但这样做的话题目就没有意义了。我们知道对于一个未知大小的数组，假设它的大小为N，那么对于一个从0开始的足够大的范围来说，前N个是可以取到值的，对于这种特性，只要给定数组大小上限，就能得到当前数组的大小，根据题意数组大小上限已知，因此可以考虑用两次二分法，首先用二分法确定数组长度，然后再用一次二分法找到目标元素。------方法1

当然方法1是基于数组大小上限已知的情况，如果上限也是未知的，这种方法将失效。因此引入另一种方法-滑动+拥塞窗口法。该方法初始设定一个大小为1的窗口，判断右边界是否大于等于目标元素，如果小于，则窗口起始点移动到右边界位置，然后右边界扩大1倍，直到目标元素在窗口中某个位置时停止滑动，并在内部进行二分查找，从而找到最终目标。------方法2[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-多次二分法

两次二分查找，第一次确定数组大小，第二次确定目标元素。(关键词：二分查找)

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(1)$ ---76%

``` java
/**
 * // This is ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * interface ArrayReader {
 *     public int get(int index) {}
 * }
 */
class Solution {
    public int search(ArrayReader reader, int target) {
        int low = 0, high = 9999;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (reader.get(mid) == Integer.MAX_VALUE) high = mid - 1;
            else low = mid + 1;
        }
        low = 0;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (reader.get(mid) == target) return mid;
            if (reader.get(mid) < target) low = mid + 1;
            else high = mid - 1;
        }
        return -1;
    }
}
```

### 方法2-滑动窗口法[$^{[1]}$](#refer-anchor-1)

通过滑动窗口确定目标可能所在的范围，然后二分法找到索引。(关键词：二分查找、滑动窗口)

时间复杂度：$O(\log(T))$ --T是目标索引值

空间复杂度：$O(1)$

``` java
class Solution {
  public int search(ArrayReader reader, int target) {
    if (reader.get(0) == target) return 0;

    // search boundaries
    int left = 0, right = 1;
    while (reader.get(right) < target) {
      left = right;
      right <<= 1;
    }

    // binary search
    int pivot, num;
    while (left <= right) {
      pivot = left + ((right - left) >> 1);
      num = reader.get(pivot);

      if (num == target) return pivot;
      if (num > target) right = pivot - 1;
      else left = pivot + 1;
    }

    // there is no target element
    return -1;
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 702-Solution](https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/solution/)
