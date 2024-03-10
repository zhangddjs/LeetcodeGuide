func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	sum := sumArr(happiness, k)
	return sum
}

func sumArr(arr []int, k int) int64 {
	sum := int64(0)
	for i := 0; i < k; i++ {
		sum += int64(max(arr[i]-i, 0))
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}