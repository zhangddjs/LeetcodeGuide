class Solution {
    public int minSessions(int[] tasks, int sessionTime) {
        Arrays.sort(tasks);
        boolean[] done = new boolean[tasks.length];
        int cnt = 0, doneCnt = 0;
        while (doneCnt < tasks.length) {
            int cur = sessionTime;
            for (int i = tasks.length - 1; i >= 0; i--) {
                if (done[i]) continue;
                if (cur >= tasks[i]) {
                    done[i] = true;
                    cur -= tasks[i];
                    doneCnt++;
                }
            }
            cnt++;
        }
        return cnt;
    }
}