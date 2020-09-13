# [#5 Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

![Medium](/figures/Medium.svg)

## 关键词

求最值、枚举、查找、字符串、遍历、哨兵、三指针、动态规划、LCS、马拉车算法

## 题目

Given a string **s**, find the longest palindromic substring in **s**. You may assume that the maximum length of **s** is `1000`.

## 简述

**输入：** 字符串

**输出：** 最长回文串

## 思路

本题考察求最值和查找，对于求最值问题，一般都需要遍历所有符合的情况，然后用一个哨兵来记录最大值。

本题的关键也是难点在于如何在字符串中寻找回文串并枚举所有回文串，当这个难点解决后题目也就自然解决了。

首先我们从回文串的特性分析起，对于一个回文串，它必然是具有对称性的，也就是说存在一个对称轴，对称轴可能是该回文串最中间一个或两个相等字符。同时一个回文串可能包含多个子回文串，且这些子回文串的对称轴全部相同。

根据这些特性，我们可以发现，如果找到了对称轴，该轴对应的所有回文串便可以找到。

因此可以遍历字符串中的每个字符，以当前字符和当前两个字符为对称轴，通过双指针来找出所有回文串，用哨兵记录最长的回文串。------方法1

我们继续对回文串的性质进行分析，对于一个具有对称性质的回文串，如果我们将其进行逆置操作，其依然是一个回文串，且和逆置前完全相等。

有了这个性质后，如果我们把输入字符串进行逆置操作，那么其中所有的回文串都可以在原字符串中匹配到，而非回文串则不会在原字符串中匹配，那么此时，我们要找最长回文串，也就相当于找逆置后字符串中最长的无变化的串，换句话说，就是两字符串最大的公共子串。

不过需要注意类似于`aacdecaa`这种情形，按照如上的思路会很误认为`aac`是最长回文串，因此需要进行调整[$^{[1]}$](#refer-anchor-1)。也就是在得到最长公共子串时，进行逆置操作，并判断是否和原子串相等，如果相等，就说明该子串是回文串(判断回文串方法)。

寻找最大公共子串可以用基于LCS的算法来寻找。------方法2

对于最长回文子串，还有一个著名的算法-马拉车算法(Manacher's Algorithm)，可以在时间复杂度$O(n)$内完成查找。

## 解决方案

### 方法1-暴力法(多指针法)

遍历每个和每两个字符，以当前字符为基准枚举所有回文串，最终返回最长的。(关键词：遍历、枚举、三指针)

时间复杂度：$O(n^2)$ ---41%

空间复杂度：$O(1)$ ---11%

``` java
class Solution {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 2) return s;
        int maxLen = 1;
        String res = s.substring(0, 1);
        for (int i = 0; i < s.length() - 1; ++i) {
            String tmp = getLongestPalindrome(s, i, i);
            if (s.charAt(i) == s.charAt(i + 1)) {
                String tmp2 = getLongestPalindrome(s, i, i + 1);
                tmp = tmp.length() > tmp2.length() ? tmp : tmp2;
            }
            if (tmp.length() > maxLen) {
                maxLen = tmp.length();
                res = tmp;
            }
        }
        return res;
    }

    public String getLongestPalindrome(String s, int p1, int p2) {
        while (--p1 >= 0 && ++p2 < s.length() && s.charAt(p1) == s.charAt(p2));
        if (p1 == -1) return s.substring(p1 + 1, p2 + 1);
        else return s.substring(p1 + 1, p2);
    }
}
```

### 方法2-LCS

逆置字符串，返回最长公共子回文串。(关键词：逆置、LCS、动态规划)

时间复杂度：$O(n^2)$ ---16%

空间复杂度：$O(n)$ ---45%

``` java
class Solution {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 2) return s;
        String s2 = new StringBuilder(s).reverse().toString();
        int [] dp = new int[s.length()];
        int maxLen = 1, p = 0;
        //init
        for (int i = 0; i < s.length(); ++i) {
            dp[i] = s.charAt(0) == s2.charAt(i) ? 1 : 0;
        }
        //dp
        for (int i = 1; i < s.length(); ++i) {
            for (int j = s2.length() - 1; j > 0; --j) {
                dp[j] = s.charAt(i) == s2.charAt(j) ? dp[j - 1] + 1 : 0;
                if (dp[j] > maxLen) {
                    String tmp = s2.substring(j - dp[j] + 1, j + 1);
                    if (tmp.equals(new StringBuilder(tmp).reverse().toString())){
                        maxLen = dp[j];
                        p = j;
                    }
                }
            }
            dp[0] = s2.charAt(0) == s.charAt(i) ? 1 : 0;
        }
        //build res
        return s2.substring(p - maxLen + 1, p + 1);
    }
}
```

## 扩展

### 扩展方法1：更为暴力的暴力法[$^{[1]}$](#refer-anchor-1)

遍历字符串中所有子串，判断其是否是回文串，取出最长的回文串。

时间复杂度：$O(n^3)$

空间复杂度：$O(1)$

### 扩展方法2：动态规划[$^{[1]}$](#refer-anchor-1)

本方法为扩展方法1的优化法，当子串$S(i,j)$为回文串时，则代表$S(i + 1, j - 1)$且$S_i==S_j$，初始情况为$S(i,i)$,$S(i,i+1)=(S_i==S_j)$。

时间复杂度：$O(n^2)$

空间复杂度：$O(n^2)$

### 马拉车算法

单中心的四种case：

case1: 不选择当前最大回文串范围内的回文串完全属于当前最大回文串的中心点作为下一个中心。

case2: 当前最大回文串到达了字符串末尾，则不用继续。

case3: 下一个中心点选择为其回文串到达了当前回文串右边界且镜像回文串到达左边界。

case4: 不选择回文串到达右边界且镜像回文串超出左边界的中心点作为下一个中心。

对于偶数中心，则需要在每个字符间插入原串没有的字符，例如'$'符号，转化成单中心case。

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 5-Solution](https://leetcode.com/problems/longest-palindromic-substring/solution/)
