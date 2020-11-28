class Solution {
    public String reverseWords(String s) {
        if (s == null || s.length() == 0) return s;
        char[] str = s.toCharArray();
        int i = 0;
        for (int j = 0; j < s.length(); ++j) {
            if (str[j] == ' ') {
                swapWord(str, i, j - 1);
                i = j + 1;
            }
        }
        swapWord(str, i, s.length() - 1);
        return new String(str);
    }
    
    private void swapWord(char[] str, int start, int end) {
        if (str == null || str.length == 0 || start >= end) return;
        while (start < end) {
            str[start] ^= str[end];
            str[end] ^= str[start];
            str[start] ^= str[end];
            start++;
            end--;
        }
    }
}