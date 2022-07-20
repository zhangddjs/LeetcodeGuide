class Solution {
    public int maximumSum(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, List<Integer>> map = new HashMap<>();
        int max = -1, possibleMax = nums[nums.length - 1];
        for (int i = nums.length - 1; i >= 0; i--) {
            if (possibleMax + nums[i] < max) return max;
            int digitSum = computeDigitSum(nums[i]);
            if (map.get(digitSum) == null) map.put(digitSum, new ArrayList<>(Arrays.asList(nums[i])));
            else if (map.get(digitSum).size() == 1) {
                map.get(digitSum).add(nums[i]);
                max = Math.max(max, map.get(digitSum).get(0) + map.get(digitSum).get(1));
                if (possibleMax == map.get(digitSum).get(0)) possibleMax = computePossibleMax(map);
            }
        }
        return max;
    }

    private int computeDigitSum(int num) {
        int i = 0;
        while (num != 0) {
            i += num % 10;
            num /= 10;
        }
        return i;
    }

    private int computePossibleMax(Map<Integer, List<Integer>> map) {
        int possible = -1;
        for (List<Integer> nums : map.values()) {
            if (nums.size() == 1) {
                possible = Math.max(possible, nums.get(0));
            }
        }
        return possible;
    }
}