# [#1 Two Sum](https://leetcode.com/problems/two-sum/)

![Easy](/figures/Easy.svg)

## 关键词

查找、遍历、枚举、HashMap

## 题目

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

## 简述

**输入：** 整型数组; 目标值

**输出：** 数组中相加为目标值的2个数的索引

**Notes：**

+ 2 <= 输入数组长度 <= 10$^5$
+ $-10^9$ <= 每个数、目标值 <= $10^9$
+ 数组中只存在一个符合的答案

## 思路

本题考察查找，如果输入有序，则可以尝试二分查找或滑动窗口法，由于本题输入数组无序，所以尝试顺序遍历的方法。

很容易想到一种暴力解法，遍历数组中所有可能的组合，然后将相加的值和目标对比，如果相同，直接返回即可。------方法1

枚举所有组合的时间复杂度是$O(n^2)$，我们可以发现暴力法中会有一定的重复计算，例如对于第一个数字，遍历所有和第一个数字的组合后如果没有发现匹配，则下次遇到和第一个数字相同的数字，将会重复计算一遍，因此可以添加备忘录来记录遍历过的数字以减少重复计算，用空间换时间------方法2

不管是暴力法还是优化的暴力法，都无法突破时间的瓶颈，甚至牺牲很大的空间来换取少量的时间。

可以得知，对于暴力法，无非就是对当前数字`i`的后面找一个数字`j`，使得`i + j == target`也就是`j == target - i`，如果能找到，那么可推导出对于`j`来说，它的前面则存在一个数字`i`，使得`i == target - j`。

因此我们可以得到一个更加简单的方法，就是从左到右顺序遍历数组，用哈希表存储遍历过的数字及其索引信息。对于当前的数字`i`，首先判断它前面是否存在数字`target - i`，存在则返回，否则将`i`和索引存入哈希表(或者对于当前`i`，判断它是否是前面数字的目标，是则返回，否则将`i`的目标`target - i`及`i`的索引存入哈希表)。------方法3

## 解决方案

### 方法1-暴力法

遍历枚举所有组合，返回符合条件的结果。(关键词：遍历、枚举)

时间复杂度：$O(n^2)$ ---29%

空间复杂度：$O(1)$ ---75%

``` java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length - 1; ++i)
            for (int j = i + 1; j < nums.length; ++j)
                if (nums[i] + nums[j] == target) return new int[]{i, j};
        return null;
    }
}
```

### 方法2-优化的暴力法

遍历枚举所有组合，跳过遍历过的数字，返回符合条件的结果。(关键词：遍历、枚举、HashSet)

时间复杂度：$O(n^2)$ ---37%

空间复杂度：$O(n)$ ---77%

``` java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < nums.length - 1; ++i) {
            if (set.contains(nums[i])) continue;
            set.add(nums[i]);
            for (int j = i + 1; j < nums.length; ++j)
                if (nums[i] + nums[j] == target) return new int[]{i, j};
        }
        return null;
    }
}
```

### 方法3-动态规划法

遍历数组，用hash表记录每个数字要找的目标，返回符合条件的结果。(关键词：遍历、HashMap)

时间复杂度：$O(n)$ ---77%

空间复杂度：$O(n)$ ---84%

``` java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; ++i) {
            if (map.keySet().contains(nums[i])) return new int[]{map.get(nums[i]), i};
            map.put(target - nums[i], i);
        }
        return null;
    }
}
```
