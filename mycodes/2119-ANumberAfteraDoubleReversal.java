class Solution {
    public boolean isSameAfterReversals(int num) {
        return num == reverse(reverse(num));
    }

    private int reverse(int num) {
        int res = 0;
        while (num != 0) {
            res = res * 10 + num % 10;
            num /= 10;
        }
        return res;
    }
}