func resultArray(nums []int) []int {
	arr1, arr2 := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])
	for i := 2; i < len(nums); i++ {
		if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
			arr1 = append(arr1, nums[i])
		} else {
			arr2 = append(arr2, nums[i])
		}
	}
	return append(arr1, arr2...)
}