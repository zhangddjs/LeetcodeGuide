//Time 50%
//Space 75%
class Solution {
    public int singleNumber(int[] nums) {
        int res = 0, cntZero = 0, cntOne = 0;
        for (int i = 0; i < 32; ++i) {
            cntZero = 0;
            cntOne = 0;
            for (int num : nums) {
                if ((num & (1 << i)) == 0) cntZero++;
                else cntOne++;
            }
            if (cntZero % 3 == 0) res ^= (1 << i);
        }
        return res;
    }
}