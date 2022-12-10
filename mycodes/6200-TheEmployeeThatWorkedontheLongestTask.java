class Solution {
    public int hardestWorker(int n, int[][] logs) {
        int res = logs[0][0], longestTime = logs[0][1];
        for (int i = 1; i < logs.length; i++) {
            int time = logs[i][1] - logs[i - 1][1];
            if (time > longestTime){
                res = logs[i][0];
                longestTime = time;
            } else if (time == longestTime) res = Math.min(res, logs[i][0]);
        }
        return res;
    }
}