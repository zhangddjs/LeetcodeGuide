/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;
        Deque<TreeNode> queue = new LinkedList<TreeNode>();
        queue.offerLast(root);
        int size = queue.size(), level = 1;
        List<Integer> curLevel = new ArrayList<>();
        while (!queue.isEmpty()) {
            if (size == 0) {
                level++;
                size = queue.size();
                res.add(curLevel);
                curLevel = new ArrayList<>();
            }
            TreeNode node = level % 2 != 0 ? queue.pollFirst() : queue.pollLast();
            curLevel.add(node.val);
            if (level % 2 != 0) {
                if (node.left != null) queue.offerLast(node.left);
                if (node.right != null) queue.offerLast(node.right);
            } else {
                if (node.right != null) queue.offerFirst(node.right);
                if (node.left != null) queue.offerFirst(node.left);
            }
            size--;
        }
        res.add(curLevel);
        return res;
    }
}