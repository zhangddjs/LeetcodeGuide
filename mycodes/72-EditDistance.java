//Bug Free!!!
//Time 13%
//Space 10%

class Solution {
    int [][] dp;
    
    public int minDistance(String word1, String word2) {
        if (word1 == null) return word2 == null ? 0 : word2.length();
        if (word2 == null) return word1.length();
        dp = new int[word1.length() + 1][word2.length() + 1];
        for (int i = 0; i <= word1.length(); ++i) Arrays.fill(dp[i], -1);
        return helper(word1, word2);
    }
    
    private int helper(String word1, String word2) {
        int m = word1.length(), n = word2.length();
        if (m == 0) return n;
        if (n == 0) return m;
        if (dp[m][n] != -1) return dp[m][n];
        if (word1.charAt(0) == word2.charAt(0)) {
            dp[m][n] = helper(word1.substring(1), word2.substring(1));
            return dp[m][n];
        }
        dp[m][n] = 1 + min(helper(word1.substring(1), word2),                   //delete
                           helper(word1, word2.substring(1)),                   //insert
                           helper(word1.substring(1), word2.substring(1)));     //replace
        return dp[m][n];
    }
    
    private int min(int a, int b, int c) {
        return Math.min(a, Math.min(b, c));
    }
}