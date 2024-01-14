// Still TE
func beautifulIndices(s string, a string, b string, k int) []int {
	aIdxs, bIdxs := getIndices(s, a), getIndices(s, b)
	res := make([]int, 0, len(aIdxs))
	if len(aIdxs) == 0 || len(bIdxs) == 0 {
		return res
	}
	i := 0
	for _, bi := range bIdxs {
		if i >= len(aIdxs) {
			break
		}
		leftbound, rightbound := leftBound(aIdxs, i, len(aIdxs), bi, k), rightBound(aIdxs, i, len(aIdxs), bi, k)
		res = append(res, aIdxs[leftbound:rightbound+1]...)
		i = rightbound + 1
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

func leftBound(idxs []int, left, right, target, k int) int {
	if len(idxs) == 0 || left < 0 || right > len(idxs) {
		return -1
	}
	if idxs[left] > target+k || idxs[right-1] < target-k {
		return left
	}
	for left < right {
		mid := (left + right) / 2
		if idxs[mid] >= target-k && idxs[mid] <= target+k {
			right = mid
		} else if idxs[mid] < target-k {
			left = mid + 1
		} else if idxs[mid] > target+k {
			right = mid
		}
	}
	return left
}

func rightBound(idxs []int, left, right, target, k int) int {
	if len(idxs) == 0 || left < 0 || right > len(idxs) {
		return -1
	}
	if idxs[left] > target+k || idxs[right-1] < target-k {
		return left - 1
	}
	for left < right {
		mid := (left + right) / 2
		if idxs[mid] >= target-k && idxs[mid] <= target+k {
			left = mid + 1
		} else if idxs[mid] < target-k {
			left = mid + 1
		} else if idxs[mid] > target+k {
			right = mid
		}
	}
	return left - 1
}