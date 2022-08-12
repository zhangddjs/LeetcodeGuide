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
    public ListNode removeZeroSumSublists(ListNode head) {
        Map<ListNode, Integer> nodeSum = new HashMap<>();
        Map<Integer, ListNode> sumNode = new HashMap<>();
        ListNode n = head;
        int sum = 0;
        while (n != null) {
            sum += n.val;
            if (sum != 0 && sumNode.containsKey(sum)) {
                for (ListNode p = sumNode.get(sum).next; p != n; p = p.next)
                    sumNode.remove(nodeSum.remove(p));
                sumNode.get(sum).next = n.next;
            } else if (sum == 0) {
                for (ListNode p = head; p != n; p = p.next)
                    sumNode.remove(nodeSum.remove(p));
                head = n.next;
            } else {
                nodeSum.put(n, sum);
                sumNode.put(sum, n);
            }
            n = n.next;
        }
        return head;
    }
}