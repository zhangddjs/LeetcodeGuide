/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
  if root != nil && subRoot != nil {
    return traverseSync(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
  }
  return root == subRoot
}

func traverseSync(root *TreeNode, subRoot *TreeNode) bool {
  if root != nil && subRoot != nil  {
    return root.Val == subRoot.Val && traverseSync(root.Left, subRoot.Left) && traverseSync(root.Right, subRoot.Right)
  }
  return root == subRoot
}
