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
        return helper(root, null);
    }

    public Node helper(Node root, Node brother) {
        if (root == null) return null;
        root.next = brother;
        helper(root.left, root.right);
        helper(root.right, brother == null ? null : brother.left);
        return root;
    }
}

class Solution2 {
    public Node connect(Node root) {
        if (root == null) return root;
        Node nextFirst = root.left;
        Node node = root;
        while (nextFirst != null) {
            node.left.next = node.right;
            node.right.next = node.next == null ? null : node.next.left;
            node = node.next;
            if (node == null) {
                node = nextFirst;
                nextFirst = node.left;
            }
        }
        return root;
    }
}