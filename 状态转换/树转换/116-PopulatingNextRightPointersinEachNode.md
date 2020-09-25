# [#116 Populating Next Right Pointers in Each Node](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、二叉树、满二叉树、就地、递归、层次遍历

## 题目

You are given a **perfect binary tree** where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

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

**输入：** 满二叉树

**输出：** 每一层节点连接后的满二叉树

**Notes：**

+ 树节点小于4096个
+ -1000 <= 节点值 <= 1000

## 思路

本题考察二叉树状态转换，需要按题目要求推出合适的方案来进行操作。

因为只需连接每一层的节点，所以很快可以想到通过深度优先和层次遍历的方法进行节点连接。但根据题目要求除递归栈外只能常数空间复杂度，因此需要尝试进行其他方法。

经过分析可以发现，对于每个节点，如果其是父节点的左孩子，则连接其兄弟节点，否则连接其父节点连接的节点的左孩子。根据该子问题，我们可以开始用递归(DFS)来进行连接。------方法1

递归法确实很好，但其空间复杂度还是不满足常数复杂度，因此我们再进一步考虑如何减少空间复杂度。

我们可以知道，不管用DFS还是用BFS(层次)遍历，都需要额外非常数空间，传统的层次遍历需要额外空间的原因是因为每一层的节点之间并没有关联，需要通过队列来维持当前一层和下一层的节点顺序。

因此也可以知道，下一层的节点顺序是由其上一层节点顺序确定的，所以如果在遍历当前层时知道上一层的节点顺序，当前层节点顺序也就知道了。那么有没有办法不通过队列维持上一层节点顺序呢？答案是有的，即如果上一层节点已经按顺序相连，顺序信息也就知道了。

在本题中，我们要做的恰恰是对每一层节点进行关联，当前一层连接后，下一层不需要借助额外空间也可以进行连接。我们继续尝试按照层次遍历的思想，对于当前层每个节点，令其左孩子和其右孩子连接，令其右孩子和其连接的节点(当前层下一个节点)的左孩子连接，而其连接的节点又可以在遍历上一层时确定。当遍历到最后一层时，算法结束。------方法2

## 解决方案

### 方法1-分治法(递归DFS)

递归连接每个左孩子节点和其兄弟节点，每个右孩子节点和其父节点连接的节点的左孩子节点。(关键词：递归)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---17%

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
        return helper(root, null);
    }

    public Node helper(Node root, Node brother) {
        if (root == null) return null;
        root.next = brother;
        helper(root.left, root.right);
        helper(root.right, brother == null ? null : brother.left);
        return root;
    }
}
```

### 方法2-层次遍历法(BFS)

按层次将该层每个节点的左孩子和其兄弟节点连接，将右孩子和该节点连接的节点的左孩子连接。(关键词：层次遍历)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(1)$ ---99%

``` java
class Solution {
    public Node connect(Node root) {
        if (root == null) return root;
        Node nextFirst = root.left;
        Node node = root;
        while (nextFirst != null) {
            node.left.next = node.right;
            node.right.next = node.next == null ? null : node.next.left;
            node = node.next;
            if (node == null) {
                node = nextFirst;
                nextFirst = node.left;
            }
        }
        return root;
    }
}
```
