class Solution {
    public TreeNode insertIntoBST(TreeNode root, int val) {
        if (root == null) return new TreeNode(val);
        TreeNode node = root, parent = null;
        while (node.val != val) {
            if (node.val > val && node.left == null) node.left = new TreeNode(val);
            else if (node.val < val && node.right == null) node.right = new TreeNode(val);
            node = node.val > val ? node.left : node.right;
        }
        return root;
    }
}