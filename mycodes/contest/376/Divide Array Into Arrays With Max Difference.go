func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}
		res = append(res, []int{nums[i], nums[i+1], nums[i+2]})
	}
	return res
}