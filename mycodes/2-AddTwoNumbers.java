class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        if (l1 == null || l2 == null) return l1 == null ? l2 : l1;
        ListNode p1 = l1, p2 = l2, pre = l1;
        int carry = 0;
        while (p1 != null && p2 != null) {
            p1.val += p2.val + carry;
            carry = p1.val / 10;
            p1.val %= 10;
            pre = p1;
            p1 = p1.next;
            p2 = p2.next;
        }
        if (p1 == null) {
            pre.next = p2;
            p1 = p2;
        }
        while (carry != 0) {
            if (p1 == null) {
                pre.next = new ListNode(carry);
                break;
            } else {
                p1.val += carry;
                carry = p1.val / 10;
                p1.val %= 10;
                pre = p1;
                p1 = p1.next;
            }
        }
        return l1;
    }
}