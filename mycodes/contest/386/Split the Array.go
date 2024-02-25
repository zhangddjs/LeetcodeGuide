func isPossibleToSplit(nums []int) bool {
	numMap := make(map[int]int, len(nums))
	for i := range nums {
		numMap[nums[i]] = numMap[nums[i]] + 1
		if numMap[nums[i]] > 2 {
			return false
		}
	}
	return true
}