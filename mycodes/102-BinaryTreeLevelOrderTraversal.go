/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
  if root == nil {
    return [][]int{}
  }
  res := make([][]int, 0)
  q := make([]*TreeNode, 0)
  q = append(q, root)
  size := 1
  cur := make([]int, 0, len(q))
  for len(q) > 0 {
    node := q[0]
    q = q[1:]
    size--
    cur = append(cur, node.Val)
    if node.Left != nil {
      q = append(q, node.Left)
    }
    if node.Right != nil {
      q = append(q, node.Right)
    }
    if size == 0 {
      size = len(q)
      res = append(res, cur)
      cur = make([]int, 0, len(q))
    }
  }
  return res
}
