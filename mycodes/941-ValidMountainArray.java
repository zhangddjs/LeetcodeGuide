class Solution {
    public boolean validMountainArray(int[] arr) {
        if (arr.length < 3 || arr[1] <= arr[0] || arr[arr.length - 1] >= arr[arr.length - 2]) return false;
        int i = 1;
        for (; i < arr.length && arr[i] > arr[i - 1]; i++);
        for (; i < arr.length && arr[i] < arr[i - 1]; i++);
        return i == arr.length;
    }
}