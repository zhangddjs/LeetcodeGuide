# [#235 Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree)

![Easy](/figures/Easy.svg)

## 关键词

查找、二叉树、BST、公共节点

## 题目

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow **a node to be a descendant of itself**).”

## 简述

**输入：** BST; BST中的两个不同节点

**输出：** 最小公共祖先

**Notes：**

+ 所有节点值唯一

## 思路

本题考察二叉树的查找，通常做法需要遍历每个节点，但根据BST的性质，可以一定程度简化计算。

经过分析可以发现，对于二叉树中的任意两个节点，其公共祖先节点一定包含根节点，同时若某一节点为这两个节点的祖先节点，那么该节点的父节点也为这两节点的祖先节点。

因此我们可以得到一个暴力方法，即通过两次深度遍历寻找着两个节点，同时记录遍历路径，这两个路径中必然有一段是公共路径，路径上节点也就是公共祖先节点，公共路径上最后一个节点即为最小公共祖先。本方法也是寻找普通二叉树任意两节点最小公共节点LCA的通用暴力法。------方法1

不过根据BST的性质，我们可以在寻找目标节点时进行一个优化操作，如果目标小于当前节点，则向左子树遍历，否则遍历右子树。------方法2

在方法2按BST搜索的基础上，我们可以发现从源点到目标点的路径可以一次性找到，不需要遍历其他路径来确定该路径，因此我们可以同时遍历搜索两个目标节点，如果两个节点同时小于或大于当前节点，则说明当前节点为一个公共祖先，否则接下来的路径会在当前节点出现分岔，从而确定当前节点为最小公共祖先。------方法3

## 解决方案

### 方法1-暴力法

深度遍历两次找到两个节点，记录各自路径，寻找公共路径中最后一个节点，即为最小公共祖先节点。(关键词：DFS、记录路径)

时间复杂度：$O(n)$ ---72%

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

### 方法2-优化的暴力法

BST搜索方式搜索两次找到两个节点，记录各自路径，寻找公共路径中最后一个节点，即为最小公共祖先节点。(关键词：BST搜索)

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(\log(n))$ ---83%

``` java
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
        if (list.isEmpty() && root != desc) dfs(root.val > desc.val ? root.left : root.right, desc, list);
        if (!list.isEmpty() || root == desc) list.add(root);
    }
}
```

### 方法3-BST搜索法

按BST搜索方式同时搜索两个节点，如果两个节点同时小于或大于当前节点，则继续，否则当前节点即为最小公共祖先节点。(关键词：BST搜索、记录路径)

#### 迭代方式

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(1)$ ---56%

``` java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        while ((root.val > p.val && root.val > q.val) || (root.val < p.val && root.val < q.val))
            root = root.val > p.val ? root.left : root.right;
        return root;
    }
}
```

#### 递归方式

时间复杂度：$O(\log(n))$ ---100%

空间复杂度：$O(\log(n))$ ---83%

``` java
class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root.val > p.val && root.val > q.val) return lowestCommonAncestor(root.left, p, q);
        else if (root.val < p.val && root.val < q.val) return lowestCommonAncestor(root.right, p, q);
        else return root;
    }
}
```

## 扩展

### TODO扩展一些知识和方法[$^{[1]}$](#refer-anchor-1)

内容

``` java
/**
 * copyright: LeetCode(https://leetcode.com)
 * 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
 */
//Extension Solution
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 0-Solution]()
