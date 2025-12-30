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
  pre, cur := head, head.Next
  for cur != nil {
    if cur.Val == pre.Val {
      pre.Next = cur.Next
    } else {
      pre = pre.Next
    }
    cur = cur.Next
  }
  return head
}
