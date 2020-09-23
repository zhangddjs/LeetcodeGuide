# [#141 Linked List Cycle](https://leetcode.com/problems/linked-list-cycle)

![Easy](/figures/Easy.svg)

## 关键词

判断、链表、快慢指针、HashSet

## 题目

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that `pos` is not passed as a parameter**.

Return `true` _if there is a cycle in the linked list_. Otherwise, return `false`.

**Follow up:**

Can you solve it using `O(1)` (i.e. constant) memory?

## 简述

**输入：** 链表表头节点

**输出：** 是否有环

**Notes：**

+ 节点数范围`[0, 10^4]`
+ -10$^5$ <= 节点值 <= 10$^5$

## 思路

本题考察判断，需要知道true和false的达成条件，然后根据情况进行处理。

判断链表是否有环可以通过快-慢指针的方式，用双指针对链表进行遍历，一个指针每次走2步，另一个每次走一步，如果链表有环，那么这两个指针必然相遇(已有证明)。因此如果遍历过程中两个指针相遇，则返回`true`，否则快指针必定先遍历到结尾，此时返回`false`。------方法1

## 解决方案

### 方法1-快慢指针法

快慢指针遍历链表，若在遍历完之前相遇则有环，否则无环。(关键词：快慢指针)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---56%

``` java
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) return false;
        ListNode p1 = head, p2 = head.next;
        while (p1 != p2) {
            if (p2 == null || p2.next == null) return false;
            p2 = p2.next.next;
            p1 = p1.next;
        }
        return true;
    }
}
```

## 扩展

### 扩展方法-哈希表法[$^{[1]}$](#refer-anchor-1)

用哈希表记录节点是否访问过，如果遍历时遇到访问过的节点，则说明存在环。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
public boolean hasCycle(ListNode head) {
    Set<ListNode> nodesSeen = new HashSet<>();
    while (head != null) {
        if (nodesSeen.contains(head)) {
            return true;
        } else {
            nodesSeen.add(head);
        }
        head = head.next;
    }
    return false;
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 141-Solution](https://leetcode.com/problems/linked-list-cycle/solution/)
