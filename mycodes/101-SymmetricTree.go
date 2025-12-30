/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
  if root == nil {
    return true
  }
  q := make([]*TreeNode, 0)
  nextq := make([]*TreeNode, 0)
  q = append(q, root)
  for len(q) > 0 {
    node := q[0]
    q = q[1:]
    if node != nil {
      nextq = append(nextq, node.Left, node.Right)
    }
    if len(q) == 0 {
      if !isParlindrone(nextq) {
        return false
      }
      q = nextq
      nextq = make([]*TreeNode, 0)
    }
  }
  return true
}

func isParlindrone(q []*TreeNode) bool {
  i, j := 0, len(q)-1
  for i < j {
    if q[i] != nil && q[j] != nil && q[i].Val != q[j].Val {
      return false
    } else if q[i] != q[j] && (q[i] == nil || q[j] == nil) {
      return false
    }
    i++
    j--
  }
  return true
}
