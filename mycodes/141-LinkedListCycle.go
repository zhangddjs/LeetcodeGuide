/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
  if head == nil || head.Next == nil {
    return false
  }
  s, f := head, head
  for f.Next != nil && f.Next.Next != nil {
    s = s.Next
    f = f.Next.Next
    if f == s {
      return true
    }
  }
  return false
}
