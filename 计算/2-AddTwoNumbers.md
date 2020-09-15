# [#2 Add Two Numbers](https://leetcode.com/problems/add-two-numbers)

![Medium](/figures/Medium.svg)

## 关键词

计算、链表、遍历、基础数学

## 题目

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

## 简述

**输入：** 两个代表正数数字的链表

**输出：** 两链表数字求和后的数字链表

## 思路

本题考察链表加法计算。

因为链表中数字逆序，也就是说头节点个位数，第二个节点十位数，依次类推。因此我们可以直接同时从头结点遍历两个链表，用临时变量记录进位，将每个位上两节点与进位的相加结果作为节点存入新链表，或更新进某一输入链表(就地)，如果遍历到较长的链表最后时进位非0，则再插入一个新节点。------方法1

## 解决方案

### 方法1-暴力法

从头节点并行遍历两链表，按位相加。(关键词：遍历、基础数学)

时间复杂度：$O(n)$ ---66%

空间复杂度：$O(1)$ ---26%

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
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) return l1 == null ? l2 : l1;
        ListNode p1 = l1, p2 = l2, pre = l1;
        int carry = 0;
        while (p1 != null && p2 != null) {
            p1.val += p2.val + carry;
            carry = p1.val / 10;
            p1.val %= 10;
            pre = p1;
            p1 = p1.next;
            p2 = p2.next;
        }
        if (p1 == null) {
            pre.next = p2;
            p1 = p2;
        }
        while (carry != 0) {
            if (p1 == null) {
                pre.next = new ListNode(carry);
                break;
            } else {
                p1.val += carry;
                carry = p1.val / 10;
                p1.val %= 10;
                pre = p1;
                p1 = p1.next;
            }
        }
        return l1;
    }
}
```
