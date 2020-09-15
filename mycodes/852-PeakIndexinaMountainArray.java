class Solution {
    public int peakIndexInMountainArray(int[] A) {
        int [] temp = new int[2];
        for (int i = 0; i < A.length; ++i) {
            if (A[i] > temp[0]) {
                temp[0] = A[i];
                temp[1] = i;
            }
        }
        return temp[1];
    }
}

class Solution2 {
    public int peakIndexInMountainArray(int[] A) {
        for (int i = 1; i < A.length - 1; ++i)
            if (A[i] > A[i + 1]) return i;
        return -1;
    }
}

class Solution3 {
    public int peakIndexInMountainArray(int[] A) {
        int low = 1, high = A.length - 1, mid;
        while (low <= high) {
            mid = low + (high - low) / 2;
            if (A[mid] > A[mid - 1] && A[mid] > A[mid + 1]) return mid;
            if (A[mid] < A[mid + 1]) low = mid + 1;
            else high = mid - 1;
        }
        return -1;
    }
}