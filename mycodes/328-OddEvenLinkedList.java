/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode oddEvenList(ListNode head) {
        if (head == null || head.next == null || head.next.next == null) return head;
        ListNode oddtail = head, eventail = head.next;
        while (eventail != null && eventail.next != null) {
            ListNode node = eventail.next;
            eventail.next = node.next;
            eventail = eventail.next;
            node.next = oddtail.next;
            oddtail.next = node;
            oddtail = oddtail.next;
        }
        return head;
    }
}