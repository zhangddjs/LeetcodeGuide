package main
import "fmt"

func summaryRanges(nums []int) []string {
  res := make([]string, 0, len(nums))
  if len(nums) == 0 {
    return res
  }
  cur := fmt.Sprintf("%v", nums[0])
  curNum := nums[0]
  for i := 1; i < len(nums); i++ {
    if nums[i] - nums[i-1] > 1 {
      if nums[i-1] != curNum {
        res = append(res, fmt.Sprintf("%v->%v", cur, nums[i-1]))
      } else {
        res = append(res, cur)
      }
      cur = fmt.Sprintf("%v", nums[i])
      curNum = nums[i]
    }
  }
  if nums[len(nums)-1] != curNum {
    res = append(res, fmt.Sprintf("%v->%v", cur, nums[len(nums)-1]))
  } else {
    res = append(res, cur)
  }
  return res
}
