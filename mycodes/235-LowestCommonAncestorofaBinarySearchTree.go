/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
var path []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil || p == nil || q == nil {
    return nil
  }
  path = make([]*TreeNode, 0)
  pathp, pathq := make([]*TreeNode, 0), make([]*TreeNode, 0)
  findPath(root, p)
  pathp = append(pathp, path...)
  pathp = append(pathp, root)
  path = make([]*TreeNode, 0)
  findPath(root, q)
  pathq = append(pathq, path...)
  pathq = append(pathq, root)
  i, j := len(pathp)-1, len(pathq)-1
  for i, j = i, j; i >= 0 && j >= 0; i, j = i-1, j-1 {
    if pathp[i] != pathq[j] {
      break
    }
  }
  return pathp[i+1]
}

func findPath(root, node *TreeNode) bool {
  if root == nil {
    return false
  }
  if root == node {
    return true
  }
  if findPath(root.Left, node) {
    path = append(path, root.Left)
    return true
  }
  if findPath(root.Right, node) {
    path = append(path, root.Right)
    return true
  }
  return false
}
