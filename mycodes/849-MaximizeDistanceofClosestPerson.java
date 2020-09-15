class Solution {
    public int maxDistToClosest(int[] seats) {
        int [] shortest = new int[seats.length];
        int prev = -seats.length;
        for (int i = 0; i < seats.length; ++i) {
            if (seats[i] == 0) shortest[i] = i - prev;
            else prev = i;
        }
        prev = 2 * seats.length;
        for (int i = seats.length - 1; i >= 0; --i) {
            if (seats[i] == 0) shortest[i] = Math.min(shortest[i], prev - i);
            else prev = i;
        }
        Arrays.sort(shortest);
        return shortest[shortest.length - 1];
    }
}

class Solution {
    public int maxDistToClosest(int[] seats) {
        int i = -seats.length, j, res = 0;
        for (j = 0; j < seats.length && seats[j] != 1; ++j);
        for (int k = 0; k < seats.length; ++k) {
            if (k != j) res = Math.max(Math.min(j - k, k - i), res);
            else {
                i = j;
                for (j = j + 1; j < seats.length && seats[j] != 1; ++j);
                j = j == seats.length ? 2 * j : j;
            }
        }
        return res;
    }
}