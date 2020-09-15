class Solution {
    public int[] twoSum(int[] nums, int target) {
        for (int i = 0; i < nums.length - 1; ++i)
            for (int j = i + 1; j < nums.length; ++j)
                if (nums[i] + nums[j] == target) return new int[]{i, j};
        return null;
    }
}

class Solution2 {
    public int[] twoSum(int[] nums, int target) {
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < nums.length - 1; ++i) {
            if (set.contains(nums[i])) continue;
            set.add(nums[i]);
            for (int j = i + 1; j < nums.length; ++j)
                if (nums[i] + nums[j] == target) return new int[]{i, j};
        }
        return null;
    }
}

class Solution3 {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; ++i) {
            if (map.keySet().contains(nums[i])) return new int[]{map.get(nums[i]), i};
            map.put(target - nums[i], i);
        }
        return null;
    }
}