/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  cur := head.Next
  for cur != nil && head.Val == cur.Val {
    cur = cur.Next
  }
  if cur != head.Next {
    return deleteDuplicates(cur)
  }
  head.Next = deleteDuplicates(cur)
  return head
}
