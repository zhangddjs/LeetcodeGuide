/**
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

    rightRoot := upsideDownBinaryTree(root.Right)
    leftRoot := upsideDownBinaryTree(root.Left)
    leftRoot.Right = root
    leftRoot.Left = rightRoot
    return leftRoot
}