func findMinimumOperations(s1 string, s2 string, s3 string) int {
	b1, b2, b3 := []byte(s1), []byte(s2), []byte(s3)
	sameCnt := 0
	for i := 0; i < len(b1) && i < len(b2) && i < len(b3); i++ {
		if b1[i] == b2[i] && b2[i] == b3[i] {
			sameCnt++
		} else {
			break
		}
	}
	if sameCnt < 1 {
		return -1
	}
	return len(b1) + len(b2) + len(b3) - 3*sameCnt
}