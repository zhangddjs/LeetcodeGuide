class Solution {
    public int commonFactors(int a, int b) {
        int m = Math.min(a, b), cnt = 0;
        for (int i = 1; i <= m; i++) if (a % i == 0 && b % i == 0) cnt++;
        return cnt;
    }
}