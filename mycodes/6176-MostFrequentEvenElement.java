class Solution {
    public int mostFrequentEven(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            if (num % 2 == 0) map.put(num, map.getOrDefault(num, 0) + 1);
        }
        int res = -1, max = 0;
        for (Iterator<Map.Entry<Integer,Integer>> it = map.entrySet().iterator(); it.hasNext();) {
            Map.Entry<Integer,Integer> entry = it.next();
            if (entry.getValue() > max) {
                res = entry.getKey();
                max = entry.getValue();
            } else if (entry.getValue() == max) res = Math.min(res, entry.getKey());
        }
        return res;
    }
}