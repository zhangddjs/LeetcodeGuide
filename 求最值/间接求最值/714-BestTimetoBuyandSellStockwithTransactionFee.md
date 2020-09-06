# [#714 Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)

![Medium](/figures/Medium.svg)

## 关键词

求最值、求最优方案、统计、枚举、股票问题、状态转移、动态规划、规律、贪心法、三指针、滑动窗口、双指针

## 题目

Your are given an array of integers `prices`, for which the `i`-th element is the price of a given stock on day `i`; and a non-negative integer `fee` representing a transaction fee.

You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)

Return the maximum profit you can make.

## 简述

**输入：** 价格数组; 税

**输出：** 最大利润

**Notes：**

+ 0 < 价格数组 <= 50000
+ 0 < 价格 < 50000
+ 0 <= 税 < 50000

## 思路

本题考察求最值，按求最值的框架应该是遍历或枚举所有方案，用哨兵记录最大值。

如何遍历每个情况呢？分析得知，在每个节点有买入、卖出两个操作，当进行买入时，后面任意比当前价格高的节点都可以卖出，在每个节点卖出都会产生一个利润，再加上从卖出节点下一个节点位置开始的最大利润，通过回溯法便可取得全局最大值。(换言之就是对于每一个小段，都有很多种交易方案，并且存在一种交易方案使得该段利润最大，全部小段所得利润拼接后即可得到全局最大利润)

基于回溯思想之上，我们可以设计出动态规划的解法，分别用两个变量记录第`i`天持有股票或未持股时的最大利润，那么第`i + 1`天持有股票的最大利润由第`i`天未持股的最大利润推倒出，同时未持股的最大利润也可以由第`i`天持股的最大利润导出(**卖出/不买->未持股，买入/不卖->持股**)。遍历结束后返回未持股的最大利润即可。------方法1[$^{[1]}$](#refer-anchor-1)

我们知道，动态规划法在满足一定条件时可以优化为贪心法。经过进一步分析得知，买入位置一般在极小值处，卖出位置一般在极大值(贪心)，但由于有交易税存在，所以当极小值大于等于上一个极大值-税的时候(第一个极小值除外)，将不考虑买入，当下一个极大值小于等于上一极小值+税的时候，将不考虑卖出。在这样的条件下，连续出现两个买入事件时将选择价格更低的那个，连续出现两个卖出事件时将选择价格更高的那个，出现一个买入，一个卖出事件时将利润进行统计。

因此可以使用三指针(滑动窗口)的方法，左边界为暂定买入点，右边界不断右移，如果发现在卖出之前有更好的买入点就更新左边界，否则以当前暂定买入点右边的最大值且满足卖出条件的点作为暂定卖出点，在有暂定买入点和卖出点后再遇到满足买入条件的点时就进行卖出操作并更新左边界为当前点。------方法2

## 解决方案

### 方法1-动态规划法[$^{[1]}$](#refer-anchor-1)

遍历数组，用两个变量记录当前时刻继续持股和未持股的最大利润，根据状态转移方程更新，最后返回在最后时刻未持股最大利润。(关键词：状态转移)

时间复杂度：$O(n)$

空间复杂度：$O(1)$

``` java
class Solution {
    public int maxProfit(int[] prices, int fee) {
        int cash = 0, hold = -prices[0];
        for (int i = 1; i < prices.length; i++) {
            cash = Math.max(cash, hold + prices[i] - fee);
            hold = Math.max(hold, cash - prices[i]);
        }
        return cash;
    }
}
```

### 方法2-贪心法(多指针法)

遍历数组，用双指针记录满足条件的暂定买入点和卖出点，统计所有交易的利润总和。(关键词：规律、三指针、滑动窗口)

时间复杂度：$O(n)$ ---38%

空间复杂度：$O(1)$ ---32%

``` java
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
```

## 扩展

### 扩展方法-双指针贪心法[$^{[2]}$](#refer-anchor-2)

有大神在三指针的贪心法基础上继续进行了十分巧妙的优化，实现了非常简洁的双指针贪心法。(双指针)

``` python
class Solution:
    def maxProfit(self, prices, fee):
        """
        :type prices: List[int]
        :type fee: int
        :rtype: int
        """
        n = len(prices)
        if n < 2:
             return 0
        ans = 0
        minimum = prices[0]
        for i in range(1, n):
            if prices[i] < minimum:
                minimum = prices[i]
            elif prices[i] > minimum + fee:
                ans += prices[i] - fee - minimum
                minimum = prices[i] - fee
        return ans
```

## 参考

<div id="refer-anchor-1"></div>

+ [1] [Leetcode. 714-Solution](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/solution/)

<div id="refer-anchor-2"></div>

+ [2] [Leetcode. 714-Discuss](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/201603/Python.-Greedy-is-good.)
