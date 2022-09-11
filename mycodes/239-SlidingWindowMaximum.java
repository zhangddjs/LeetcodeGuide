class Solution {
    public int[] maxSlidingWindow(int[] nums, int k) {
        Deque<QueueElem> q = new LinkedList<>();
        int[] res = new int[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            while (!q.isEmpty() && q.getLast().val < nums[i]) q.pollLast();
            if (!q.isEmpty() && q.getFirst().idx <= i - k) q.pollFirst();
            q.addLast(new QueueElem(i, nums[i]));
            if (i >= k - 1) res[i - k + 1] = q.getFirst().val;
        }
        return res;
    }

    class QueueElem {
        int idx;
        int val;
        public QueueElem(int idx, int val) {
            this.idx = idx;
            this.val = val;
        }
    }

}