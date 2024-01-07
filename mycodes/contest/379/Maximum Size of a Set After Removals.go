// beat 100%
func maximumSetSize(nums1 []int, nums2 []int) int {
	map1, map2, res := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	n := len(nums1)
	for _, num := range nums1 {
		map1[num] = true
		res[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
		res[num] = true
	}
	if len(map1) <= n/2 && len(map2) <= n/2 {
		return len(res)
	}
	sameMap := make(map[int]bool)
	for k := range map1 {
		if map2[k] {
			sameMap[k] = true
		}
	}
	needReduce1, needReduce2 := len(map1)-n/2, len(map2)-n/2
	if needReduce1 < 0 {
		needReduce1 = 0
	}
	if needReduce2 < 0 {
		needReduce2 = 0
	}
	if needReduce1+needReduce2 <= len(sameMap) {
		return len(res)
	}
	return len(res) - (needReduce1 + needReduce2 - len(sameMap))
}