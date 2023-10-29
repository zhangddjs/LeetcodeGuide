class Solution {
    public int findKOr(int[] nums, int k) {
        int [] count = compute(nums);
        int res = 0;
        for (int i = 0; i < 32; i++) {
            if (count[i] >= k) res |= 1 << i;
        }
        return res;
    }

    private int[] compute(int[] nums) {
        int [] res = new int[32];
        for (int i = 0; i < nums.length; i++)
            for (int j = 0; j < 32; j++)
                if ((nums[i] >> j & 1) == 1) res[j]++;
        return res;
    }
}