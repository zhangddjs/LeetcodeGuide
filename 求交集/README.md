# 求交集

## 哪些情形-解决方案

### 判断是否相交$^1$

+ 简单题$^1$

  + [836题-两个矩形是否重叠]-寻找相交条件、枚举true条件(相交)、几何数学

### 判断是否能从一种相交态(相离、部分相交、嵌套、完全相交)转换到另一种相交态$^1$

+ 简单题$^1$

  + [![[引用][比较]](/figures/Ref-Compare.svg) 859题-兄弟字符串](/比较/859-BuddyStrings.md)-遍历、枚举false条件

## 此类问题解决方案汇总

\*总题数$^2$

暴力法：遍历$^1$、寻找true条件$^1$(相交$^1$)、枚举true条件$^1$(相交$^1$)、true条件和false条件互推$^1$(不相交条件和相交条件互推$^1$)、枚举false条件$^2$(不相交$^1$)

逆转思维$^1$、几何数学$^1$、条件优化$^1$、条件判别法则不唯一$^1$

> for all intervals inputs,
Sweep Line method should be the first intuition you come up with.
> For each request [i,j],
we set count[i]++ and count[j + 1]--,
Then we sweep once the whole count,
we can find the frequency for count[i].

> t[start] +=1 means every index after this one will be counted 1 more time,

> t[end+1] -= 1, means every index this one will be counted 1 less time

## 常见最优时间复杂度

$O(1)^1$

<!-- 题目链接 -->
[836题-两个矩形是否重叠]:836-RectangleOverlap.md
