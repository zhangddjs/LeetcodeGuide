# [#8 String to Integer (atoi)](https://leetcode.com/problemset/all/?search=String%20to%20Integer%20(atoi))

![Medium](/figures/Medium.svg)

## 关键词

字符串转换、扫描字符串、分情况处理、字符处理

## 题目

Implement `atoi` which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

## 简述

**输入：** 字符串

**输出：** 转换成的整型数字

**Notes：**

+ 如果字符串对应整型超过32位时，直接转成整型最大值或最小值($2^{31} - 1$ or $-2^{31}$)

## 思路

本题考察字符串转换。

首先可以用trim方法去掉开头和结尾的空格字符，然后遍历字符串，并根据字符串和整型的性质分情况进行分析与处理。

1. 以`'+'`、`'-'`符号开头，代表整型符号位 ==> 用整型`1`、`-1`记录，最初为`1`。
2. 以非数字开头或者字符串为空 ==> 返回`0`。
3. 去掉符号位(如果有)后以数字开头 ==> 记录到整型结果缓存`buf`，最初为`0`。
4. 当前字符是数字，且`buf`大小在整型范围内 ==> 更新`buf`
5. 当前字符不是数字 ==> 返回`buf * 符号位`的结果。
6. 当前`buf`超出整型范围 ==> 返回INT_MAX或INT_MIN。

根据如上情况处理字符串后便可得到目标结果。------方法1

## 解决方案

### 方法1-暴力方法

遍历字符串，按条件进行字符处理。(关键词：遍历、拆分条件、字符处理)

时间复杂度：$O(n)$ ---83%

空间复杂度：$O(1)$ ---84%

``` java
class Solution {
    public int myAtoi(String str) {
        int i = 0;
        long buf = 0, sign = 1;
        while (i < str.length() && str.charAt(i) == ' ') i++;
        if(i == str.length()) return 0;
        else if(str.charAt(i) == '+' || str.charAt(i) == '-') sign = str.charAt(i++) == '+' ? 1 : -1;
        else if(!Character.isDigit(str.charAt(i))) return 0;
        while (i != str.length() && Character.isDigit(str.charAt(i)) && 
               sign * buf < Integer.MAX_VALUE && sign * buf > Integer.MIN_VALUE)
            buf = buf * 10 + (str.charAt(i++) - '0');
        if (sign * buf < Integer.MAX_VALUE && sign * buf > Integer.MIN_VALUE)
            return new Long(buf * sign).intValue();
        return sign * buf >= Integer.MAX_VALUE ? Integer.MAX_VALUE : Integer.MIN_VALUE;
    }
}
```
