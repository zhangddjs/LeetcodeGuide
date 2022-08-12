class Solution {
    public int arithmeticTriplets(int[] nums, int diff) {
        int cnt = 0;
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = i + 1; j < nums.length; j++) {
                if (nums[j] - nums[i] < diff) continue;
                else if (nums[j] - nums[i] > diff) break;
                for (int k = j + 1; k < nums.length; k++) {
                    if (nums[k] - nums[j] < diff) continue;
                    else if (nums[k] - nums[j] > diff) break;
                    cnt++;
                }
            }
        }
        return cnt;
    }
}