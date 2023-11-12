package test

func maximumStrongPairXor(nums []int) int {
	maxXor := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= min(nums[i], nums[j]) {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
