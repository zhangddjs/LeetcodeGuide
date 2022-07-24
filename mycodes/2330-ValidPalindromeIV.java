class Solution {
    public boolean makePalindrome(String s) {
        int cnt = 0;
        for (int i = 0, j = s.length() - 1; i < j && cnt < 3; i++, j--) if (s.charAt(i) != s.charAt(j)) cnt++;
        return cnt < 3;
    }
}