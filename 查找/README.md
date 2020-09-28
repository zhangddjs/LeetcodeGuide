# 查找

## 哪些情形-解决方案

### 求目标元素或元素索引$^5$

+ 简单题$^3$

  + [852题-山峰数组峰顶索引]-遍历、哨兵、数组有序、二分法、单峰问题、黄金分割查找

  + [160题-两链表交点]-HashSet、遍历、情况转化、双指针、双指针隐藏信息

  + [235题-BST的LCA]-LCA、DFS、记录路径、BST搜索

+ 中等题$^2$

  + [702题-求有序未知大小数组的目标元素索引]-二分查找、滑动窗口优化

  + [236题-BT的LCA]-LCA、DFS、记录路径、递归

### 求满足条件的元素或元素组合的索引$^2$

+ 简单题$^1$

  + [1题-TwoSum]-遍历、枚举、HashMap

+ 中等题$^1$

  + [![[引用][求最值]](/figures/Ref-MaximumAndMinimum.svg) 5题-最长回文子串](/求最值/间接求最值/5-LongestPalindromicSubstring.md)-遍历、哨兵、三指针、动态规划、LCS、马拉车算法

## 此类问题解决方案汇总

\*总题数$^7$

暴力法：遍历$^4$、哨兵$^2$、枚举$^1$

LCA$^2$

二分法$^2$、数组有序$^2$、数组大小未知$^1$、单峰问题$^1$、黄金分割查找$^1$、滑动窗口优化$^1$、BST搜索$^1$

HashMap$^1$、HashSet$^1$、多指针法$^2$(双指针$^1$、三指针$^1$)、动态规划$^1$、LCS$^1$、马拉车算法$^1$、双指针隐藏信息$^1$、DFS$^2$、递归$^1$

情况转化$^1$、记录路径$^2$

## 常见最优时间复杂度

$O(\log(n))^3$、$O(n)^3$

<!-- 题目链接 -->
[852题-山峰数组峰顶索引]:852-PeakIndexinaMountainArray.md
[702题-求有序未知大小数组的目标元素索引]:702-SearchinaSortedArrayofUnknownSize.md
[1题-TwoSum]:1-TwoSum.md
[160题-两链表交点]:160-IntersectionofTwoLinkedLists.md
[235题-BST的LCA]:235-LowestCommonAncestorofaBinarySearchTree.md
[236题-BT的LCA]:236-LowestCommonAncestorofaBinaryTree.md
