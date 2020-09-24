# [#103 Binary Tree Zigzag Level Order Traversal](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal)

![Easy](/figures/Easy.svg)
![Medium](/figures/Medium.svg)
![Hard](/figures/Hard.svg)

## 关键词

排序、二叉树、层次遍历、双端队列、DFS

## 题目

Given a binary tree, return the `zigzag level order` traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

## 简述

**输入：** 二叉树

**输出：** `Z`型层次遍历序列

## 思路

本题考察二叉树结点的层次遍历，为数据结构的基本知识。

本题和[[引用]102题](102-BinaryTreeLevelOrderTraversal.md)类似，都是按层次遍历二叉树，但本题在遍历顺序上重新做了要求，即第一层从左往右，第二层从右往左，第三层从左往右，以此类推。

因此本题的解法和102题也类似，用一个队列维护，只不过该队列改成双端队列，奇数层从前端出元素，从后端入元素，偶数层从后端出元素，从前端入元素。------方法1

当然本题同样可以用DFS法遍历，在奇数层往对应集合的后面插入，偶数层往前面插入。

## 解决方案

### 方法1-双端队列法(BFS)

用队列维护当前层和下一层节点信息，按指定顺序入队出队遍历。(关键词：双端队列)

时间复杂度：$O(n)$ ---5%

空间复杂度：$O(n)$ ---8%

``` java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;
        Deque<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offerLast(root);
        int size = queue.size(), level = 1;
        List<Integer> curLevel = new ArrayList<>();
        while (!queue.isEmpty()) {
            if (size == 0) {
                level++;
                size = queue.size();
                res.add(curLevel);
                curLevel = new ArrayList<>();
            }
            TreeNode node = level % 2 != 0 ? queue.pollFirst() : queue.pollLast();
            curLevel.add(node.val);
            if (level % 2 != 0) {
                if (node.left != null) queue.offerLast(node.left);
                if (node.right != null) queue.offerLast(node.right);
            } else {
                if (node.right != null) queue.offerFirst(node.right);
                if (node.left != null) queue.offerFirst(node.left);
            }
            size--;
        }
        res.add(curLevel);
        return res;
    }
}
```
