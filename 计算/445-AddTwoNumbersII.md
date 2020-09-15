# [#445 Add Two Numbers II](https://leetcode.com/problems/add-two-numbers-ii)

![Medium](/figures/Medium.svg)

## 关键词

计算、链表、遍历、栈、头插法、逆置思想、多次遍历、逆置、就地存取进位信息

## 题目

You are given two **non-empty** linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Follow up:**

What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

## 简述

**输入：** 两个代表正数数字的链表

**输出：** 两链表数字求和后的数字链表

## 思路

本题考察链表加法计算。

本题是[[引用]2题](2-AddTwoNumbers.md)的扩展，如果题目允许修改链表，那么我们就可以先进行链表逆置[[引用]206题](/状态转换/基本数据结构转换/206-ReverseLinkedList.md)，然后这道题便转化成了第2题。

不过这道题不能修改输入链表，因此还需要寻找其它可行的办法。

对于单链表来说，从当前节点找到后面节点容易，反之则难。如果我们按顺序遍历，从最高位开始相加，那么很可能到某个元素时候出现进位，导致前面的元素也出现进位，此时回溯修改前面遍历过的元素将会很麻烦。因此需要从最低位开始遍历进行相加。

因此虽然不可以逆置链表，但借助逆置的思想仍然是可行的。我们可以遍历两链表并将元素存入两个栈中，遍历完后栈顶元素将是最低位，同时出栈进行相加操作，并将加上进位后的结果用头插法插入新链表，当两栈都空时返回表头即可。------方法1

## 解决方案

### 方法1-暴力法

遍历双链表，推入双栈，出栈求和，头插法插入输出链表。(关键词：遍历、栈、头插法、逆置思想)

时间复杂度：$O(n)$ ---31%

空间复杂度：$O(n)$ ---79%

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
        ListNode res = new ListNode(0);
        ListNode p1 = l1, p2 = l2;
        int carry = 0;
        Stack<Integer> stack1 = new Stack<>();
        Stack<Integer> stack2 = new Stack<>();
        while (p1 != null) {
            stack1.push(p1.val);
            p1 = p1.next;
        }
        while (p2 != null) {
            stack2.push(p2.val);
            p2 = p2.next;
        }
        while (!stack1.isEmpty() || !stack2.isEmpty() || carry != 0) {
            int x = stack1.isEmpty() ? 0 : stack1.pop();
            int y = stack2.isEmpty() ? 0 : stack2.pop();
            x += y + carry;
            carry = x / 10;
            ListNode node = new ListNode(x % 10);
            node.next = res.next;
            res.next = node;
        }
        return res.next;
    }
}
```

## 扩展

### 扩展方法-就地实现[$^{[1]}$](#refer-anchor-1)

从高位向低位进行相加操作也是可以的，同时既然不允许逆置输入，那么逆置输出也可以。过程如下：

首先遍历两链表找到短链表最高位在长链表中的位置;

再遍历一遍，到达共同最高位时直接对每对节点中的值进行相加操作，用头插法构建新链表，遍历完后新链表从低位到高位排序，同时每个节点存储了进位信息;

然后再对新链表进行遍历，每个节点加上前一节点的进位信息，同时在前一节点删除进位信息，并用头插法将当前节点插入到链表头(逆置)，遍历完后如果还有进位，则再在头节点前插入节点来存储。(关键词：多次遍历、逆置、就地存取进位信息、头插法)

时间复杂度：$O(n)$

空间复杂度：$O(1)$ _不计输出链表空间_

``` c++
//代码见参考[1]
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 445-Discuss](https://leetcode.com/problems/add-two-numbers-ii/discuss/92624/C++-O(1)-extra-space-except-for-output.-Reverse-output-instead.-Is-this-cheating)
