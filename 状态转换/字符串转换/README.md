# 字符串转换

## 哪些情形-解决方案

### 按题目要求判断和操作每个单词或字符$^4$

+ 简单题$^3$

  + [824题-GoatLatin]-拆分字符串、拼接字符串、单词缓存、字符处理、字符串截取

  + [![[引用][求最值]](/figures/Ref-MaximumAndMinimum.svg) 819题-最频繁单词](/求最值/间接求最值/819-MostCommonWord.md)-正则、拆分字符串、单词缓存、字符处理

  + [![[引用][求最值]](/figures/Ref-MaximumAndMinimum.svg) 821题-最短字符距离](/求最值/直接求最值/821-ShortestDistancetoaCharacter.md)-字符查找

+ 中等题$^1$

  + [8题-字符串转整型](8-StringtoInteger(atoi).md)-扫描字符串、分情况处理、字符处理

### 判断某状态是否能转换到目标状态$^1$

+ 简单题$^1$

  + [![[引用][比较]](/figures/Ref-Compare.svg) 859题-兄弟字符串](/比较/859-BuddyStrings.md)-字符处理

### 某状态能转换到目标状态的最优转换方案$^1$

+ 中等题$^1$

  + [![[引用][求最值]](/figures/Ref-MaximumAndMinimum.svg) 712题-求最小的删除字母的ASCII和](/求最值/间接求最值/712-MinimumASCIIDeleteSumforTwoStrings.md)-多指针

### 表达式分析$^1$

+ 复杂题$^1$

  + [736题-转化Lisp表达式]-字符串分析、拆分字符串、分情况处理、栈、HashMap、递归

## 此类问题解决方案汇总

\*总题数$^7$

常规操作：扫描字符串$^1$、拆分字符串$^3$、拼接字符串$^1$、正则$^1$、字符串截取$^1$

字符处理$^4$、单词缓存$^2$、字符查找$^1$

多指针$^1$

字符串分析$^1$、分情况处理$^2$、栈$^1$、HashMap$^1$、递归$^1$

## 常见最优时间复杂度

$O(n)^1$、$O(n^2)^1$

## 注意

+ Java，Python中String是不可变的，而c++中string是可变的，根据不同语言的不同特性，往往有多种解法。本模块主要针对Java。

+ 某些字符串操作的封装方法的复杂度不包含在复杂度分析中。

## 扩展

### 扩展一些字符串正则处理方式[$^{[1]}$](#refer-anchor-1)

``` java
String normalizedStr = paragraph.replaceAll("[^a-zA-Z0-9 ]", " ").toLowerCase();
String[] words = normalizedStr.split("\\s+");

String[] words = p.replaceAll("\\W+" , " ").toLowerCase().split("\\s+");
```

<!-- 题目链接 -->
[824题-GoatLatin]:824-GoatLatin.md
[736题-转化Lisp表达式]:736-ParseLispExpression.md
[8题-字符串转整型]:8-StringtoInteger(atoi).md
