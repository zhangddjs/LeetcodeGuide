// 这是错误的思路， this is a wrong idea for greedy
func minSessions(tasks []int, sessionTime int) int {
    sort.Ints(tasks)
	remains := make([]int, len(tasks))
	need := 1
	for i := len(tasks) - 1; i >= 0; i-- {
		for j := 0; j < len(remains); j++ {
			if j >= need {
				need = j + 1
			}
			if remains[j] + tasks[i] <= sessionTime {
				remains[j] = remains[j] + tasks[i]
				break
			}
		}
	}
	return need
}