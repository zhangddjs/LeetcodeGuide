public class Solution {
    public boolean hasCycle(ListNode head) {
        if (head == null) return false;
        ListNode p1 = head, p2 = head.next;
        while (p1 != p2) {
            if (p2 == null || p2.next == null) return false;
            p2 = p2.next.next;
            p1 = p1.next;
        }
        return true;
    }
}