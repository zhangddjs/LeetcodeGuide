class Solution {
    public boolean checkPowersOfThree(int n) {
        for (int i = (int) (Math.log(n) / Math.log(3)); i >= 0; i--) {
            if (Math.pow(3, i) <= n) n -= Math.pow(3, i);
        }
        return n == 0;
    }
}