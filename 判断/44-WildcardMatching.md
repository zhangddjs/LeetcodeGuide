# [#44 Wildcard Matching](https://leetcode.com/problems/wildcard-matching)

![Hard](/figures/Hard.svg)

## 关键词

判断、字符串、模式匹配、

## 题目

Given an input string (`s`) and a pattern (`p`), implement wildcard pattern matching with support for `'?'` and `'*'` where:

+ `'?'` Matches any single character.
+ `'*'` Matches any sequence of characters (including the empty sequence).

The matching should cover the **entire** input string (not partial).

e.g.

``` text
      Input:  s = "adcebfg", p = "*a*b?"
     Output:  false
Explanation:  The first '*' matches the empty sequence, while the second '*' matches the substring "dce". The '?' couldn't match "fg".
```

## 简述

**输入：** 字符串; 模式串

**输出：** 是否匹配

**Notes：**

+ 0 <= 字符串长度，模式串长度 <= 2000
+ 字符串只包含小写英文字母
+ 模式串包含小写英文字母和`'?'`、`'*'`

## 思路

本题考察字符串的模式匹配，可以先确定什么情况下匹配，什么情况下不匹配，然后再分析输入串并做出判断。

我们可以先从模式串入手进行分析，模式串主要有以下4种情况：

+ Case 1: 不包含`'?'`、`'*'`
+ Case 2: 只包含`'?'`
+ Case 3: 只包含`'*'`
+ Case 4: 同时包含`'?'`、`'*'`

对于情况1，一一对比字符串和模式串每个字符即可。

对于情况2，一一对比字符串和模式串每个字符，遇到`'?'`时可以匹配任意字符。

对于情况3和4，一一对比字符串和模式串每个字符，遇到`'*'`时，可以匹配任意长度的子串，这时我们可以借助回溯法，即`'*'`匹配0个字符、`'*'`匹配1个字符、`'*'`匹配2个字符...

总结下来，我们可以得到一个暴力法，即一一对比两串的每个字符，模式串遇到`'?'`可以匹配字符串任意字符，遇到`'*'`则可以用回溯的方法探测每一种情况。------方法1

但显然暴力回溯法时间复杂度太高，而且往往会做大量的重复运算，因此我们需要将算法进行优化，寻找出重复计算的地方，提前终止回溯。



## 解决方案

### 方法1-暴力回溯法

一一对比字符，遇到`'?'`匹配任意字符，遇到`'*'`则使用回溯法。(关键词：对比、回溯)

时间复杂度：$O()$ ---TLE

空间复杂度：$O()$ ---TLE

``` java
class Solution {
    public boolean isMatch(String s, String p) {
        if (s.length() == 0) {
            for(char c : p.toCharArray()) if (c != '*') return false;
            return true;
        }

        if (p.length() == 0) return false;

        boolean flag = false;

        if (p.charAt(0) == s.charAt(0) || p.charAt(0) == '?') {
            flag = isMatch(s.substring(1), p.substring(1));
        } else if (p.charAt(0) == '*') {
            for (int i = 0; i <= s.length() && !flag; ++i) {
                flag |= isMatch(s.substring(i), p.substring(1));
            }
        }
        return flag;
    }
}
```

## 扩展

### TODO扩展一些知识和方法[$^{[1]}$](#refer-anchor-1)

内容

``` java
/**
 * Copyright: LeetCode(https://leetcode.com)
 * Author: covfefe
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
//Extension Solution
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 0-Solution]()
