# [#23 Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists)

![Hard](/figures/Hard.svg)

## 关键词

状态转换、链表、合并、有序链表、k指针、尾插法、优先队列、问题转化、分治法

## 题目

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

_Merge all the linked-lists into one sorted linked-list and return it._

## 简述

**输入：** `k`个有序单链表

**输出：** 合并后的有序单链表

**Notes：**

+ 0 <= k <= 10$^4$
+ 0 <= 每个链表元素个数 <= 500
+ -10$^4$ <= 链表元素值 <= 10$^4$
+ 所有链表总长 <= 10$^4$

## 思路

本题考察链表合并操作。

[[引用]21题](21-MergeTwoSortedLists.md)已经提供了一个合并两个链表的解决方案，本题需要合并k个链表，那么能否用合并两个链表的方法呢？

首先我们可以想到k指针法，用k个指针并行遍历k个链表，每次取出当前k个指针中最小的值，尾插法插入到结果链表中并令该指针右移一个，当所有指针都遍历到末尾时算法结束。------方法1

另一个可以很快想到的方法是使用优先队列法，将所有链表元素全部存入优先队列后再取出组成结果。或者是全部存入集合，然后进行排序操作再组成结果链表，其时间空间复杂度是一样的。------方法2

有了优先队列的思想后，我们可以对方法1进行一个优化，从而使得寻找k个指针中的最小值时不用一一比较，只需从优先队列中取出最小值，同时将该指针下一个元素放入优先队列，用尾插法插入结果链表中。[$^{[1]}$](#refer-anchor-1)------方法3

## 解决方案

### 方法1-暴力法(多指针法)

并行遍历k链表，用尾插法优先插入当前最小元素到新链表。(关键词：k指针、尾插法)

时间复杂度：$O(kn)$ ---8% _$n$为所有元素的个数_

空间复杂度：$O(1)$ ---95% _不计结果空间，使用输入数组存储指针信息_

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
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Integer min = 0;
        while (min != null) {
            min = null;
            ListNode tmp = null;
            int index = -1;
            for (int i = 0; i < lists.length; ++i) {
                if (lists[i] != null && (min == null || lists[i].val < min)) {
                    min = lists[i].val;
                    tmp = lists[i].next;
                    index = i;
                }
            }
            if (min != null) {
                tail.next = new ListNode(min);
                tail = tail.next;
                lists[index] = tmp;
            }
        }
        return res.next;
    }
}
```

### 方法2-暴力法(优先队列法)

将k个链表插入优先队列，再取出用尾插法组成新链表。(关键词：优先队列、尾插法)

时间复杂度：$O(n\log(n))$ ---56%

空间复杂度：$O(n)$ ---49%

``` java
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Queue<ListNode> queue = new PriorityQueue((a, b) -> ((ListNode)a).val - ((ListNode)b).val);
        for (ListNode list : lists) {
            ListNode node = list;
            while (node != null) {
                queue.offer(node);
                node = node.next;
            }
        }
        while (!queue.isEmpty()) {
            tail.next = queue.poll();
            tail = tail.next;
        }
        tail.next = null;
        return res.next;
    }
}
```

### 方法3-优化的暴力法(多指针法)

用优先队列存储k指针，从而使得取最小指针时更加快速。(关键词：k指针、优先队列、尾插法)

时间复杂度：$O(n\log(k))$ ---40%

空间复杂度：$O(k)$ ---89%

``` java
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Queue<ListNode> queue = new PriorityQueue<ListNode>((a, b) -> a.val - b.val);
        Map<ListNode, Integer> map = new HashMap<>();
        for (int i = 0; i < lists.length; ++i)
            if (lists[i] != null) {
                queue.offer(lists[i]);
                map.put(lists[i], i);
            }
        while (!queue.isEmpty()) {
            ListNode min = queue.poll();
            int index = map.get(min);
            tail.next = min;
            tail = tail.next;
            lists[index] = lists[index].next;
            map.remove(min);
            if (lists[index] != null) {
                map.put(lists[index], index);
                queue.offer(lists[index]);
            }
        }
        tail.next = null;
        return res.next;
    }
}
```

## 扩展

### 扩展方法1-一个一个合并[$^{[1]}$](#refer-anchor-1)

第一个链表和第二个合并，合并后的链表再和第三个合并，一直合并到最后一个。问题转化成了合并两个有序链表的问题。(关键词：问题转化)

时间复杂度：$O(kn)$

空间复杂度：$O(1)$

### 扩展方法2-两两合并(分治法)[$^{[1]}$](#refer-anchor-1)

类似于归并排序思维。第一个链表和第二个组成一对，第三个和第四个组成一对……将每一对进行合并，产生k/2个新链表，继续进行匹配合并操作，直到最后只剩一个链表，即是结果。(关键词：分治法)

时间复杂度：$O(n\log(k))$

空间复杂度：$O(1)$

``` python
#代码见参考[1]
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 23-Solution](https://leetcode.com/problems/merge-k-sorted-lists/solution/)
