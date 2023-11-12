package test

import "sort"

//WA
func maximumStrongPairXor(nums []int) int {
	maxXor := 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := len(nums) - 1; i > 0; i-- {
	inner:
		for j := i - 1; j >= 0; j-- {
			if nums[i]-nums[j] > nums[j] {
				for i > j && nums[i]-nums[j] > nums[j] {
					i--
					if nums[i+1]-nums[i] > nums[i] {
						maxXor = max(maxXor, nums[i]^nums[i+1])
					}
				}
				i++
				break inner
			} else {
				maxXor = max(maxXor, nums[i]^nums[j])
			}
		}
	}

	return maxXor
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
