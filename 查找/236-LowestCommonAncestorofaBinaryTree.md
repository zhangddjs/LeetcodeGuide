# [#236 Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree)

![Medium](/figures/Medium.svg)

## 关键词

查找、二叉树、LCA、DFS、记录路径、递归

## 题目

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow **a node to be a descendant of itself**).”

## 简述

**输入：** 二叉树; 二叉树中两个不同节点

**输出：** 两个节点最小公共祖先

**Notes：**

+ 所有节点值唯一

## 思路

本题考察二叉树的查找，通常做法需要遍历每个节点。

本题和[[引用]235题](235-LowestCommonAncestorofaBinarySearchTree.md)很相似，都是求最小祖先节点，但本题给的二叉树不是BST，因此不能使用BST的性质简化计算。因此对于本题我们可以直接使用暴力法，即通过两次DFS遍历找到输入条件中的两个点并记录其到根节点的路径，然后再从路径中找到公共路径的最后一个节点，即为最小公共祖先。------方法1

在得到暴力法后还可以继续进行优化，可以将整个问题划分为子问题，即对于一个子树，如果该子树包含了两个节点，则返回其包含两个节点的子树的遍历结果（即为LCA），如果只包含一个节点，返回该节点，如果不包含这两个节点，返回空。------方法2[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-暴力法

DFS遍历两次找到两个节点，记录其路径，返回公共路径最后一个节点。(关键词：DFS、记录路径)

时间复杂度：$O(n)$ ---43%

空间复杂度：$O(n)$ ---93%

``` java
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        List<TreeNode> list1 = new ArrayList<>(), list2 = new ArrayList<>();
        dfs(root, p, list1);
        dfs(root, q, list2);
        int i = list1.size() - 1, j = list2.size() - 1;
        for (i = i, j = j; i >= 1 && j >= 1 && list1.get(i - 1) == list2.get(j - 1); i--, j--);
        return list1.get(i);
    }

    public void dfs (TreeNode root, TreeNode desc, List<TreeNode> list) {
        if (root == null) return;
        if (list.isEmpty() && root != desc) dfs(root.left, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
        if (list.isEmpty() && root != desc) dfs(root.right, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
    }
}
```

### 方法2-分治法[$^{[1]}$](#refer-anchor-1)

递归左右子树，返回同时包含两个节点的最小子树的根节点。(关键词：递归)

时间复杂度：$O(n)$

空间复杂度：$O(n)$

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
    if (root == null || root == p || root == q) return root;
    TreeNode left = lowestCommonAncestor(root.left, p, q);
    TreeNode right = lowestCommonAncestor(root.right, p, q);
    return left == null ? right : right == null ? left : root;
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 236-Discuss](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/65225/4-lines-C++JavaPythonRuby)
