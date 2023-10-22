class Solution {
  public int minimumSum(int[] nums) {
      int[] leftMin = new int[nums.length];
      int[] rightMin = new int[nums.length];
      leftMin[0] = nums[0];
      rightMin[nums.length - 1] = nums[nums.length - 1];
      int sum = -1;
      for (int i = 1; i < nums.length; i++)
          if (nums[i] < leftMin[i - 1]) leftMin[i] = nums[i];
          else leftMin[i] = leftMin[i - 1];
      for (int i = nums.length - 2; i >= 0; i--)
          if (nums[i] < rightMin[i + 1]) rightMin[i] = nums[i];
          else rightMin[i] = rightMin[i + 1];
      
      for (int i = 1; i < nums.length - 1; i++) {
          int minLeft = leftMin[i - 1];
          int minRight = rightMin[i + 1];
          if (minLeft >= nums[i] || minRight >= nums[i]) continue;
          sum = sum == -1 ? minLeft + minRight + nums[i] : Math.min(sum, minLeft + minRight + nums[i]);
      }
      return sum;
  }
}
