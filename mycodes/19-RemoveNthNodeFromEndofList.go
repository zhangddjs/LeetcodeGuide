/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  p := &ListNode{0, head}
  s, f := p, p
  for i := 0; i < n; i++ {
    f = f.Next
  }
  for f.Next != nil {
    f = f.Next
    s = s.Next
  }
  s.Next = s.Next.Next
  return p.Next
}
