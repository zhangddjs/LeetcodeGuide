func maximumLengthSubstring(s string) int {
	maxLen := 0
	for i := range s {
		maxLen = max(maxLen, lenOfSubstring(s[i:]))
	}
	return maxLen
}

func lenOfSubstring(s string) int {
	occr := make(map[byte]int)
	len := 0
	for i := range s {
		if v, ok := occr[s[i]]; ok && v == 2 {
			return len
		}
		len += 1
		occr[s[i]]++
	}
	return len
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}