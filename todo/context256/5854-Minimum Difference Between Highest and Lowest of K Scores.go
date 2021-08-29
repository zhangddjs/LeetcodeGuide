func minimumDifference(nums []int, k int) int {
	if k == 1 {
		return 0
	}

	min := nums[k - 1] - nums[0]
	sort.Ints(nums)
	for i := 1; i <= len(nums) - k; i++ {
		diff := nums[i + k - 1] - nums[i]
		if diff < min {
			min = diff
		}
	}
	return min
}