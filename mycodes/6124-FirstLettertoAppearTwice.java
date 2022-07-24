class Solution {
    public char repeatedCharacter(String s) {
        boolean[] exist = new boolean[26];
        for (char c : s.toCharArray()) {
            if (!exist[c - 'a']) exist[c - 'a'] = true;
            else return c;
        }
        return ' ';
    }
}