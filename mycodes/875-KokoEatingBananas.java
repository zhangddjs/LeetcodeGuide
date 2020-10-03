//Time 86%
//Space 95%

class Solution {
    public int minEatingSpeed(int[] piles, int H) {
        //越慢越好->K越小越好
        //可以运用二分查找，low为1，high为piles中的最大值
        //对于每一个K，若消耗的时间大于H，则K小了，否则可能比要求的最小K大了
        //时间复杂度O(nlogm) m为香蕉最大数，n为串数。
        if (H < piles.length) return -1;
        int K, low = 1, high = getMax(piles);
        while (low < high) {
            K = low + (high - low) / 2;
            if (getEatTime(piles, K) > H) low = K + 1;
            else high = K;
        }
        return high;
    }
    
    public int getMax(int[] piles) {
        int max = Integer.MIN_VALUE;
        for (int v : piles) max = Math.max(max, v);
        return max;
    }
    
    public int getEatTime(int[] piles, int K) {
        int res = 0;
        for (int v : piles) res += (v - 1) / K + 1;
        return res;
    }
}