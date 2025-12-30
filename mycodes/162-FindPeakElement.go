func findPeakElement(nums []int) int {
  n := len(nums)
  if n == 1 || nums[0] > nums[1] {
    return 0
  }
  if nums[n-1] > nums[n-2] {
    return n-1
  }
  l, h := 1, n-2
  for l < h {
    m := l + (h-l)/2
    if nums[m] > nums[m-1] && nums[m] > nums[m+1] {
      return m
    }
    if nums[m] > nums[m-1] {
      l = m+1
    } else {
      h = m-1
    }
  }
  return h
}
