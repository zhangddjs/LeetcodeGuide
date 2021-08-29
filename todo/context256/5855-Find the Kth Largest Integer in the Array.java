class Solution {
    public String kthLargestNumber(String[] nums, int k) {
        PriorityQueue<String> pq = new PriorityQueue<>(k, (a, b) -> compareStr(a, b));
        for (int i = 0; i < nums.length; ++i) {
            if (pq.size() < k) {
                pq.offer(nums[i]);
            } else if (compareStr(pq.peek(), nums[i]) < 0) {
                pq.poll();
                pq.offer(nums[i]);
            }
        }
        return pq.peek();
    }
    
    int compareStr(String a, String b) {
        if (a.length() < b.length()) {
            return -1;
        } else if (a.length() > b.length()) {
            return 1;
        }
        for (int i = 0; i < a.length(); ++i) {
            if (a.charAt(i) < b.charAt(i)) {
                return -1;
            } else if (a.charAt(i) > b.charAt(i)) {
                return 1;
            }
        }
        return 1;
    }
}