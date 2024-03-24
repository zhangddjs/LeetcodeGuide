class Solution {
    public long[] mostFrequentIDs(int[] nums, int[] freq) {
        TreeMap<Long, Set<Integer>> freqMap = new TreeMap<>();
        Map<Integer, Long> numFreq = new HashMap<>();
        long[] res = new long[nums.length];
        for (int i = 0; i < nums.length; i++) {
            Long prevFreq = numFreq.getOrDefault(nums[i], new Long(0));
            Long curFreq = prevFreq + freq[i];
            numFreq.put(nums[i], curFreq);
            freqMap.getOrDefault(prevFreq, new HashSet<Integer>()).remove(nums[i]);
            if (freqMap.get(prevFreq) == null || freqMap.get(prevFreq).size() == 0) {
                freqMap.remove(prevFreq);
            }
            Set<Integer> freqSet = freqMap.getOrDefault(curFreq, new HashSet<Integer>());
            freqSet.add(nums[i]);
            freqMap.put(curFreq, freqSet);
            res[i] = freqMap.lastKey();
        }
        return res;
    }
}