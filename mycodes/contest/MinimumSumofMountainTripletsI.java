class Solution {
    public int minimumSum(int[] nums) {
      int minLeft, minRight;
      int sum = -1;
      for (int i = 1; i < nums.length - 1; i++) {
          minLeft = getMinLeft(nums, i);
          minRight = getMinRight(nums, i);
          if (minLeft >= nums[i] || minRight >= nums[i]) continue;
          sum = sum == -1 ? minLeft + minRight + nums[i] : Math.min(sum, minLeft + minRight + nums[i]);
      }
      return sum;
    }

    private int getMinLeft(int[] nums, int i) {
      int min = nums[0];
      for (int j = 1; j < i; j++)
        if (nums[j] < min) min = nums[j];
      return min;
    }

    private int getMinRight(int[] nums, int i) {
      int min = nums[nums.length - 1];
      for (int j = nums.length - 2; j > i; j--)
        if (nums[j] < min) min = nums[j];
      return min;
    }
}
