//Time 100%
//Space 30%

/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

class Solution {
    class Result{
        Node head, tail;
        public Result(Node head, Node tail) {
            this.head = head;
            this.tail = tail;
        }
    }
    
    public Node treeToDoublyList(Node root) {
        if (root == null) return null;
        Result res = convert(root);
        
        return res.head;
    }
    
    private Result convert(Node root) {
        if (root == null) return null;
        Result left = convert(root.left);
        Result right = convert(root.right);
        Result res = new Result(root, root);
        if (left != null) {
            left.tail.right = root;
            root.left = left.tail;
            res.head = left.head;
        }
        if (right != null) {
            right.head.left = root;
            root.right = right.head;
            res.tail = right.tail;
        }
        res.head.left = res.tail;
        res.tail.right = res.head;
        return res;
    }
}