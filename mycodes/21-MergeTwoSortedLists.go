/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  if list1 == nil {
    return list2
  } else if list2 == nil {
    return list1
  }
  p := list1
  if list1.Val < list2.Val {
    n := mergeTwoLists(list1.Next, list2)
    p.Next = n
  } else {
    p = list2
    n := mergeTwoLists(list1, list2.Next)
    p.Next = n
  }
  return p
}
