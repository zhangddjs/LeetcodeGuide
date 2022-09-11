class Solution {
    public int numberOfWays(int startPos, int endPos, int k) {
        if (endPos < startPos) {
            startPos = endPos + startPos;
            endPos = startPos - endPos;
            startPos -= endPos;
        }
        if ((endPos - startPos) % 2 != k % 2 || k < endPos - startPos) return 0;
        return (int)(C(k, (k - (endPos - startPos)) / 2) % (10e9+7));
    }

    private long C(int n, int k) {
        long a = 1, b = 1;
        if (k > n / 2) k = n - k;
        for (long i = 1; i <= k; i++) {
            a *= (n + 1 - i);
            b *= i;
        }
        return a / b;
    }
}