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
    Map<TreeNode, BSTNodeInfo> map;

    class BSTNodeInfo {
        long sum;
        int max, min;

        public BSTNodeInfo(long sum, int min, int max) {
            this.sum = sum;
            this.min = min;
            this.max = max;
        }
    }

    public int maxSumBST(TreeNode root) {
        map = new HashMap<>();
        buildBSTInfoMap(root);
        return computeMax();
    }

    private void buildBSTInfoMap(TreeNode root) {
        if (root == null) return;
        if (root.left == null && root.right == null) {
            map.put(root, new BSTNodeInfo((long) root.val, root.val, root.val));
            return;
        }
        buildBSTInfoMap(root.left);
        buildBSTInfoMap(root.right);
        if (isBSTRoot(root)) {
            map.put(root, buildBSTRootInfo(root));
        }
    }

    private BSTNodeInfo buildBSTRootInfo(TreeNode root) {
        BSTNodeInfo rootInfo = new BSTNodeInfo((long) root.val, root.val, root.val);
        if (root.left != null) {
            BSTNodeInfo leftInfo = map.get(root.left);
            rootInfo.sum += leftInfo.sum;
            rootInfo.min = Math.min(rootInfo.min, leftInfo.min);
        }
        if (root.right != null) {
            BSTNodeInfo rightInfo = map.get(root.right);
            rootInfo.sum += rightInfo.sum;
            rootInfo.max = Math.max(rootInfo.max, rightInfo.max);
        }
        return rootInfo;
    }

    private boolean isBSTRoot(TreeNode root) {
        return !((root.left != null && (!map.containsKey(root.left) || map.get(root.left).max >= root.val)) ||
                 (root.right != null && (!map.containsKey(root.right) || map.get(root.right).min <= root.val)));
    }

    private int computeMax() {
        long max = 0;
        for (TreeNode key : map.keySet()) max = Math.max(map.get(key).sum, max);
        return (int) max;
    }
}