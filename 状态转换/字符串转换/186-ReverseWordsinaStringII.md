# [#186 Reverse Words in a String II](https://leetcode.com/problems/reverse-words-in-a-string-ii/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、字符串、数组、逆置、字符处理、两次逆置

## 题目

Given an input string , reverse the string word by word.

## 简述

**输入：** 字符串的字符数组

**输出：** 将每个单词逆置后的字符串

**Notes：**

+ 单词为一个无空格字符序列
+ 输入字符串字符数组首尾无空格
+ 单词间的空格只有单个

## 思路

本题考察字符串转换，是[[引用]151题-ReverseWordsinaString](151-ReverseWordsinaString.md)的扩展题，输入参数从字符串变成了字符数组，条件也有细微差别，因此这道题考察的重点在于如何操作字符数组。

在151题中，共介绍了4种逆置方式，有直接调用内置方法的，也有字符处理+单词缓存的方式实现逆置操作的。本题可以继续沿用151题中的方法。即先将字符数组逆置，然后再将每个单词逆置。------方法1

## 解决方案

### 方法1-两次逆置法

先对整个数组进行逆置，再逆置每个单词。(关键词：字符处理、两次逆置)

时间复杂度：$O(n)$ ---44%

空间复杂度：$O(1)$ ---30%

``` java
class Solution {
    public void reverseWords(char[] s) {
        reverse(0, s.length - 1, s);
        int i = 0, j = -1;
        while (++j < s.length) {
            if (s[j] == ' '){
                reverse(i, j - 1, s);
                i = j + 1;
            }
        }
        reverse(i, j - 1, s);
    }

    public void reverse(int i, int j, char[] s) {
        for (i = i, j = j; i < j; ++i, --j) {
            char tmp = s[i];
            s[i] = s[j];
            s[j] = tmp;
        }
    }
}
```
