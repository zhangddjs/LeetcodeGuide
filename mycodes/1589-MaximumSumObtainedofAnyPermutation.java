// TLE Solution
class Solution {
    public int maxSumRangeQuery(int[] nums, int[][] requests) {
        Arrays.sort(nums);
        int[] idxCnt = new int[nums.length];
        long res = 0;
        for (int[] req : requests) {
            for (int i = req[0]; i <= req[1]; i++) {
                idxCnt[i]++;
            }
        }
        Queue<Integer> q = new PriorityQueue<Integer>(Comparator.reverseOrder());
        for (int i = 0; i < idxCnt.length; i++) {
            if (idxCnt[i] != 0) q.offer(idxCnt[i]);
        }
        int cur = nums.length - 1;
        while (!q.isEmpty()) {
            int cnt = q.poll();
            res += cnt * nums[cur--];
        }
        return (int) (res % (Math.pow(10, 9) + 7));
    }
}