class Solution {
    public int lengthOfLIS(int[] nums, int k) {
        int[] maxLen = new int[nums.length];
        maxLen[nums.length - 1] = 1;
        int res = 1;
        for (int i = nums.length - 2; i >= 0; i--) {
            int max = 0;
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] > nums[i] && nums[i] + k >= nums[j]) {
                    max = Math.max(max, maxLen[j]);
                }
            }
            maxLen[i] = max + 1;
            res = Math.max(res, max + 1);
        }
        return res;
    }
}