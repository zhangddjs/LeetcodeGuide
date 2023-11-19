func minimumSteps(s string) int64 {
	b := []byte(s)
	cnt := int64(0)
	cnt0 := int64(0)
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == '1' {
			cnt += cnt0
			continue
		}
		cnt0++
	}
	return cnt
}