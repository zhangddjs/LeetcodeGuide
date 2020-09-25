/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    public Node connect(Node root) {
        Node node = root;
        Node nextLevel = node == null ? null : node.left != null ? node.left : node.right;
        if (node != null && node.left != null && node.right != null) node.left.next = node.right;
        while (nextLevel != null) {
            node = nextLevel;
            nextLevel = null;
            Node pre = null;
            while (node != null) {
                if (nextLevel == null) nextLevel = node.left == null ? node.right : node.left;
                if (node.left != null && node.right != null) node.left.next = node.right;
                if (pre != null) pre.next = node.left == null ? node.right : node.left;
                else pre = node.right != null ? node.right : node.left;
                while (pre != null && pre.next != null) pre = pre.next;
                node = node.next;
            }
        }
        return root;
    }
}