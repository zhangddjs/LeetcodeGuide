# [#117 Populating Next Right Pointers in Each Node II](https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、二叉树、就地、层次遍历、双指针

## 题目

Given a binary tree

``` c++
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```

Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to `NULL`.

Initially, all next pointers are set to `NULL`.

**Follow up:**

+ You may only use constant extra space.
+ Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.

## 简述

**输入：** 二叉树

**输出：** 每一层节点连接后的满二叉树

**Notes：**

+ 树节点小于6000个
+ -100 <= 节点值 <= 100

## 思路

本题考察二叉树状态转换，需要按题目要求推出合适的方案来进行操作。

本题和[[引用]116题](116-PopulatingNextRightPointersinEachNode.md)题非常类似，区别仅在于116题中的二叉树为满二叉树。因此本题不能直接将节点的左孩子和右孩子连接，或右孩子和该节点连接的节点的左节点相连，而是需要用另一个指针向右寻找子节点所在层的下一个非空节点，与其进行连接。------方法1

## 解决方案

### 方法1-双指针法

层次遍历并令每个节点的非空左孩子、非空右孩子和子节点所在层的下一个非空的节点连接，用双指针记录当前节点和下一层已连接节点的尾部。(关键词：层次遍历、双指针)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---51%

``` java
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    public Node connect(Node root) {
        Node node = root;
        Node nextLevel = node == null ? null : node.left != null ? node.left : node.right;
        if (node != null && node.left != null && node.right != null) node.left.next = node.right;
        while (nextLevel != null) {
            node = nextLevel;
            nextLevel = null;
            Node pre = null;
            while (node != null) {
                if (nextLevel == null) nextLevel = node.left == null ? node.right : node.left;
                if (node.left != null && node.right != null) node.left.next = node.right;
                if (pre != null) pre.next = node.left == null ? node.right : node.left;
                else pre = node.right != null ? node.right : node.left;
                while (pre != null && pre.next != null) pre = pre.next;
                node = node.next;
            }
        }
        return root;
    }
}
```
