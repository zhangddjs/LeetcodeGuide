# [#94 Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal)

![Medium](/figures/Medium.svg)

## 关键词

排序、二叉树、中序遍历、递归、栈、Morris遍历

## 题目

Given the `root` of a binary tree, return _the inorder traversal of its nodes' values_.

**Follow up:** Recursive solution is trivial, could you do it iteratively?

## 简述

**输入：** 二叉树

**输出：** 中序遍历序列

**Notes：**

+ 0 <= 节点数 <= 100
+ -100 <= 节点值 <= 100

## 思路

本题考察二叉树结点的中序遍历，为数据结构的基本知识。

按左-根-右的顺序递归遍历，将遍历的值按顺序添加到结果集合输出即可。------方法1

迭代法需要用一个栈来维护根节点，当访问右节点前需要出栈。------方法2

## 解决方案

### 方法1-递归法

左-根-右递归。(关键词：中序遍历、递归)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---73%

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
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if(root == null) return res;
        res.addAll(inorderTraversal(root.left));
        res.add(root.val);
        res.addAll(inorderTraversal(root.right));
        return res;
    }
}
```

### 方法2-迭代法

左-根-右迭代。(关键词：中序遍历、栈)

时间复杂度：$O(n)$ ---21%

空间复杂度：$O(n)$ ---30%

``` java
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if(root == null) return res;
        Stack<TreeNode> stack = new Stack<>();
        TreeNode node = root;
        while (!stack.isEmpty() || node != null){
            while (node != null) {
                stack.push(node);
                node = node.left;
            }
            node = stack.pop();
            res.add(node.val);
            node = node.right;
        }
        return res;
    }
}
```

## 扩展

### 扩展方法-Morris遍历[$^{[1]}$](#refer-anchor-1)

遍历二叉树，如果当前节点没有左子树，则访问当前节点并遍历右子树，否则令该节点为左子树的最右孩子的右节点，然后从当前节点原先的左孩子开始继续遍历。

时间复杂度：$O(n)$

空间复杂度：$O(n)$

不过方法改变了输入二叉树结构，参考[${[2]}$](#refer-anchor-2)中提供了一个还原输入二叉树的方法。

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 94-Solution](https://leetcode.com/problems/binary-tree-inorder-traversal/solution/)

<div id="refer-anchor-1"></div>

+ [2] [Leetcode. 94-Comment](https://leetcode.com/problems/binary-tree-inorder-traversal/solution/154815)