package main

func maxProfit(prices []int) int {
  minprice := prices[0]
  res := 0
  for i := 1; i < len(prices); i++ {
    if prices[i] < minprice {
      minprice = prices[i]
    } else {
      res = max(res, prices[i]-minprice)
    }
  }
  return res
}
