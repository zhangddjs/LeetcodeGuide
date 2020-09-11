# [#125 Valid Palindrome](https://leetcode.com/problems/valid-palindrome/)

![Easy](/figures/Easy.svg)

## 关键词

判断、字符串、回文、遍历、双指针

## 题目

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

E.g.

"A man, a plan, a canal: Panama" -> true

"race a car" -> false

## 简述

**输入：** 字符串

**输出：** 是否回文

**Notes：**

+ 空字符串也是回文
+ 字符串只包含ASCII字符

## 思路

本题考察判断，需要分析列举返回true或false的条件，然后进行进一步处理。

回文字符串的特点是对称，要判断字符串是否是回文，可以想到用双指针的方法，遍历比较首尾字符，如果有字符不等，则不是回文，否则当指针都指向中间字符时，返回true。------方法1

## 解决方案

### 方法1-多指针法

双指针遍历字符串，两两比较字符串两端字符。(关键词：遍历、双指针)

时间复杂度：$O(n)$ ---73%

空间复杂度：$O(1)$ ---96%

``` java
class Solution {
    public boolean isPalindrome(String s) {
        int p1 = 0, p2 = s.length() - 1;
        while (p1 < p2) {
            while(p1 < p2 && !(Character.isLetter(s.charAt(p1)) || Character.isDigit(s.charAt(p1)))) p1++;  //Character.isLetterOrDigit
            while(p1 < p2 && !(Character.isLetter(s.charAt(p2)) || Character.isDigit(s.charAt(p2)))) p2--;
            if (Character.toLowerCase(s.charAt(p1)) == Character.toLowerCase(s.charAt(p2))) {
                p1++;
                p2--;
            } else return false;
        }
        return true;
    }
}
```
