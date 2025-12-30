func minSubArrayLen(target int, nums []int) int {
  minLen := len(nums) + 1
  l, r, sum := 0, 0, 0
  for r < len(nums) {
    sum += nums[r]
    for l <= r && sum >= target {
      minLen = min(minLen, r-l+1)
      sum -= nums[l]
      l++
    }
    r++
  }
  if minLen > len(nums) {
    return 0
  }
  return minLen
}
