/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
  if root == nil {
    return true
  }
  hl, hr := getH(root.Left), getH(root.Right)
  if abs(hl - hr) <= 1 {
    return isBalanced(root.Left) && isBalanced(root.Right)
  }
  return false
}

func getH(root *TreeNode) int {
  if root == nil {
    return 0
  }
  q := make([]*TreeNode, 0)
  q = append(q, root)
  size, dep := 1, 0
  for len(q) > 0 {
    n := q[0]
    q = q[1:]
    if n.Left != nil {
      q = append(q, n.Left)
    }
    if n.Right != nil {
      q = append(q, n.Right)
    }
    size--
    if size == 0 {
      size = len(q)
      dep++
    }
  }
  return dep
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}
