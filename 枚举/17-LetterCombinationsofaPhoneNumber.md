# [#17 Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number)

![Medium](/figures/Medium.svg)

## 关键词

枚举、

## 题目

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order**.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png)

## 简述

**输入：** 手机键盘按键序列

**输出：** 所有可能的字母组合(无视顺序)

**Notes：**

+ 0 <= 序列长度 <= 4
+ 按键范围`['2', '9']`

## 思路

本题考察枚举，通常需要运用回溯法思想来解决此类问题。

对于本题，可以联想成图或者n叉树，然后使用深度遍历方法，对于序列中某个位置键位，将其对应的所有可能的字母和之前的遍历路径进行组合，然后对于每个组合继续遍历序列中下一个键位。------方法1

## 解决方案

### TODO方法1-什么方法

过程简述。(关键词：keyword、keyword)

时间复杂度：$O()$ ---%

空间复杂度：$O()$ ---%

``` java
//Solution
////Written after seeing @xxx's title "xxx" but before looking at its code.
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
