# [#42 Trapping Rain Water](https://leetcode.com/explore/interview/card/microsoft/30/array-and-strings/211/)

![Hard](/figures/Hard.svg)

## 关键词

统计、数组、求容量、拆分情况、多次遍历、动态规划、哨兵、栈、双指针

## 题目

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.

![eg](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. **Thanks Marcos** for contributing this image!

## 简述

**输入：** 代表海拔的数组

**输出：** 积水量

## 思路

本题考察统计。

根据短板和万有引力效应，在盆地中水面高度和海拔低的地面高度齐平，并且上坡或下坡环境中不会形成积水，因此只有当海拔出现`高(极大)-低(极小)-高(极大)`的情况时才可以积水，水面高度和较小的极大海拔齐平。

基于这个特性，我们可以将数组拆分成几段`高-低-高`情况的组合，这些片段可以再大致细分为如下几种情况：

1. 高-低-高
2. 高-低-次高-低-高
3. 次高-低-高
4. 高-低-次高
5. 高-低-次高-低-次低
6. 次低-低-次高-低-高

以上这些情况最终可以归类为如下2类：

1. 两端相等
2. 一端比另一端高

我们可以发现，对于第一类，从任意一端遍历都可以很快求出这一段的积水量。对于第二类，从低端向高端遍历，遇到低海拔时进行积水量统计(当前海拔与遍历过的最高海拔之差)，遇到高海拔的更新最高海拔为当前海拔。

分析到这里，题目也就可以迎刃而解了。我们可以先遍历一遍找到最高海拔，然后从两端向最高海拔进行遍历，从而统计出总积水量------方法1

## 解决方案

### 方法1-动态规划法(多次遍历法)

遍历一遍找到最高海拔，再从两端向最高海拔遍历，统计积水量。(关键词：拆分情况、多次遍历、哨兵)

时间复杂度：$O(n)$ ---94%

空间复杂度：$O(1)$ ---95%

``` java
class Solution {
    public int trap(int[] height) {
        if (height == null || height.length == 0) return 0;
        int max = height[0], maxIndex = 0;
        for (int i = 1; i < height.length; ++i) {
            if (max < height[i]) {
                max = height[i];
                maxIndex = i;
            }
        }
        int maximal = 0, res = 0;
        for (int i = 0; i < maxIndex; ++i) {
            if (height[i] < maximal) res += maximal - height[i];
            else maximal = height[i];
        }
        maximal = 0;
        for (int i = height.length - 1; i > maxIndex; --i) {
            if (height[i] < maximal) res += maximal - height[i];
            else maximal = height[i];
        }
        return res;
    }
}
```

## 扩展

### 扩展方法1-OnePass栈[$^{[1]}$](#refer-anchor-1)

动态规划法的OnePass改进法，遍历数组，如果当前元素小于等于栈顶元素，则入栈，否则，出栈直到栈顶元素比当前元素大，并记录储水量。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` c++
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
int trap(vector<int>& height)
{
    int ans = 0, current = 0;
    stack<int> st;
    while (current < height.size()) {
        while (!st.empty() && height[current] > height[st.top()]) {
            int top = st.top();
            st.pop();
            if (st.empty())
                break;
            int distance = current - st.top() - 1;
            int bounded_height = min(height[current], height[st.top()]) - height[top];
            ans += distance * bounded_height;
        }
        st.push(current++);
    }
    return ans;
}
```

### 扩展方法2-OnePass双指针法[$^{[1]}$](#refer-anchor-1)

动态规划法的OnePass改进法，令双指针从左右两端开始遍历，并分别记录遍历过的最大高度，最大高度更矮的指针优先遍历，并记录积水量，直到两指针相遇，也就是都指向全局最大高度。

时间复杂度：$O(n)$

空间复杂度：$O(1)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
int trap(vector<int>& height)
{
    int left = 0, right = height.size() - 1;
    int ans = 0;
    int left_max = 0, right_max = 0;
    while (left < right) {
        if (height[left] < height[right]) {
            height[left] >= left_max ? (left_max = height[left]) : ans += (left_max - height[left]);
            ++left;
        }
        else {
            height[right] >= right_max ? (right_max = height[right]) : ans += (right_max - height[right]);
            --right;
        }
    }
    return ans;
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 42-Solution](https://leetcode.com/problems/trapping-rain-water/solution/)
