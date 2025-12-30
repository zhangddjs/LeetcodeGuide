/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
  if head == nil {
    return head
  }
  s, f := head, head
  for f != nil && f.Next != nil {
    f = f.Next.Next
    s = s.Next
  }
  return s
}
