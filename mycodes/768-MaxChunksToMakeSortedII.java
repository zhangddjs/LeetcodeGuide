//Time 100%
//Space 96%
class Solution {
    public int maxChunksToSorted(int[] arr) {
        //记录每个块最小最大值
        //如果后面的块的最小值大于前面的块的最大值，则可以分
        int res = 0;
        int [] partitions = new int[arr.length];
        partitions[0] = arr[0];
        for (int i = 1; i < arr.length; ++i) {
            int tmp = res;
            if (arr[i] >= partitions[res]) partitions[++res] = arr[i];
            else {
                while (res - 1 >= 0 && arr[i] < partitions[res - 1]) res--;
                partitions[res] = partitions[tmp];
            }
        }
        return res + 1;
    }
}