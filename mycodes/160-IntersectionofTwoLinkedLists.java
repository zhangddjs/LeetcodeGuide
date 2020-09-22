/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) return null;
        int lenA = 0, lenB = 0;
        ListNode pA = headA, pB = headB;
        while (pA != null) {
            pA = pA.next;
            lenA++;
        }
        while (pB != null) {
            pB = pB.next;
            lenB++;
        }
        pA = headA;
        pB = headB;
        int lenDiff = Math.abs(lenA - lenB);
        if (lenA > lenB) while (lenDiff-- != 0) pA = pA.next;
        else while (lenDiff-- != 0) pB = pB.next;
        while (pA != null) {
            if (pA == pB) return pA;
            pA = pA.next;
            pB = pB.next;
        }
        return null;
    }
}

public class Solution2 {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode pA = headA, pB = headB;
        while (pA != null || pB != null) {
            if (pA != null && pB != null && pA == pB) return pA;
            pA = pA == null ? headB : pA.next;
            pB = pB == null ? headA : pB.next;
        }
        return null;
    }
}