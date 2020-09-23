# [#344 Reverse String](https://leetcode.com/problemset/all/?search=Reverse%20String)

![Easy](/figures/Easy.svg)

## 关键词

状态转换、字符串、逆置、数组、字符操作、遍历、双指针

## 题目

Write a function that reverses a string. The input string is given as an array of characters `char[]`.

Do not allocate extra space for another array, you must do this by **modifying the input array in-place** with O(1) extra memory.

You may assume all the characters consist of **printable ascii characters**.

## 简述

**输入：** 字符数组

**输出：** 就地逆置后的字符数组

## 思路

本题考察字符串转换，根据题目的要求对输入进行操作即可。

对于数组的逆置，比较容易想到的常用方法是双指针法，一个指针从前往后遍历，另一个指针从后往前遍历，在遍历时交换两个指针指向的字符元素，当指针相遇时算法结束，得到逆置结果。------方法1

## 解决方案

### 方法1-多指针法

双指针从两端开始遍历数组，交换指针指向元素，相遇时停止。(关键词：遍历、双指针)

时间复杂度：$O(n)$ ---70%

空间复杂度：$O(1)$ ---23%

``` java
class Solution {
    public void reverseString(char[] s) {
        char tmp = '0';
        for(int i = 0; i < s.length / 2; i++){
            tmp = s[i];
            s[i] = s[s.length - 1 - i];
            s[s.length - 1 - i] = tmp;
        }
    }
}
```

## 扩展

### 扩展方法1-Life is short, use Python. (c)[$^{[1]}$](#refer-anchor-1)

人生苦短，使用Python。

``` python
#
# copyright: LeetCode(https://leetcode.com)
# 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
#
class Solution:
    def reverseString(self, s):
        s.reverse()
```

### 扩展方法2-递归法[$^{[1]}$](#refer-anchor-1)

递归交换两端字符，但空间复杂度稍高一些。

``` java
class Solution {
  public void helper(char[] s, int left, int right) {
    if (left >= right) return;
    char tmp = s[left];
    s[left++] = s[right];
    s[right--] = tmp;
    helper(s, left, right);
  }

  public void reverseString(char[] s) {
    helper(s, 0, s.length - 1);
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 344-Solution](https://leetcode.com/problems/reverse-string/solution/)
