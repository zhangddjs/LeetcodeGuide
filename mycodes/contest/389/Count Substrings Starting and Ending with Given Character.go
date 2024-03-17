func countSubstrings(s string, c byte) int64 {
	cnt := countChars(s, c)
	return factorial(cnt)
}

func countChars(s string, c byte) int64 {
	cnt := int64(0)
	for i := range s {
		if s[i] == c {
			cnt++
		}
	}
	return cnt
}

func factorial(n int64) int64 {
	res := int64(0)
	for n > 0 {
		res += n
		n--
	}
	return res
}