package test

func minOperations(nums1 []int, nums2 []int) int {
	n := len(nums1)
	opt1 := compute(nums1, nums2)
	if opt1 == -1 {
		return -1
	}
	nums1[n-1], nums2[n-1] = nums2[n-1], nums1[n-1]
	opt2 := compute(nums1, nums2)
	opt2++
	if opt2 < opt1 {
		return opt2
	}
	return opt1
}

func compute(nums1, nums2 []int) int {
	opt := 0
	n := len(nums1)
	for i := 0; i < n-1; i++ {
		if nums1[i] > nums1[n-1] && nums2[i] > nums1[n-1] {
			return -1
		}
		if nums1[i] > nums2[n-1] && nums2[i] > nums2[n-1] {
			return -1
		}
		if nums1[i] > nums1[n-1] && nums1[i] > nums2[n-1] {
			return -1
		}
		if nums2[i] > nums1[n-1] && nums2[i] > nums2[n-1] {
			return -1
		}
		if nums1[i] > nums1[n-1] || nums2[i] > nums2[n-1] {
			opt++
		}
	}
	return opt
}
