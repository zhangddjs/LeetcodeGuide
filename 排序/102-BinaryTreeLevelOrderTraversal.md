# [#102 Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal)

![Medium](/figures/Medium.svg)

## 关键词

排序、二叉树、层次遍历、队列、DFS

## 题目

Given a binary tree, return the _level order_ traversal of its nodes' values. (ie, from left to right, level by level).

## 简述

**输入：** 二叉树

**输出：** 层次遍历序列

## 思路

本题考察二叉树结点的层次遍历，为数据结构的基本知识。

将根节点入队列，作为第一层，将下一层节点(根节点的左右孩子)入队，然后从队列中取出当前层下一个要访问的节点或下一层第一个节点，并在遍历当前层每个节点时将其左右孩子入队列从而构成下一层要遍历的节点，直到队列空时所有层遍历结束。

在遍历每一层时需要记录当前层剩余未遍历的节点数量。当遍历完当前层时，队列中剩余的元素数即为下一层节点数量，此时当前层节点数量正好为0，再令其等于队列长度即可。------方法1

## 解决方案

### 方法1-队列法(BFS)

用队列维护当前层和下一层节点信息，按顺序入队出队遍历。(关键词：队列)

时间复杂度：$O(n)$ ---62%

空间复杂度：$O(n)$ ---16%

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
    public List<List<Integer>> levelOrder(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;
        queue.offer(root);
        int rest = queue.size();
        List<Integer> curLevel = new ArrayList<>();
        while (!queue.isEmpty()) {
            if (rest == 0) {
                rest = queue.size();
                res.add(curLevel);
                curLevel = new ArrayList<>();
            }
            TreeNode node = queue.poll();
            curLevel.add(node.val);
            rest--;
            if (node.left != null) queue.offer(node.left);
            if (node.right != null) queue.offer(node.right);
        }
        res.add(curLevel);
        return res;
    }
}
```

## 扩展

### 扩展方法1-递归法(DFS)[$^{[1]}$](#refer-anchor-1)

用深度优先递归遍历，记录当前遍历深度，加到对应集合中。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
 class Solution {
    List<List<Integer>> levels = new ArrayList<List<Integer>>();

    public void helper(TreeNode node, int level) {
        // start the current level
        if (levels.size() == level)
            levels.add(new ArrayList<Integer>());

         // fulfil the current level
         levels.get(level).add(node.val);

         // process child nodes for the next level
         if (node.left != null)
            helper(node.left, level + 1);
         if (node.right != null)
            helper(node.right, level + 1);
    }

    public List<List<Integer>> levelOrder(TreeNode root) {
        if (root == null) return levels;
        helper(root, 0);
        return levels;
    }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 102-Solution](https://leetcode.com/problems/binary-tree-level-order-traversal/solution/)
