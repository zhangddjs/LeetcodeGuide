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
    public List<List<Integer>> levelOrder(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<TreeNode>();
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) return res;
        queue.offer(root);
        int rest = queue.size();
        List<Integer> curLevel = new ArrayList<>();
        while (!queue.isEmpty()) {
            if (rest == 0) {
                rest = queue.size();
                res.add(curLevel);
                curLevel = new ArrayList<>();
            }
            TreeNode node = queue.poll();
            curLevel.add(node.val);
            rest--;
            if (node.left != null) queue.offer(node.left);
            if (node.right != null) queue.offer(node.right);
        }
        res.add(curLevel);
        return res;
    }
}