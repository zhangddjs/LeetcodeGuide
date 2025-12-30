/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  res := &ListNode{
    Val: 0,
    Next: head,
  }
  tail := head
  for tail.Next != nil {
    p := tail.Next
    tail.Next = p.Next
    p.Next = res.Next
    res.Next = p
  }
  return res.Next
}
