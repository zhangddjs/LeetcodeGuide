func beautifulIndices(s string, a string, b string, k int) []int {
	aIdxs, bIdxs := getIndices(s, a), getIndices(s, b)
	res := make([]int, 0, len(aIdxs))
	if len(aIdxs) == 0 || len(bIdxs) == 0 {
		return res
	}
	i := 0
	for _, bi := range bIdxs {
		for i < len(aIdxs) && bi-k > aIdxs[i] {
			i++
		}
		for i < len(aIdxs) && bi-k <= aIdxs[i] && bi+k >= aIdxs[i] {
			res = append(res, aIdxs[i])
			i++
		}
	}
	return res
}

func getIndices(s string, sub string) []int {
	idx := -1
	res := make([]int, 0)
	for {
		idx = indexAt(s, sub, idx+1)
		if idx == -1 {
			break
		}
		res = append(res, idx)
	}
	return res
}

func indexAt(s, sub string, n int) int {
	idx := strings.Index(s[n:], sub)
	if idx > -1 {
		idx += n
	}
	return idx
}