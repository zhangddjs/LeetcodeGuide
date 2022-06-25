class Solution {
    public int longestPalindrome(String word1, String word2) {
        int len1 = word1.length(), len2 = word2.length();
        int starti = Integer.MAX_VALUE, startj = Integer.MIN_VALUE;
        for (int i = 0; i < len1; i++) {
            for (int j = 0; j < len2; j++) {
                if (word1.charAt(i) == word2.charAt(j)) {
                    starti = Math.min(starti, i);
                    startj = Math.max(startj, j);
                }
            }
        }
        if (starti == Integer.MAX_VALUE) return 0;

        word1 = word1.substring(starti) + word2.substring(0, startj + 1);
        word2 = new StringBuilder(word1).reverse().toString();
        int n = word1.length();
        int[][] buf = new int[n + 1][n + 1];
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                buf[i][j] = Math.max(buf[i - 1][j], buf[i][j - 1]);
                if (word1.charAt(i - 1) == word2.charAt(j - 1) && buf[i - 1][j - 1] == buf[i][j - 1])
                    buf[i][j] = buf[i - 1][j - 1] + 1;
            }
        }
        return buf[n][n];
    }
}