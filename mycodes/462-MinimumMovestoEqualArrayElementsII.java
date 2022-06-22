class Solution {
    public int minMoves2(int[] nums) {
        int mid = getMid(nums);
        int res = 0;
        for (int num : nums) {
            res += Math.abs(num - mid);
        }
        return res;
    }
    
    public int getMid(int[] nums) {
        Queue<Integer> small = new PriorityQueue();
        Queue<Integer> large = new PriorityQueue(Collections.reverseOrder());
        for (int num : nums) {
            if (small.size() != 0 && num >= small.peek()) {
                small.offer(num);
            } else if(large.size() != 0 && num <= large.peek()) {
                large.offer(num);
            } else {
                small.offer(num);
            }
            if (small.size() - large.size() > 1) {
                large.offer(small.poll());
            } else if (large.size() - small.size() > 1) {
                small.offer(large.poll());
            }
        }
        return large.size() > small.size() ? large.peek() : small.peek();
    }
}