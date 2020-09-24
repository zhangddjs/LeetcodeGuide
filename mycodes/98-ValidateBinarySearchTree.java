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
    public boolean isValidBST(TreeNode root) {
        if (root == null) return true;
        boolean isValid = true;
        isValid &= root.left == null ? true : root.left.val < root.val;
        isValid &= root.right == null ? true : root.right.val > root.val;
        isValid = isValid && helper(root.left, null, root) && helper(root.right, root, null);
        return isValid;
    }

    public boolean helper(TreeNode root, TreeNode leftparent, TreeNode rightparent) {
        boolean isValid = true;
        if (root == null || (root.left == null && root.right == null)) return true;
        isValid &= root.left == null ? true : root.left.val < root.val;
        isValid &= root.right == null ? true : root.right.val > root.val;
        isValid &= root.right == null || rightparent == null ? true : root.right.val < rightparent.val;
        isValid &= root.left == null || leftparent == null ? true : root.left.val > leftparent.val;
        isValid = isValid && helper(root.left, leftparent, root) && helper(root.right, root, rightparent);
        return isValid;
    }
}

class Solution2 {
    long pre = Long.MIN_VALUE;
    public boolean isValidBST(TreeNode root) {
        if (root == null) return true;
        boolean isValid = isValidBST(root.left) & root.val > pre;
        pre = root.val;
        return isValid && isValidBST(root.right);
    }
}