//TLE Solution

func resultArray(nums []int) []int {
	arr1, arr2 := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	arr1Unsort, arr2Unsort := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	arr1, arr2 = append(arr1, nums[0]), append(arr2, nums[1])
	arr1Unsort, arr2Unsort = append(arr1Unsort, nums[0]), append(arr2Unsort, nums[1])
	for i := 2; i < len(nums); i++ {
		num := nums[i]
		idx1, idx2 := bsearch(arr1, num), bsearch(arr2, num)
		arr1, arr2, arr1Unsort, arr2Unsort = compareAndAppend(arr1, arr2, arr1Unsort, arr2Unsort, num, idx1, idx2)
	}
	return append(arr1Unsort, arr2Unsort...)
}

func bsearch(arr []int, num int) int {
	n := len(arr)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func compareAndAppend(arr1, arr2, arr1Unsort, arr2Unsort []int, num, idx1, idx2 int) ([]int, []int, []int, []int) {
	if idx1 > idx2 {
		arr1 = append(arr1[:idx1], append([]int{num}, arr1[idx1:]...)...)
		arr1Unsort = append(arr1Unsort, num)
	} else if idx1 < idx2 {
		arr2 = append(arr2[:idx2], append([]int{num}, arr2[idx2:]...)...)
		arr2Unsort = append(arr2Unsort, num)
	} else {
		if len(arr1) > len(arr2) {
			arr2 = append(arr2[:idx2], append([]int{num}, arr2[idx2:]...)...)
			arr2Unsort = append(arr2Unsort, num)
		} else {
			arr1 = append(arr1[:idx1], append([]int{num}, arr1[idx1:]...)...)
			arr1Unsort = append(arr1Unsort, num)
		}
	}
	return arr1, arr2, arr1Unsort, arr2Unsort
}

//----------------------------- AC Version
func resultArray(nums []int) []int {
	arr1, arr2 := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	arr1Unsort, arr2Unsort := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	arr1, arr2 = append(arr1, nums[0]), append(arr2, nums[1])
	arr1Unsort, arr2Unsort = append(arr1Unsort, nums[0]), append(arr2Unsort, nums[1])
	for i := 2; i < len(nums); i++ {
		num := nums[i]
		idx1, idx2 := bsearch(arr1, num), bsearch(arr2, num)
		arr1, arr2, arr1Unsort, arr2Unsort = compareAndAppend(arr1, arr2, arr1Unsort, arr2Unsort, num, idx1, idx2)
	}
	return append(arr1Unsort, arr2Unsort...)
}

func bsearch(arr []int, num int) int {
	n := len(arr)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func compareAndAppend(arr1, arr2, arr1Unsort, arr2Unsort []int, num, idx1, idx2 int) ([]int, []int, []int, []int) {
	if idx1 > idx2 {
		arr1 = efficientInsert(arr1, num, idx1)
		arr1Unsort = append(arr1Unsort, num)
	} else if idx1 < idx2 {
		arr2 = efficientInsert(arr2, num, idx2)
		arr2Unsort = append(arr2Unsort, num)
	} else {
		if len(arr1) > len(arr2) {
			arr2 = efficientInsert(arr2, num, idx2)
			arr2Unsort = append(arr2Unsort, num)
		} else {
			arr1 = efficientInsert(arr1, num, idx1)
			arr1Unsort = append(arr1Unsort, num)
		}
	}
	return arr1, arr2, arr1Unsort, arr2Unsort
}

func efficientInsert(arr []int, num, idx int) []int {
	arr = append(arr, 0)         // Step 1
	copy(arr[idx+1:], arr[idx:]) // Step 2
	arr[idx] = num               // Step 3
	return arr
}