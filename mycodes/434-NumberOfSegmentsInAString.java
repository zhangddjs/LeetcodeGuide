/**
 * 
 * TIME: 100.00%
 * SPACE: 54.55%
 */

class Solution {
    public int countSegments(String s) {
        int cnt = 0;
        boolean flag = false;
        if(s.length() == 0) {
            return 0;
        }
        for (int i = 0; i < s.length(); i++) {
            if(s.charAt(i) == ' ') {
                if (flag) {
                    flag = false;
                    cnt++;
                }
            } else {
                flag = true;
            }
        }
        cnt += flag ? 1 : 0;
        return cnt;
    }
}