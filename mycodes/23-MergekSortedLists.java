class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Integer min = 0;
        while (min != null) {
            min = null;
            ListNode tmp = null;
            int index = -1;
            for (int i = 0; i < lists.length; ++i) {
                if (lists[i] != null && (min == null || lists[i].val < min)) {
                    min = lists[i].val;
                    tmp = lists[i].next;
                    index = i;
                }
            }
            if (min != null) {
                tail.next = new ListNode(min);
                tail = tail.next;
                lists[index] = tmp;
            }
        }
        return res.next;
    }
}

class Solution2 {
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Queue<ListNode> queue = new PriorityQueue((a, b) -> ((ListNode)a).val - ((ListNode)b).val);
        for (ListNode list : lists) {
            ListNode node = list;
            while (node != null) {
                queue.offer(node);
                node = node.next;
            }
        }
        while (!queue.isEmpty()) {
            tail.next = queue.poll();
            tail = tail.next;
        }
        tail.next = null;
        return res.next;
    }
}

class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        ListNode res = new ListNode();
        ListNode tail = res;
        Queue<ListNode> queue = new PriorityQueue<ListNode>((a, b) -> a.val - b.val);
        Map<ListNode, Integer> map = new HashMap<>();
        for (int i = 0; i < lists.length; ++i)
            if (lists[i] != null) {
                queue.offer(lists[i]);
                map.put(lists[i], i);
            }
        while (!queue.isEmpty()) {
            ListNode min = queue.poll();
            int index = map.get(min);
            tail.next = min;
            tail = tail.next;
            lists[index] = lists[index].next;
            map.remove(min);
            if (lists[index] != null) {
                map.put(lists[index], index);
                queue.offer(lists[index]);
            }
        }
        tail.next = null;
        return res.next;
    }
}