class Solution {
    public void reverseWords(char[] s) {
        reverse(0, s.length - 1, s);
        int i = 0, j = -1;
        while (++j < s.length) {
            if (s[j] == ' '){
                reverse(i, j - 1, s);
                i = j + 1;
            }
        }
        reverse(i, j - 1, s);
    }

    public void reverse(int i, int j, char[] s) {
        for (i = i, j = j; i < j; ++i, --j) {
            char tmp = s[i];
            s[i] = s[j];
            s[j] = tmp;
        }
    }
}