class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null) return head;
        Stack<ListNode> stack = new Stack<>();
        while (head != null) {
            stack.push(head);
            head = head.next;
        }
        ListNode res = stack.pop();
        ListNode tail = res;
        while (!stack.isEmpty()) {
            tail.next = stack.pop();
            tail = tail.next;
        }
        tail.next = null;
        return res;
    }
}

class Solution2 {
    public ListNode reverseList(ListNode head) {
        ListNode tail = head, pre = head, res = head;
        while (tail != null && tail.next != null) {
            res = tail.next;
            tail.next = res.next;
            res.next = pre;
            pre = res;
        }
        return res;
    }
}

class Solution3 {
    ListNode res;
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) return head;
        reverse(head).next = null;
        return res;
    }

    public ListNode reverse(ListNode head) {
        if (head.next == null) {
            res = head;
            return head;
        }
        ListNode tail = reverse(head.next);
        tail.next = head;
        tail = tail.next;
        return tail;
    }
}