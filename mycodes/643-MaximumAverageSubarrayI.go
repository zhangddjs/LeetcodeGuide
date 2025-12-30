func findMaxAverage(nums []int, k int) float64 {
  total, res := 0, float64(0)
  for i := 0; i < k; i++ {
    total += nums[i]
  }
  res = float64(total) / float64(k)
  for i := 1; i <= len(nums)-k; i++ {
    total -= nums[i-1]
    total += nums[i+k-1]
    res = max(res, float64(total) / float64(k))
  }
  return res
}

