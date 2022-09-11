class Solution {
    public boolean checkDistances(String s, int[] distance) {
        int[][] pos = new int[2][26];
        for (int i = 0; i < s.length(); i++) {
            int ch = s.charAt(i) - 'a';
            if (pos[0][ch] != 0) pos[1][ch] = i + 1;
            else pos[0][ch] = i + 1;
        }
        for (int i = 0; i < distance.length; i++) {
            if (pos[0][i] != 0 && distance[i] + 1 != pos[1][i] - pos[0][i]) return false;
        }
        return true;
    }
}