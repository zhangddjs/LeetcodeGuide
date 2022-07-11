class Solution {
    public boolean canChange(String start, String target) {
        int n = start.length();
        int cntL, cntR, cntL2, cntR2;
        for (int i = 0; i < n; i++) {
            if (start.charAt[i] == 'L') cntL++;
            if (start.charAt(i) == 'R') cntR++;
            if (target.charAt(i) == 'L') cntL2++;
            if (target.charAt(i) == 'R') cntR2++;
        }
        if (cntL != cntL2 || cntR != cntR2) return false;
        int lastj = 0;
        for (int i = 0; i < n; i++) {
            char ch = start.charAt(i);
            if (ch == '_') continue;
            int j = lastj;
            for (; j < n; j++) {
                char cj = target.charAt(j);
                if (cj == '_') continue;
                if (cj == ch) {
                    if ((cj == 'R' && j < i) || (cj == 'L' && j > i)) return false;
                    lastj = j + 1;
                    break;
                }
                return false;
            }
            if (j == n) return false;
        }
        return true;
    }
}