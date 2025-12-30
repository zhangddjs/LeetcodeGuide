/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
  if root == nil {
    return nil
  }
  q := make([]*TreeNode, 0)
  q = append(q, root)
  for len(q) > 0 {
    p := q[0]
    q = q[1:]
    if p.Left != nil {
      q = append(q, p.Left)
    }
    if p.Right != nil {
      q = append(q, p.Right)
    }
    p.Right, p.Left = p.Left, p.Right
  }
  return root
}
