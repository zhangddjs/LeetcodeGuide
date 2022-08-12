// TLE
class Solution {
    public int totalStrength(int[] strength) {
        long res = 0;
        for (int i = 0; i < strength.length; i++) {
            for (int j = 0; j < strength.length - i; j++) {
                res += min(strength, j, j + i) * sum(strength, j, j + i);
            }
        }
        return (int)(res % (10e9 + 7));
    }

    private int min(int[] strength, int i, int j) {
        int min = Integer.MAX_VALUE;
        for (i = i; i <= j; i++) {
            min = Math.min(min, strength[i]);
        }
        return min;
    }

    private int sum(int[] strength, int i, int j) {
        int sum = 0;
        for (i = i; i <= j; i++) {
            sum += strength[i];
        }
        return sum;
    }
}


