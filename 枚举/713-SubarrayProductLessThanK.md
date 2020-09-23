# [#713 Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

![Medium](/figures/Medium.svg)

## 关键词

统计、枚举、压缩、双指针、滑动窗口、哨兵、问题转化、前缀和、二分法

## 题目

Your are given an array of positive integers `nums`.

Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than `k`.

## 简述

**输入：** 整型数组; 整型k

**输出：** 累乘积小于k的(元素连续)子数组数

**Notes：**

+ 0 < 数组长度 <= 50000
+ 0 < 数组元素 <= 1000
+ 0 <= k < $10^6$

## 思路

本题考察统计和枚举，需要枚举所有满足条件的解并统计个数。

枚举题往往可以想到回溯法来遍历所有的可能性，如果出现大量重复计算，可以添加备忘录来减少计算量，并优化成自底向上的动态规划的方法。

这道题我们可以这样枚举：令子数组的长度从0增加到数组长度，记录每个长度每个子数组的乘积到备忘录中，这样一来增加一个元素时就可以直接乘备忘录中上一长度的值即可，从而避免重复运算。

不过根据子数组连续的特性，可以发现对于一些case是可以压缩的，比如对于某个子数组下标i到下标j，如果它们的累乘积小于k，则它们中的任意的连续子数组都满足此条件，因此我们使用双指针法(滑动窗口法)，当窗口累积小于k，则加上窗口大小到结果(也就是新增一个元素时新增的子数组组合数)，并右移右边界，否则右移左边界，窗口累积用哨兵记录。------方法1

## 解决方案

### 方法1-多指针法(滑动窗口法)

双指针代表子数组边界，哨兵记录子数组累积，在满足条件时右移右边界并统计，否则右移左边界。(关键词：压缩、双指针、滑动窗口、哨兵)

时间复杂度：$O(n)$ ---99%

空间复杂度：$O(1)$ ---60%

``` java
class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        if (k == 0) return 0;
        int res = 0, i = 0, j = 0, product = 1;
        for (j = i; j < nums.length; ++j) {
            product *= nums[j];
            while (product >= k && i <= j) product /= nums[i++];
            res += (j - i + 1);
        }
        return res;
    }
}
```

## 扩展

### 累乘累加转换

$\log(\prod_i x_i)=\sum_i \log(x_i)$

### 扩展方法-对数二分搜索[$^{[1]}$](#refer-anchor-1)

根据数学知识$\log(\prod_i x_i)=\sum_i \log(x_i)$，可以将累乘问题降级为累加问题，因为某些子数组的乘积可能太大。

转换后，问题也变得更加熟悉。可以通过前缀和数组来记录，`prefix[i+1] = nums[0] + nums[1] + ... + nums[i]`，那么要求的便是`prefix[j] - prefix[i] < k`.

前缀和数组是单调增的，所以可以引入二分搜索，搜索`[i,j]`，记录`[i,k]`宽度`(k <= j)`从而统计所有子数组数。(关键词：问题转化、前缀和、二分法)

时间复杂度：$O(n\log(n))$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        if (k == 0) return 0;
        double logk = Math.log(k);
        double[] prefix = new double[nums.length + 1];
        for (int i = 0; i < nums.length; i++) {
            prefix[i+1] = prefix[i] + Math.log(nums[i]);
        }

        int ans = 0;
        for (int i = 0; i < prefix.length; i++) {
            int lo = i + 1, hi = prefix.length;
            while (lo < hi) {
                int mi = lo + (hi - lo) / 2;
                if (prefix[mi] < prefix[i] + logk - 1e-9) lo = mi + 1;
                else hi = mi;
            }
            ans += lo - i - 1;
        }
        return ans;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 713-Solution](https://leetcode.com/problems/subarray-product-less-than-k/solution/)
