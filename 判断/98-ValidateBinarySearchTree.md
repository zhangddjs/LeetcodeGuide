# [#98 Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree)

![Medium](/figures/Medium.svg)

## 关键词

判断、二叉搜索树、递归、拆分情况、情况转化、中序遍历、pre指针

## 题目

Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

+ The left subtree of a node contains only nodes with keys **less than** the node's key.
+ The right subtree of a node contains only nodes with keys **greater than** the node's key.
+ Both the left and right subtrees must also be binary search trees.

## 简述

**输入：** 二叉树

**输出：** 是否是BST

## 思路

本题考察对二叉搜索树的判断，需要知道那些情况二叉树是BST，哪些情况不是。

可以知道，一个二叉树如果是BST，则必然满足如下两个情况：

+ 情况1：对于任意节点，其节点值大于左孩子且小于右孩子
+ 情况2：对于任意节点，其节点值大于左子树所有节点，小于右子树所有节点

因此本题可以拆分成子问题，然后通过递归的方式来判断每个节点是否满足如上情况，如果全都满足，则二叉树是BST。

对于情况1，只需进行简单判断即可，情况2的判断则需要遍历左右子树，这样做的话时间复杂度将很高，因此我们需要将情况2进行一个转换，即对于任意节点，该节点可能同时属于某节点的左子树和某另一节点的右子树，所以当前节点值必须小于所在左子树的父节点，且大于所在右子树的父节点。这样一来，对每个正在递归遍历的节点，我们只需要记录其所在左子树和右子树的父节点即可完成情况2的判断。------方法1

另一种方法是根据BST的特殊性质来判断，即BST的中序遍历序列是升序排序的。因此我们只需用中序遍历的方式遍历二叉树，然后在遍历时记录前一个遍历的节点值，并判断当前节点是否大于前一节点，如果不满足，则不是BST。------方法2

## 解决方案

### 方法1-分治法

递归判断二叉树每个节点是否满足上述分析的情况，如果满足，则是BST。(关键词：递归、拆分情况、情况转化)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---81%

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
    public boolean isValidBST(TreeNode root) {
        if (root == null) return true;
        boolean isValid = true;
        isValid &= root.left == null ? true : root.left.val < root.val;
        isValid &= root.right == null ? true : root.right.val > root.val;
        isValid = isValid && helper(root.left, null, root) && helper(root.right, root, null);
        return isValid;
    }

    public boolean helper(TreeNode root, TreeNode leftparent, TreeNode rightparent) {
        boolean isValid = true;
        if (root == null || (root.left == null && root.right == null)) return true;
        isValid &= root.left == null ? true : root.left.val < root.val;
        isValid &= root.right == null ? true : root.right.val > root.val;
        isValid &= root.right == null || rightparent == null ? true : root.right.val < rightparent.val;
        isValid &= root.left == null || leftparent == null ? true : root.left.val > leftparent.val;
        isValid = isValid && helper(root.left, leftparent, root) && helper(root.right, root, rightparent);
        return isValid;
    }
}
```

### 方法2-中序遍历法

中序遍历二叉树并判断每个节点是否大于上一节点，如果满足，则是BST。(关键词：中序遍历、pre指针)

时间复杂度：$O(n)$ ---100%

空间复杂度：$O(n)$ ---74%

``` java
class Solution {
    long pre = Long.MIN_VALUE;
    public boolean isValidBST(TreeNode root) {
        if (root == null) return true;
        boolean isValid = isValidBST(root.left) & root.val > pre;
        pre = root.val;
        return isValid && isValidBST(root.right);
    }
}
```
