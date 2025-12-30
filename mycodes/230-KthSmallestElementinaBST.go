/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var kk int

func kthSmallest(root *TreeNode, k int) int {
  if k <= 0 || root == nil {
    return -1
  }
  kk = k
  return traverse(root)
}


func traverse(root *TreeNode) int {
  if kk <= 0 || root == nil {
    return -1
  }
  left := traverse(root.Left)
  if left != -1 {
    return left
  }
  kk--
  if kk == 0 {
    return root.Val
  }
  return traverse(root.Right)
}
