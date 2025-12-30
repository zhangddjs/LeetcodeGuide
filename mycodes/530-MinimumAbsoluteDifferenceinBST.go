/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var arr []int

func getMinimumDifference(root *TreeNode) int {
  arr = make([]int, 0)
  traverse(root)
  var res int = 1000000
  for i := 1; i < len(arr); i++ {
    res = min(res, abs(arr[i]-arr[i-1]))
  }
  return res
}

func traverse(root *TreeNode) {
  if root == nil {
    return
  }
  traverse(root.Left)
  arr = append(arr, root.Val)
  traverse(root.Right)
}

func abs(a int) int {
  if a < 0 {
    return -a
  }
  return a
}
