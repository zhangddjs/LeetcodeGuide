/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
  if root == nil || (root.Left == nil && root.Right == nil) {
    return 0
  }
  lh, ld := diameterAndHeight(root.Left)
  rh, rd := diameterAndHeight(root.Right)
  return max(ld, rd, lh+rh)
}

func diameterAndHeight(root *TreeNode) (int, int) {
  if root == nil {
    return 0, 0
  }
  hl, dl := diameterAndHeight(root.Left)
  hr, dr := diameterAndHeight(root.Right)
  curh := 1+max(hl, hr)
  curd := max(hl+hr, dl, dr)
  return curh, curd
}
