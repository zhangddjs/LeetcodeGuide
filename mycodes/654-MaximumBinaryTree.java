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
    class MaxWithIdx{
        int max;
        int idx;
        MaxWithIdx(int max, int idx) {
            this.max = max;
            this.idx = idx;
        }
    }

    public TreeNode constructMaximumBinaryTree(int[] nums) {
        return build(nums, 0, nums.length);
    }

    private TreeNode build(int[] nums, int left, int right) {
        if (left >= right) return null;
        MaxWithIdx maxWithIdx = getMax(nums, left, right);
        TreeNode root = new TreeNode(maxWithIdx.max);
        root.left = build(nums, left, maxWithIdx.idx);
        root.right = build(nums, maxWithIdx.idx + 1, right);
        return root;
    }

    private MaxWithIdx getMax(int[] nums, int left, int right) {
        int max = Integer.MIN_VALUE, idx = left;
        for (int i = left; i < right; i++) {
            if (nums[i] > max) {
                max = nums[i];
                idx = i;
            }
        }
        return new MaxWithIdx(max, idx);
    }
}