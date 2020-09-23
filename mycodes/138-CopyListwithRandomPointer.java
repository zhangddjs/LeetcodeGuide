class Solution1 {
    public Node copyRandomList(Node head) {
        if (head == null) return null;
        Node head2 = new Node(head.val);
        Node  n1 = head.next, n2 = head2;
        while (n1 != null) {
            n2.next = new Node(n1.val);
            n1 = n1.next;
            n2 = n2.next;
        }
        n1 = head;
        n2 = head2;
        Node p1 = head, p2 = head2;
        while (n1 != null) {
            Node random = n1.random;
            while (p1 != random) {
                p1 = p1.next;
                p2 = p2.next;
            }
            n2.random = p2;
            n1 = n1.next;
            n2 = n2.next;
            p1 = head;
            p2 = head2;
        }
        return head2;
    }
}

class Solution2 {
    public Node copyRandomList(Node head) {
        Node dummy = new Node(0);
        Node  n1 = head, n2 = dummy;
        Map<Node, Node> map = new HashMap<>();
        while (n1 != null) {
            if (n1.random != null) map.put(n1.random, new Node(n1.random.val));
            n1 = n1.next;
        }
        n1 = head;
        while (n1 != null) {
            n2.next = map.getOrDefault(n1, new Node(n1.val));
            n2 = n2.next;
            n2.random = map.get(n1.random);
            n1 = n1.next;
        }
        return dummy.next;
    }
}

class Solution3 {
    public Node copyRandomList(Node head) {
        if (head == null) return null;
        Node dummy = new Node(0);
        Node  n1 = head, n2 = dummy;
        while (n1 != null) {
            Node tmp = n1.next;
            n2.next = n1;
            n2.next.next = new Node(n1.val);
            n1 = tmp;
            n2 = n2.next.next;
        }
        n1 = head;
        n2 = n1.next;
        while (n1 != null) {
            n2.random = n1.random == null ? null : n1.random.next;
            n2 = n2.next == null ? dummy : n2.next.next;
            n1 = n1.next.next;
        }
        n1 = head;
        while (n2.next != null) {
            n2.next = n2.next.next;
            n2 = n2.next;
            n1.next = n2.next;
            n1 = n1.next;
        }
        return dummy.next;
    }
}