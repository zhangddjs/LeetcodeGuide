# [#206 Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

![Easy](/figures/Easy.svg)

## 关键词

状态转换、链表、遍历、栈、递归、尾插法、头插法

## 题目

Reverse a singly linked list.

**Follow up:**

A linked list can be reversed either iteratively or recursively. Could you implement both?

## 简述

**输入：** 单链表头节点

**输出：** 逆置后的单链表头节点

## 思路

本题考察基本数据结构操作。

链表逆置的暴力方法就是借助一个栈空间，然后遍历所有元素并入栈，再将栈中所有元素采用尾插法插入到新链表中，从而实现逆置操作。------方法1

当然，链表逆置还可以通过头插法就地进行逆置，在遍历元素的过程中，将当前元素作为头节点插入链表，遍历结束后便可得到逆置的链表。------方法2

以上两种方法都是通过迭代实现的逆置操作，本题还可以通过递归的方法来实现逆置，其思想和方法1类似，采用尾插法来实现逆置操作。------方法3

## 解决方案

### 方法1-暴力法

遍历链表并入栈，将栈中元素用尾插法形式组成逆置链表。(关键词：遍历、栈、尾插法)

时间复杂度：$O(n)$ ---8%

空间复杂度：$O(n)$ ---90%

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
    public ListNode reverseList(ListNode head) {
        if (head == null) return head;
        Stack<ListNode> stack = new Stack<>();
        while (head != null) {
            stack.push(head);
            head = head.next;
        }
        ListNode res = stack.pop();
        ListNode tail = res;
        while (!stack.isEmpty()) {
            tail.next = stack.pop();
            tail = tail.next;
        }
        tail.next = null;
        return res;
    }
}
```

### 方法2-头插法逆置

遍历链表，将当前元素用头插法形式插入表头位置。(关键词：遍历、头插法)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---96%

``` java
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode tail = head, pre = head, res = head;
        while (tail != null && tail.next != null) {
            res = tail.next;
            tail.next = res.next;
            res.next = pre;
            pre = res;
        }
        return res;
    }
}
```

### 方法3-尾插法递归(方法1递归版)

通过递归的方式进行逆置。(关键词：递归、尾插法)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---46%

``` java
class Solution {
    ListNode res;
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;
        reverse(head).next = null;
        return res;
    }

    public ListNode reverse(ListNode head) {
        if (head.next == null) {
            res = head;
            return head;
        }
        ListNode tail = reverse(head.next);
        tail.next = head;
        tail = tail.next;
        return tail;
    }
}
```
