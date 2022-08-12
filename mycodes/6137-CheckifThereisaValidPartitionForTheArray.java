class Solution {
    public boolean validPartition(int[] nums) {
        int a = 1, b = 1;
        for (int i = 1; i < nums.length; i++) {
            if (a == 0 && b == 0) {
                a++;
                b++;
                continue;
            }
            if (nums[i] == nums[i - 1]) {
                if (b == 2 && a == 1) return false;
                a++;
                b = a == 1 ? 0 : 1;
            } else if (nums[i] - nums[i - 1] == 1) {
                if (a == 2 && b == 2) return false;
                a = 0;
                b = (b + 1) % 3;
            } else {
                if (b == 2) return false;
                a = 1;
                b = 1
            }
        }
        return a > 1 || b == 0;
    }
}