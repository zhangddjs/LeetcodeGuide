class Solution {
    public int closetTarget(String[] words, String target, int startIndex) {
        int minCnt = 0, cnt = 0, i, j, n = words.length;
        for (i = startIndex, j = 0; j < n; i = (i + 1) % n, j++) {
            if (words[i].equals(target)) break;
            cnt++;
        }
        if (j == n) return -1;
        minCnt = cnt;
        cnt = 0;
        for (i = startIndex, j = 0; j < n; i = (i - 1 + n) % n, j++) {
            if (words[i].equals(target)) break;
            cnt++;
        }
        return Math.min(cnt, minCnt);
    }
}