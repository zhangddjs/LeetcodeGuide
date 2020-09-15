class Solution {
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        if (k == 0) return 0;
        int res = 0, i = 0, j = 0, product = 1;
        for (j = i; j < nums.length; ++j) {
            product *= nums[j];
            while (product >= k && i <= j) product /= nums[i++];
            res += (j - i + 1);
        }
        return res;
    }
}