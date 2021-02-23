//100% Time
//17%

class Solution {
    public int reverse(int x) {
        int div = 10;
        int res = x % div;
        x /= div;
        while (x != 0) {
            if ((x > 0 && (0x7fffffff - x % div) / div < res) ||
                (x < 0 && (0x80000000 - x % div) / div > res) ) return 0;
            res = res * div + x % div;
            x /= div;
        }
        return res;
    }
}