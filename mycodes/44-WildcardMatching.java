class Solution {
    int [][] dp;
    public boolean isMatch(String s, String p) {
        if (s == null || p == null) return false;
        dp = new int[s.length() + 1][p.length() + 1];
        return helper(s, p) == 1 ? true : false;
    }
    
    private int helper(String s, String p) {
        if (s.length() == 0 && p.length() == 0) return 1;
        if (p.length() == 0) return -1;
        if (dp[s.length()][p.length()] != 0) return dp[s.length()][p.length()];
        int res = -1;
        if (s.length() == 0) {
            if (p.charAt(0) == '*')
                res = helper(s, p.substring(1));
            else res = -1;
            dp[s.length()][p.length()] = res;
            return res;
        }
        if (s.charAt(0) == p.charAt(0) || p.charAt(0) == '?') {
            res = helper(s.substring(1), p.substring(1));
            dp[s.length()][p.length()] = res;
            return res;
        }
        if (p.charAt(0) == '*') {
            for (int i = 0; i <= s.length() && res != 1; ++i) {
                res = helper(s.substring(i), p.substring(1));
            }
        }
        dp[s.length()][p.length()] = res;
        return res;
    }
}