class Solution {
    public int search(ArrayReader reader, int target) {
        int low = 0, high = 9999;
        while (low < high) {
            int mid = low + (high - low) / 2;
            if (reader.get(mid) == Integer.MAX_VALUE) high = mid - 1;
            else low = mid + 1;
        }
        low = 0;
        while (low <= high) {
            int mid = low + (high - low) / 2;
            if (reader.get(mid) == target) return mid;
            if (reader.get(mid) < target) low = mid + 1;
            else high = mid - 1;
        }
        return -1;
    }
}