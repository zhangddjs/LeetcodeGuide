# [#708 Insert into a Sorted Circular Linked List](https://leetcode.com/problems/insert-into-a-sorted-circular-linked-list/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、链表、环形链表、有序链表、遍历、链表插入、拆分情况

## 题目

Given a node from a **Circular Linked List** which is sorted in ascending order, write a function to insert a value `insertVal` into the list such that it remains a sorted circular list. The given node can be a reference to any single node in the list, and may not be necessarily the smallest value in the circular list.

If there are multiple suitable places for insertion, you may choose any place to insert the new value. After the insertion, the circular list should remain sorted.

If the list is empty (i.e., given node is `null`), you should create a new single circular list and return the reference to that single node. Otherwise, you should return the original given node.

## 简述

**输入：** 环形有序链表表头节点; 插入元素的值

**输出：** 插入元素后的表头节点

**Notes：**

+ 表头节点不一定是最小值，可能为链表中任意节点
+ 0 <= 节点数 <= $5 * 10^4$
+ $-10^6$ <= 节点值、插入值 <= $10^6$
+ 如果有多个合适的插入位置，随便插入即可

## 思路

本题考察基础数据结构链表的操作。

本题需要在环形有序的链表中插入一个元素，因为是有序的，所以只需从头节点往后遍历插入到合适的位置即可。(**#case 0**)

但根据链表是环形的特点，往往存在两个节点的值是逆序的，也就是链表中最小值和最大值的位置，遇到这种情况，如果插入节点值比最小值小或者比最大值大，则插入在这两个元素中间。(**#case 1**)

如果一开始链表中只有一个节点(**#case 2**)或全部节点相等(**#case 3**)，则尾节点的下一个节点指向头节点，新元素插入到尾节点和头节点之间。

如果链表是空的，则直接构造头节点即可。(**#case 4**)------方法1

## 解决方案

### 方法1-暴力法

遍历链表，拆分情况进行插入。(关键词：遍历链表、拆分情况)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---93%

``` java
/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _next) {
        val = _val;
        next = _next;
    }
};
*/
class Solution {
    public Node insert(Node head, int insertVal) {
        Node node = head;
        if (node == null) {      //#case 4
            head = new Node(insertVal);
            head.next = head;
        }
        else {
            while (node.next != head) {     //#case 2,3
                if ((node.val > node.next.val &&
                    (node.val <= insertVal ||
                     node.next.val >= insertVal)) ||
                    (node.val <= insertVal &&
                     node.next.val >= insertVal)) break;     //#case 0,1
                node = node.next;
            }
            node.next = new Node(insertVal, node.next);
        }
        return head;
    }
}
```
