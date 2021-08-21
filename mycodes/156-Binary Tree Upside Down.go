/**
 * The description of this problem is not clearly, it didn't say that right nodes with two child wll not do the option.
 * So here I implement this recursive function which will both operate on left and right nodes who has two childs.
 *
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}

	rightRoot := upsideDownBinaryTree(root.Right) //get the new root of right subtree
	left := root.Left                             // record the origin left child of current root
	leftRoot := upsideDownBinaryTree(left)        //get the new root of left subtree
	left.Right = root
	left.Left = rightRoot
	root.Left = nil
	root.Right = nil
	return leftRoot //the new root of left subtree is also the new root of current tree
}

// maybe AC answer. test cases are wrong
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root
	}

	left := root.Left                      // record the origin left child of current root
	leftRoot := upsideDownBinaryTree(left) //get the new root of left subtree
	left.Right = root
	left.Left = root.Right
	root.Left = nil
	root.Right = nil
	return leftRoot //the new root of left subtree is also the new root of current tree
}
