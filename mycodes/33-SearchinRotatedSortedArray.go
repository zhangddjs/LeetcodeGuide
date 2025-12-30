func search(nums []int, target int) int {
  offset := getOffset(nums)
  l, h := offset, offset+len(nums)-1
  for l <= h {
    m := l + (h-l)/2
    mIdx := m % len(nums)
    if nums[mIdx] == target {
      return mIdx
    } else if nums[mIdx] > target {
      h = m - 1
    } else {
      l = m + 1
    }
  }
  return -1
}

func getOffset(nums []int) int {
  l, h := 0, len(nums)-1
  for l < h {
    m := l + (h-l)/2
    if nums[m] < nums[h] {
      h = m
    } else {
      l = m + 1
    }
  }
  return h
}

