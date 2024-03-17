func isSubstringPresent(s string) bool {
	p := make(map[string]bool)

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			return true
		}
		sub := string([]byte{s[i-1], s[i]})
		if p[reverse(sub)] {
			return true
		}
		p[sub] = true
	}
	return false
}

func reverse(s string) string {
	res := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}