func minimumBoxes(apple []int, capacity []int) int {
	sumOfApple := sumArr(apple)
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	cnt := 0
	for _, e := range capacity {
		sumOfApple -= e
		cnt++
		if sumOfApple <= 0 {
			break
		}
	}
	return cnt
}

func sumArr(arr []int) int {
	sum := 0
	for _, e := range arr {
		sum += e
	}
	return sum
}