func findMin(nums []int) int {
  l, h := 0, len(nums)-1
  if nums[h] > nums[0] {
    return nums[0]
  }
  for l < h {
    m := l + (h-l)/2
    if nums[m] < nums[0] {
      h = m
    } else {
      l = m + 1
    }
  }
  return nums[h]
}
