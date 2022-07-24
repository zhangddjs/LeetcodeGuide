class Solution {
    public int largestAltitude(int[] gain) {
        int res = 0, cur = 0;
        for (int i : gain) {
            cur += i;
            res = Math.max(cur, res);
        }
        return res;
    }
}