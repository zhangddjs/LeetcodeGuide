# [#21 Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists)

![Easy](/figures/Easy.svg)

## 关键词

状态转换、链表、合并、有序链表、双指针、尾插法、插入排序、递归

## 题目

Merge two sorted linked lists and return it as a new **sorted** list. The new list should be made by splicing together the nodes of the first two lists.

## 简述

**输入：** 两个有序链表

**输出：** 合并后的有序链表

## 思路

本题考察链表合并操作。

拿到题目可以立马想到一种暴力解法，就是通过双指针来并行遍历两个链表，将双指针指向的较小的那个元素用尾插法添加到新链表中，然后该指针向后移动，直到一个链表遍历完成。此时直接将剩下的元素插入新链表并返回即可。------方法1

另一种方法就是将一个链表用插入排序插入到另一个链表中，因为链表有序，每次插入时从上一个插入位置开始扫描当前插入位置，从而线性时间内完成插入，并能够就地完成合并操作。------方法2

## 解决方案

### 方法1-暴力法

并行遍历双链表，用尾插法优先插入较小元素到新链表。(关键词：双指针、尾插法)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---69% _不计结果空间_

``` java
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode res = new ListNode(0);
        ListNode tail = res, p1 = l1, p2 = l2;
        while (p1 != null && p2 != null) {
            tail.next = new ListNode(p1.val < p2.val ? p1.val : p2.val);
            tail = tail.next;
            if (p1.val < p2.val) p1 = p1.next;
            else p2 = p2.next;
        }
        tail.next = p1 == null ? p2 : p1;
        return res.next;
    }
}
```

### 方法2-插入排序法

遍历某一链表，用插入排序思维将其插入到另一链表。(关键词：双指针、插入排序)

时间复杂度：$O(n)$ ---31%

空间复杂度：$O(1)$ ---9% _不计结果空间_

``` java
//本方法的实现可以进行改进和浓缩，详见参考[1]
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) return l1 == null ? l2 : l1;
        ListNode p1 = l1, p2 = l2;
        if (p2.val < p1.val) {
            ListNode tmp = p2;
            p2 = p2.next;
            tmp.next = p1;
            l1 = tmp;
            p1 = l1;
        }
        while (p2 != null && p1.next != null) {
            ListNode tmp = p2;
            p2 = p2.next;
            while (p1.next != null && p1.next.val < tmp.val) p1 = p1.next;
            tmp.next = p1.next;
            p1.next = tmp;
            p1 = p1.next;
        }
        if (p2 != null) p1.next = p2;
        return l1;
    }
}
```

## 扩展

### 扩展方法-递归法[$^{[1]}$](#refer-anchor-1)

可以将方法1转化成一种递归的形式来实现。

``` java
//代码详见参考[1]
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 21-Solution](https://leetcode.com/problems/merge-two-sorted-lists/solution/)
