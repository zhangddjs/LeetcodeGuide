class Solution {
    public Node insert(Node head, int insertVal) {
        Node node = head;
        if (node == null) {      //#case 4
            head = new Node(insertVal);
            head.next = head;
        }
        else {
            while (node.next != head) {     //#case 2,3
                if ((node.val > node.next.val &&
                    (node.val <= insertVal ||
                     node.next.val >= insertVal)) ||
                    (node.val <= insertVal &&
                     node.next.val >= insertVal)) break;     //#case 0,1
                node = node.next;
            }
            node.next = new Node(insertVal, node.next);
        }
        return head;
    }
}