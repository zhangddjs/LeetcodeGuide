// OOM
class Solution {
    public int[] xorQueries(int[] arr, int[][] queries) {
        int[] answer = new int[queries.length];
        int[][] buf = new int[arr.length][arr.length];
        for (int i = 0; i < arr.length; i++) {
            int cur = 0;
            for (int j = i; j < arr.length; j++) {
                cur ^= arr[j];
                buf[i][j] = cur;
            }
        }
        for (int i = 0; i < queries.length; i++) {
            answer[i] = buf[queries[i][0]][queries[i][1]];
        }
        return answer;
    }
}

class Solution {
    public int[] xorQueries(int[] arr, int[][] queries) {
        int[] answer = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            answer[i] = compute(arr, queries[i][0], queries[i][1]);
        }
        return answer;
    }

    public int compute(int[] arr, int left, int right) {
        int res = 0;
        if (right < left || left < 0 || right >= arr.length) {
            return res;
        }
        for (int i = left; i <= right; i++) {
            res ^= arr[i];
        }
        return res;
    }
}