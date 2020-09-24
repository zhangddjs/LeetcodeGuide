# 判断

## 哪些情形-解决方案

### 判断是否能提供资源$^2$

+ 简单题$^1$

  + [![[引用][资源分配]](/figures/Ref-ResourceAllocation.svg) 860题-能否找零](/资源分配/860-LemonadeChange.md)-遍历、枚举false条件、拆分情况

+ 中等题$^1$

  + [![[引用][资源分配]](/figures/Ref-ResourceAllocation.svg) 698题-(n/k)Sum](/资源分配/698-PartitiontoKEqualSumSubsets.md)-枚举false条件

### 判断是否相交$^2$

+ 简单题$^2$

  + [141题-有环链表]-快慢指针、HashSet

  + [![[引用][求交集]](/figures/Ref-Intersection.svg) 836题-两个矩形是否重叠](/求交集/836-RectangleOverlap.md)-拆分情况、枚举true条件(相交)、枚举false条件、true和false条件转化、情况转化

### 判断相等(两元素能否在特定转换后相等)$^1$

+ 简单题$^1$

  + [![[引用][比较]](/figures/Ref-Compare.svg) 859题-兄弟字符串](/比较/859-BuddyStrings.md)-遍历、拆分情况、枚举false条件

### 判断回文、对称、合法$^3$

+ 简单题$^2$

  + [125题-合法回文]-遍历、双指针

  + [20题-合法括号]-遍历、拆分情况、栈

+ 中等题$^1$

  + [98题-合法BST]-递归、拆分情况、情况转化、中序遍历、pre指针

## 此类问题解决方案汇总

\*总题数$^8$

暴力法：遍历$^4$、拆分情况$^5$、情况转化$^2$、枚举false条件$^4$、枚举true条件$^1$、true和false条件转化$^1$

多指针法$^2$(快慢指针$^1$、双指针$^1$)、栈$^1$、HashSet$^1$、递归$^1$、pre指针$^1$

中序遍历$^1$

## 常见最优时间复杂度

$O(n)^4$

<!-- 题目链接 -->

[125题-合法回文]:125-ValidPalindrome.md
[20题-合法括号]:20-ValidParentheses.md
[141题-有环链表]:141-LinkedListCycle.md
[98题-合法BST]:98-ValidateBinarySearchTree.md
