class Solution {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 2) return s;
        int maxLen = 1;
        String res = s.substring(0, 1);
        for (int i = 0; i < s.length() - 1; ++i) {
            String tmp = getLongestPalindrome(s, i, i);
            if (s.charAt(i) == s.charAt(i + 1)) {
                String tmp2 = getLongestPalindrome(s, i, i + 1);
                tmp = tmp.length() > tmp2.length() ? tmp : tmp2;
            }
            if (tmp.length() > maxLen) {
                maxLen = tmp.length();
                res = tmp;
            }
        }
        return res;
    }

    public String getLongestPalindrome(String s, int p1, int p2) {
        while (--p1 >= 0 && ++p2 < s.length() && s.charAt(p1) == s.charAt(p2));
        if (p1 == -1) return s.substring(p1 + 1, p2 + 1);
        else return s.substring(p1 + 1, p2);
    }
}

class Solution2 {
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 2) return s;
        String s2 = new StringBuilder(s).reverse().toString();
        int [] dp = new int[s.length()];
        int maxLen = 1, p = 0;
        //init
        for (int i = 0; i < s.length(); ++i) {
            dp[i] = s.charAt(0) == s2.charAt(i) ? 1 : 0;
        }
        //dp
        for (int i = 1; i < s.length(); ++i) {
            for (int j = s2.length() - 1; j > 0; --j) {
                dp[j] = s.charAt(i) == s2.charAt(j) ? dp[j - 1] + 1 : 0;
                if (dp[j] > maxLen) {
                    String tmp = s2.substring(j - dp[j] + 1, j + 1);
                    if (tmp.equals(new StringBuilder(tmp).reverse().toString())){
                        maxLen = dp[j];
                        p = j;
                    }
                }
            }
            dp[0] = s2.charAt(0) == s.charAt(i) ? 1 : 0;
        }
        //build res
        return s2.substring(p - maxLen + 1, p + 1);
    }
}