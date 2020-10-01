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
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        Map<Integer, TreeNode> map = new HashMap<>();
        int p1 = 0, p2 = 0;
        TreeNode dummy = new TreeNode();
        TreeNode cur = dummy;
        while (p1 < preorder.length && p2 < inorder.length) {
            while (p2 < inorder.length && map.containsKey(inorder[p2])) p2++;
            if (p2 >= inorder.length) break;
            else if (p2 != 0) cur = map.get(inorder[p2 - 1]);
            cur.right = new TreeNode(preorder[p1]);
            cur = cur.right;
            map.put(preorder[p1++], cur);
            while (p1 < preorder.length && preorder[p1 - 1] != inorder[p2]) {
                cur.left = new TreeNode(preorder[p1]);
                cur = cur.left;
                map.put(preorder[p1++], cur);
            }
        }
        return dummy.right;
    }
}