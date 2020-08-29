# [#859 Buddy Strings](https://leetcode.com/problems/buddy-strings/)

![Easy](/figures/Easy.svg)

## 关键词

比大小、字符串、状态转换、求交集、字符处理、枚举false条件

## 题目

Given two strings `A` and `B` of lowercase letters, return `true` if and only if we can swap two letters in `A` so that the result equals `B`.

## 简述

**输入：** 两个字符串

**输出：** 特定转换后是否相等

**Notes：**

+ 0 <= 字符串长度 <= 20000
+ 字符串只包含小写字母

## 思路

本题考察比大小，要判断的是元素能否从一个状态特定转换到目标状态。

根据题目可以知道如下四种情况不可能转换成功，第一种是两个字符串长度不等，第二种是字符不相等的索引数大于2，第三种是两字符串中不同的字符数大于0，第四种是两字符串一开始就相等且没有重复字符。因此可以遍历字符串，一一比较当前索引的字符是否相等，并判断如果满足四种情况的任意一种便返回不等。------方法1

## 解决方案

### 方法1-暴力法(排除法)

遍历字符串，一一比较字符，枚举false条件，得到结果。(关键词：遍历、字符处理、枚举false条件)

时间复杂度：$O(n)$ ---25%

空间复杂度：$O(1)$ ---28%

``` java
class Solution {
    public boolean buddyStrings(String A, String B) {
        if(A.length() != B.length()) return false;  //case 1
        char difA = '\0', difB = '\0';
        Set<Character> set = new HashSet<>();
        int count = 0;
        boolean dup = false;
        for (int i = 0; i < A.length(); ++i) {
            if(set.remove(A.charAt(i))) dup = true;
            set.add(A.charAt(i));
            if (A.charAt(i) != B.charAt(i)) {
                count++;
                if (count > 2) return false;    //case 2
                else if (count == 1) {
                    difA = A.charAt(i);
                    difB = B.charAt(i);
                } else if (difA != B.charAt(i) || difB != A.charAt(i)) return false;     //case 3
            }
        }
        return count == 2 || (count == 0 && dup);   //case 4
    }
}
```
