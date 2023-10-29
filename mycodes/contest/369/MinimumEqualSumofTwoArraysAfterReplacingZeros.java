class Solution {
    class ArrayInfo {
        long sum;
        int numOfZero;
        ArrayInfo(long sum, int numOfZero) {
            this.sum = sum;
            this.numOfZero = numOfZero;
        }
    }

    public long minSum(int[] nums1, int[] nums2) {
        ArrayInfo info1 = getInfo(nums1);
        ArrayInfo info2 = getInfo(nums2);
        ArrayInfo bigger = info1.sum > info2.sum ? info1 : info2;
        ArrayInfo smaller = bigger == info1 ? info2 : info1;
        long diff = bigger.sum + bigger.numOfZero - smaller.sum - smaller.numOfZero;
        if (diff == 0) return bigger.sum + bigger.numOfZero;
        if ((diff < 0 && bigger.numOfZero == 0) || smaller.numOfZero == 0) {
            return -1;
        }
        if (diff <= 0) {
            return smaller.sum + smaller.numOfZero;
        } else {
            return bigger.sum + bigger.numOfZero;
        }
    }

    private ArrayInfo getInfo(int[] nums) {
        long sum = 0;
        int numOfZero = 0;
        for (int i : nums) {
            sum += i;
            if (i == 0) {
                numOfZero++;
            }
        }

        return new ArrayInfo(sum, numOfZero);
    }
}