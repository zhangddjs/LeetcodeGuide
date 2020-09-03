# [#701 InsertintoaBinarySearchTree](https://leetcode.com/problems/insert-into-a-binary-search-tree/)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、二叉树、BST、二叉树遍历、作为叶子节点插入、递归

## 题目

Given the root node of a binary search tree (BST) and a value to be inserted into the tree, insert the value into the BST. Return the root node of the BST after the insertion. It is guaranteed that the new value does not exist in the original BST.

Note that there may exist multiple valid ways for the insertion, as long as the tree remains a BST after insertion. You can return any of them.

## 简述

**输入：** BST根节点; 插入的值

**输出：** 插入元素后的BST根节点

**Notes：**

+ BST节点数量在0和$10^4$之间
+ 每个节点是唯一的，要插入的值也是新的
+ 元素大小在$-10^8$和$10^8$之间

## 思路

本题考察二叉树搜索树操作，需要熟悉二叉搜索树的添删改查流程。

常规方法是从根节点开始深度优先遍历，如果当前节点大于要插入节点，则遍历左子树，否则遍历右子树，直到即将遍历的节点为空时，将元素插入该位置即可。这种方法新元素一般会插入在树底，作为一个叶子节点插入到二叉树中。（也就是搜索的流程）------方法1

## 解决方案

### 方法1-作为叶子节点插入

从根节点向左右子树不断遍历，作为叶子节点插入。(关键词：二叉树遍历、作为叶子节点插入)

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(1)$ ---94%

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
    public TreeNode insertIntoBST(TreeNode root, int val) {
        if (root == null) return new TreeNode(val);
        TreeNode node = root, parent = null;
        while (node.val != val) {
            if (node.val > val && node.left == null) node.left = new TreeNode(val);
            else if (node.val < val && node.right == null) node.right = new TreeNode(val);
            node = node.val > val ? node.left : node.right;
        }
        return root;
    }
}
```

## 扩展

### 扩展方法1-递归插入[$^{[1]}$](#refer-anchor-1)

通过递归的方式搜索并插入。(关键词：递归)

时间复杂度：$O(\log(n))$

空间复杂度：$O(\log(n))$

``` java
class Solution {
  public TreeNode insertIntoBST(TreeNode root, int val) {
    if (root == null) return new TreeNode(val);

    // insert into the right subtree
    if (val > root.val) root.right = insertIntoBST(root.right, val);
    // insert into the left subtree
    else root.left = insertIntoBST(root.left, val);
    return root;
  }
}
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 701-Solution](https://leetcode.com/problems/insert-into-a-binary-search-tree/solution/)
