//TLE
class Solution {
    public long appealSum(String s) {
        long res = 0;
        for (int i = 0; i < s.length(); i++) {
            long cur = 1;
            res += cur;
            boolean[] exist = new boolean[26];
            exist[s.charAt(i) - 'a'] = true;
            for (int j = i + 1; j < s.length(); j++) {
                cur += exist[s.charAt(j) - 'a'] ? 0 : 1;
                exist[s.charAt(j) - 'a'] = true;
                res += cur;
            }
        }
        return res;
    }
}