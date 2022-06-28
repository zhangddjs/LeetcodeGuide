class Solution {
    public long minimumTime(int[] time, int totalTrips) {
        if (time.length == 0 || totalTrips == 0) return 0;
        int min = getMin(time);
        long low = min, high = (long)min * totalTrips;
        while (low < high) {
            long mid = (low + high) / 2;
            long trips = 0;
            for (int i = 0; i < time.length; i++) {
                trips += mid / time[i];
            }
            if (trips >= totalTrips) {
                high = mid;
            } else {
                low = mid + 1;
            }
        }
        return high;
    }
    private int getMin(int[] time) {
        int min = Integer.MAX_VALUE;
        for (int i : time) {
            min = Math.min(min, i);
        }
        return min;
    }
}