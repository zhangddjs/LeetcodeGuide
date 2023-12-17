func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	numCnt := make(map[int]int, n)
	for _, row := range grid {
		for _, v := range row {
			numCnt[v]++
		}
	}
	a, b := 0, 0
	for i := 1; i <= n*n; i++ {
		if numCnt[i] == 0 {
			b = i
		} else if numCnt[i] == 2 {
			a = i
		}
	}
	return []int{a, b}
}