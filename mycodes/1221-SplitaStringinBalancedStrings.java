class Solution {
    public int balancedStringSplit(String s) {
        int res = 0, cntL = 0;
        for (int i = 0; i < s.length(); i++) {
            cntL += s.charAt(i) == 'L' ? 1 : -1;
            res += cntL == 0 ? 1 : 0;
        }
        return res;
    }
}