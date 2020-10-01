# [#105 Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal)

![Medium](/figures/Medium.svg)

## 关键词

状态转换、构造二叉树、HashMap、双指针、子问题、规律、递归

## 题目

Given preorder and inorder traversal of a tree, construct the binary tree.

## 简述

**输入：** 二叉树前序中序遍历序列

**输出：** 构造的二叉树

**Notes：**

+ 假设不存在重复节点

## 思路

本题考察根据前序中序遍历序列构造二叉树，为数据结构基本操作。

虽然平时我们对根据遍历序列构造二叉树的操作流程很熟悉，但并没有尝试过用程序去实现这样的流程，由此可见这样的基本数据结构操作往往是一个容易忽略的重点和难点。

可以知道，前序遍历的顺序为根-左-右，中序遍历的序列为左-根-右，于是在遍历序列中，左子树和根的遍历顺序相反(逆置)，右子树则不变。

我们尝试从最简单的开始构造，即输入节点数长度为1，则直接令其为根节点并返回。如果长度为2，则判断序列是否相等，如果相等，说明第2个节点为第1个节点的右孩子，否则为左孩子。如果长度为3，则会出现多种情况，即第3个节点为前两个节点中根节点的根节点或子节点，或者前两个节点中子节点的子节点。从这些情况中找到关联以及对应的处理方法并不是很容易，因此需要尝试其他的路线。

继续观察遍历顺序可以发现，前序遍历的第一个节点必然是根节点，中序遍历的第一个节点必然是最左边的节点(可能是最左的左孩子，也可能是最左的根节点)。

**第一步：** 根据这个性质，我们从左往右遍历前序遍历序列，直到遇到中序遍历序列第一个节点，此时的路径恰好是从根节点一路向左走到了最左端，因此令每一个节点为上一个节点的左孩子，就可以构造出二叉树根节点到最左节点之间的一条边。

构造好第一条边后，中序遍历序列的第2个(下一个)节点将是最左(当前)节点所在左子树的父节点或当前节点右子树的最左节点。如果该节点已经在构造好的边(已遍历过的前序序列节点)中，则为所在左子树的父节点，否则为右子树最左节点。

而在前序遍历序列中，下一个节点为最近的已遍历过且拥有右孩子的节点的右孩子节点。

**第二步：** 那么在中序遍历序列中，如果下一节点是当前节点右子树最左节点，则此时对应的前序遍历序列中，下一个节点为当前节点(最近的已遍历过且拥有右孩子的节点)的右孩子节点。否则，代表中序序列中当前节点没有右孩子节点，继续往后遍历，直到下一个节点为其右子树最左节点，连接该节点和前序序列下一节点(右子树根节点)，此时两个序列中的下一个节点在子树中代表的意义和初始时一致，于是从第一步开始循环构造即可完成对整个二叉树的构造。

根据以上分析，我们可以定义HashMap来存储遍历过的路径，用双指针按照第一步和第二步的要求遍历前序和中序序列来完成二叉树的构建。------方法1

在前序遍历序列中，第一个为根节点，其在中序序列中对应索引前面的所有节点则构成左子树节点，后面的所有节点构成右子树节点，从而完成序列拆分。此时用前序序列第一个节点构造根节点并弹出，弹出后的前序序列第一个节点为左或右子树根节点，此时可在中序遍历序列的左或右子树节点子数组中继续拆分。据此，我们可以通过递归的方式完成构造。------方法2[$^{[1]}$](#refer-anchor-1)

## 解决方案

### 方法1-双指针法

用HashMap存储遍历过的节点，使用双指针循环构造每个子树根节点到最左下节点的连接。(关键词：HashMap、双指针、子问题、规律)

时间复杂度：$O(n)$ ---99%

空间复杂度：$O(n)$ ---100%

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
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        Map<Integer, TreeNode> map = new HashMap<>();
        int p1 = 0, p2 = 0;
        TreeNode dummy = new TreeNode();
        TreeNode cur = dummy;
        while (p1 < preorder.length && p2 < inorder.length) {
            while (p2 < inorder.length && map.containsKey(inorder[p2])) p2++;
            if (p2 >= inorder.length) break;
            else if (p2 != 0) cur = map.get(inorder[p2 - 1]);
            cur.right = new TreeNode(preorder[p1]);
            cur = cur.right;
            map.put(preorder[p1++], cur);
            while (p1 < preorder.length && preorder[p1 - 1] != inorder[p2]) {
                cur.left = new TreeNode(preorder[p1]);
                cur = cur.left;
                map.put(preorder[p1++], cur);
            }
        }
        return dummy.right;
    }
}
```

### 方法2-分治法[$^{[1]}$](#refer-anchor-1)

递归构造每个子树的根节点，弹出前序序列节点，拆分中序序列。(关键词：子问题、递归)

时间复杂度：$O(n^2)$    //可以用空间O(n)的map降低到$O(n)$

空间复杂度：$O(n)$

``` python
##
 # copyright: LeetCode(https://leetcode.com)
 # 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
##
def buildTree(self, preorder, inorder):
    if inorder:
        ind = inorder.index(preorder.pop(0))
        root = TreeNode(inorder[ind])
        root.left = self.buildTree(preorder, inorder[0:ind])
        root.right = self.buildTree(preorder, inorder[ind+1:])
        return root
```

## 扩展

### 扩展方法-优化的分治法[$^{[2]}$](#refer-anchor-2)

在方法2中最坏情况为$O(n^2)$，如果使用map存储每个元素的索引，将会额外消耗$O(n)$的空间。因此可以使用指针的方法来取代数组拆分。

``` python
##
 # copyright: LeetCode(https://leetcode.com)
 # 代码版权归LeetCode(https://leetcode.com)和力扣中国(https://leetcode-cn.com/)所有
##
def buildTree(self, preorder, inorder):
    def build(stop):
        if inorder and inorder[-1] != stop:
            root = TreeNode(preorder.pop())
            root.left = build(root.val)
            inorder.pop()
            root.right = build(stop)
            return root
    preorder.reverse()  # make cheap inorder pop
    inorder.reverse()
    return build(None)
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 105-Discuss](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/34579/Python-short-recursive-solution.)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 105-Discuss](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/34543/Simple-O(n)-without-map)