/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []int

func isValidBST(root *TreeNode) bool {
  res = make([]int, 0)
  traverse(root)
  for i := 1; i < len(res); i++ {
    if res[i] <= res[i-1] {
      return false
    }
  }
  return true
}

func traverse(root *TreeNode) {
  if root == nil {
    return
  }
  traverse(root.Left)
  res = append(res, root.Val)
  traverse(root.Right)
}
