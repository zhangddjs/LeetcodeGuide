func longestOnes(nums []int, k int) int {
  res, i := k, -1
  q := make([]int, 0, len(nums))
  if k > 0 {
    q = append(q, -1)
  }
  for j := 0 ; j < len(nums); j++ {
    if nums[j] == 0 {
      q = append(q, j)
      if len(q) > k {
        i = q[0]
        q = q[1:]
      }
    }
    res = max(res, j-i)
  }
  return res
}

