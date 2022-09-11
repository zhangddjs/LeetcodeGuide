class Solution {
    public int[] prevPermOpt1(int[] arr) {
        for (int i = arr.length - 2; i >= 0; i--){
            if (arr[i] > arr[i + 1]) {
                int maxIdx = i + 1;
                for (int j = i + 1; j < arr.length; j++) {
                    if (arr[j] > arr[maxIdx] && arr[j] < arr[i]) maxIdx = j;
                }
                swap(arr, i, maxIdx);
                break;
            }
        }
        return arr;
    }

    private void swap(int[] arr, int i, int j) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }
}