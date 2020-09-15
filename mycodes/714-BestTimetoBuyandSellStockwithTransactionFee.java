class Solution {
    public int maxProfit(int[] prices, int fee) {
        if (prices.length <= 1) return 0;
        int minimal = prices[0], maximal = prices[0], profit = 0;
        boolean isSellReady = false;
        for (int i = 1; i < prices.length; ++i) {
            if (prices[i] < minimal && !isSellReady) {
                minimal = prices[i];
                maximal = minimal;
            } else if (prices[i] > minimal + fee && prices[i] > maximal) {
                isSellReady = true;
                maximal = prices[i];
            } else if (prices[i] + fee < maximal && isSellReady) {
                profit += maximal - minimal - fee;
                minimal = prices[i];
                maximal = minimal;
                isSellReady = false;
            }
        }
        if (isSellReady) profit += maximal - minimal - fee;
        return profit;
    }
}