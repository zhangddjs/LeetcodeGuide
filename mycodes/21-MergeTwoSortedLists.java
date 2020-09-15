class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode res = new ListNode(0);
        ListNode tail = res, p1 = l1, p2 = l2;
        while (p1 != null && p2 != null) {
            tail.next = new ListNode(p1.val < p2.val ? p1.val : p2.val);
            tail = tail.next;
            if (p1.val < p2.val) p1 = p1.next;
            else p2 = p2.next;
        }
        tail.next = p1 == null ? p2 : p1;
        return res.next;
    }
}

class Solution2 {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) return l1 == null ? l2 : l1;
        ListNode p1 = l1, p2 = l2;
        if (p2.val < p1.val) {
            ListNode tmp = p2;
            p2 = p2.next;
            tmp.next = p1;
            l1 = tmp;
            p1 = l1;
        }
        while (p2 != null && p1.next != null) {
            ListNode tmp = p2;
            p2 = p2.next;
            while (p1.next != null && p1.next.val < tmp.val) p1 = p1.next;
            tmp.next = p1.next;
            p1.next = tmp;
            p1 = p1.next;
        }
        if (p2 != null) p1.next = p2;
        return l1;
    }
}