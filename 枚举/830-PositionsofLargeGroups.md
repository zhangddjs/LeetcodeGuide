# [#830 Positions of Large Groups](https://leetcode.com/problems/positions-of-large-groups/)

![Easy](/figures/Easy.svg)

## 关键词

字符串、枚举、遍历、双指针

## 题目

In a string `S` of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like `S = "abbxxxxzyy"` has the groups `"a"`, `"bb"`, `"xxxx"`, `"z"` and `"yy"`.

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.

## 描述

**输入：** 字符串

**输出：** 满足条件的所有子串在输入字符串中的起止点

**Notes：**

+ 1 <= 字符串长度 <= 1000.
+ 输入字符串中只有小写字母
+ 输出按字典序。

## 思路

本题考察枚举，需要按题目要求枚举出所有满足条件的情况。

本题不需要对输入字符串进行处理，直接遍历字符串的每个字符即可。可以使用双指针法遍历，最初指针1指向字符串开头，然后指针2从指针1的下一个位置向后遍历，如果指针2所指字符与指针1的不一致，则判断当前指针2和指针1的距离，如果大于等于3，则满足条件，将起止位置加入到结果集中。------方法1

## 解决方案

### 方法1-多指针法

双指针遍历字符串，将符合条件的结果加入到输出数组中。(关键词：遍历，双指针)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---61%

``` java
class Solution {
    public List<List<Integer>> largeGroupPositions(String S) {
        int i = 0, j = 1;
        List<List<Integer>> res = new ArrayList<>();
        while (j < S.length()) {
            if (S.charAt(j) != S.charAt(i)) {
                if(j - i >= 3) res.add(Arrays.asList(i, j - 1));
                i = j;
            }
            j++;
        }
        if(j - i >= 3) res.add(Arrays.asList(i, j - 1));
        return res;
    }
}
```
