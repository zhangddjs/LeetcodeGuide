//TLE
class Solution {
    int max[];

    public int deleteString(String s) {
        int m = s.length();
        max = new int[m];
        max[m - 1] = 1;
        for (int j = m - 2; j >= 0; j--) {
            int cnt = 1;
            String sub = s.substring(j);
            m = s.length() - j;
            for (int len = 1; len <= m / 2; len++) {
                if (canDelete(sub, len)) {
                    cnt = Math.max(cnt, 1 + max[j + len]);
                }
            }
            max[j] = cnt;
        }
        return max[0];
    }
    
    private boolean canDelete(String s, int len) {
        int j = len;
        for (int k = 0; k < len; k++)
            if (s.charAt(k) != s.charAt(j + k)) return false;
        return true;
    }
}