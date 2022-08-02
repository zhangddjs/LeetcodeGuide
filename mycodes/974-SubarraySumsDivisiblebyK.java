// TLE Brute Force
class Solution {
    public int subarraysDivByK(int[] nums, int k) {
        int cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            int sum = nums[i];
            cnt += sum % k == 0 ? 1 : 0;
            for (int j = i + 1; j < nums.length; j++) {
                sum += nums[j];
                cnt += sum % k == 0 ? 1 : 0;
            }
        }
        return cnt;
    }
}
