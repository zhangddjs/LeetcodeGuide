/**
 * O(n) S(1)
 */
class Solution {
    public int minOperations(String[] logs) {
        int res = 0;
        for (String log : logs) {
            if (log.equals("./")) {
                continue;
            } else if (log.equals("../")) {
                res = res == 0 ? res : res - 1;
            } else {
                res++;
            }
        }
        return res;
    }
}