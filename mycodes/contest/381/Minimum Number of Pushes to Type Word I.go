func minimumPushes(word string) int {
	byts := []byte(word)
	cnts := make([]int, 26)
	for i := 0; i < len(byts); i++ {
		cnts[byts[i]-byte('a')]++
	}
	sort.Slice(cnts, func(i, j int) bool {
		return cnts[j] < cnts[i]
	})
	distCnt := 0
	res := 0
	for _, n := range cnts {
		res += n * (distCnt/8 + 1)
		distCnt++
	}
	return res
}