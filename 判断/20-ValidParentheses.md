# [#20 Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)

![Easy](/figures/Easy.svg)

## 关键词

判断、字符串、

## 题目

Given a string s containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

## 简述

**输入：** 只包括号的字符串

**输出：** 字符串是否满足条件

**Notes：**

+ 1 <= 字符串长度 <= 10$^4$
+ 字符串只包含字符`'()[]{}'`

## 思路

本题考察判断，应知道字符串合法和不合法的各种情形，再判断字符串属于哪种情形，从而推出合法不合法。

分析可知，字符串不合法的情形主要有3种(定义字符串中左括号在右括号的左边)：

1. 同类型括号的左括号数量小于右括号数量
2. 同类型括号的右括号数量小于左括号数量
3. 在某一对括号内出现如上两种情况

因此，我们可以从左到右扫描字符串，并借助于栈来比较左右括号是否能够匹配且顺序合法。在扫描字符串时，当遇到左括号时，进行入栈操作，右括号出现时，与栈顶元素进行匹配，如果匹配不上，则对应情况3，如果栈空，则对应情况1，否则出栈。如果遍历结束后栈非空，则对应情况2。------方法1

## 解决方案

### 方法1-暴力法

扫描字符串，分情况进行栈操作。(关键词：扫描字符串、拆分情况、栈)

时间复杂度：$O(n)$ ---44%

空间复杂度：$O(n)$ ---72%

``` java
class Solution {
    public boolean isValid(String s) {
        Set<Character> set = new HashSet<>(Arrays.asList('{', '[', '('));
        Map<Character, Character> map = new HashMap<>();
        Stack<Character> stack = new Stack<>();
        map.put('}', '{');
        map.put(']', '[');
        map.put(')', '(');
        for(int i = 0; i < s.length(); ++i) {
            if (set.contains(s.charAt(i))) stack.push(s.charAt(i));
            else if (stack.isEmpty() || stack.peek() != map.get(s.charAt(i))) return false;
            else stack.pop();
        }
        return stack.isEmpty();
    }
}
```
