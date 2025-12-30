/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
  if root == nil {
    return 0
  }
  q := make([]*TreeNode, 0)
  q = append(q, root)
  size, depth := 1, 0
  for len(q) > 0 {
    n := q[0]
    q = q[1:]
    if n.Left != nil{
      q = append(q, n.Left)
    }
    if n.Right != nil{
      q = append(q, n.Right)
    }
    size--
    if size == 0 {
      size = len(q)
      depth++
    }
  }
  return depth
}
