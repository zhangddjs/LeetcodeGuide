# [#160 Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists)

![Easy](/figures/Easy.svg)

## 关键词

查找、链表、HashSet、遍历、情况转化、双指针、双指针隐藏信息

## 题目

Write a program to find the node at which the intersection of two singly linked lists begins.

## 简述

**输入：** 两个链表表头节点

**输出：** 相交点的引用

**Notes：**

+ 时间复杂度必须是O(n)，空间复杂度O(1)
+ 如果2个链表没有交点，返回空
+ 链表节点范围为`[1, 10^9]`
+ 不能改变链表结构
+ 链表无环

## 思路

本题考察对基本数据结构的查找，需要掌握基本数据结构的性质。

可以很容易想到一种暴力解法，就是遍历一遍双链表中的一个链表，然后用HashSet来存储遍历过的每个节点，再遍历第二条链表，遍历每个节点时判断HashSet中是否存在该节点，如果存在，则为交叉点，返回即可。（关键词：遍历、HashSet）

但暴力法的空间复杂度为$O(n)$，不满足题目要求的$O(1)$，所以需要进一步优化。

根据题目所给的链表数据结构，我们可以知道链表无法随机搜索的特性，根据该特性我们又能够想到多指针的方法来搜索目标，因此如何确定多指针的搜索方案将是本题的一个重点，确定了搜索方案就可以解决本问题。

首先我们分析一个特殊情况，就是如果两个链表的长度相等，如何找公共节点？答案显而易见，用两个指针同时从两个链表表头节点开始往后遍历，并判断指针指向的节点是否相同，如果相同，那么该节点就是交点，如果遍历到末尾还没找到这样的点，则认为没有交点。

在题目中，两个链表的长度可能不同，但是如果两个链表有交点，那么从该节点到尾节点的长度在两个链表中是相等的，从头节点到该节点的长度则是不等的，如果能够令后者这两段长度相等，此时情况则能够转化成如上所分析的情况，问题也就得到解决。

因此我们可以先遍历两个链表确定链表的长度差k，然后再让长链表遍历指针先走k个节点，此时令双指针同时开始遍历并判断节点是否相等即可完成搜索。------方法1

在方法1中，我们通过两次遍历，第一次遍历用于得到两链表的长度差k，得到长度差k的方法是求出两个链表的长度然后作差求得，不过这一步骤其实还可以通过双指针携带的隐藏信息进一步进行优化。

方法1的核心是长度差k的获取，经过分析可以发现，我们可以令双指针同时遍历两个链表，当有一个指针遍历到末尾时，这两个指针之间的距离边已经是我们要求的k，且未遍历结束指针遍历的链表为更长的那个。所以此时令一个指针从该链表头结点开始，让该指针和先前没遍历结束的指针继续同步遍历，直到先前的指针指向末尾，此时第二个指针则已经走了k步，后续操作则和等长链表找交叉点一致。------方法2

## 解决方案

### 方法1-多指针法

遍历两个链表得到长度差k，令长链表遍历指针先走k个节点，再用双指针同时遍历搜索。(关键词：遍历、情况转化、双指针)

时间复杂度：$O(n)$ ---99%

空间复杂度：$O(1)$ ---82%

``` java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) return null;
        int lenA = 0, lenB = 0;
        ListNode pA = headA, pB = headB;
        while (pA != null) {
            pA = pA.next;
            lenA++;
        }
        while (pB != null) {
            pB = pB.next;
            lenB++;
        }
        pA = headA;
        pB = headB;
        int lenDiff = Math.abs(lenA - lenB);
        if (lenA > lenB) while (lenDiff-- != 0) pA = pA.next;
        else while (lenDiff-- != 0) pB = pB.next;
        while (pA != null) {
            if (pA == pB) return pA;
            pA = pA.next;
            pB = pB.next;
        }
        return null;
    }
}
```

### 方法2-优化的多指针法

双指针遍历两个链表得到长度差k，令长链表遍历指针先走k个节点，再用双指针同时遍历搜索，此方法不显示记录k值。(关键词：遍历、情况转化、双指针、双指针隐藏信息)

时间复杂度：$O(n)$ ---99%

空间复杂度：$O(1)$ ---99%

``` java
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode pA = headA, pB = headB;
        while (pA != null || pB != null) {
            if (pA != null && pB != null && pA == pB) return pA;
            pA = pA == null ? headB : pA.next;
            pB = pB == null ? headA : pB.next;
        }
        return null;
    }
}
```
