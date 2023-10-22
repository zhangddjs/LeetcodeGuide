class Solution {
    public int takeCharacters(String s, int k) {
        if (k == 0) return 0;
        int [] cnts = new int[3];
        int res;
        int pnt;
        for (pnt = 0; pnt < s.length(); pnt++) {
            cnts[s.charAt(pnt) - 'a']++;
            if (isEnough(cnts, k)) break;
        }
        if (!isEnough(cnts, k)) return -1;
        res = pnt + 1;
        int[] canMinus = new int[]{cnts[0] - k, cnts[1] - k, cnts[2] - k};
        int pnt2;
        for (pnt2 = s.length() - 1; pnt2 >= 0; pnt2--) {
            canMinus[s.charAt(pnt2) - 'a']++;
            while (pnt >= 0 && canMinus[s.charAt(pnt) - 'a'] > 0) {
                canMinus[s.charAt(pnt) - 'a']--;
                pnt--;
            }
            res = Math.min(res, pnt + 1 + (s.length() - pnt2));
        }
        return res;
    }

    private boolean isEnough(int[] cnts, int k) {
        return cnts[0] >= k && cnts[1] >= k && cnts[2] >= k;
    }
}