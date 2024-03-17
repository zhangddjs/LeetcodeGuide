func minimumDeletions(word string, k int) int {
	freqMap := make(map[byte]int)
	for c := byte('a'); c <= 'z'; c++ {
		f := freq(word, c)
		if f != 0 {
			freqMap[c] = freq(word, c)
		}
	}
	return computeMinNeedDeleted(freqMap, k)
}

func computeMinNeedDeleted(freqMap map[byte]int, k int) int {
	if len(freqMap) == 1 {
		return 0
	}
	key := findMinFreqKey(freqMap)
	freq := freqMap[key]
	newMap := copyMap(freqMap)
	delete(newMap, key)
	cnt1 := computeNeedDeleted(freqMap, freq, k)
	cnt2 := computeMinNeedDeleted(newMap, k) + freq
	return min(cnt1, cnt2)
}

func findMinFreqKey(freqMap map[byte]int) byte {
	min := 100001
	res := byte('A')
	for k, v := range freqMap {
		if v <= min {
			min = v
			res = k
		}
	}
	return res
}

func computeNeedDeleted(freqMap map[byte]int, freq, k int) int {
	res := 0
	bar := freq + k
	for _, v := range freqMap {
		if v > bar {
			res += v - bar
		}
	}
	return res
}

func copyMap(freqMap map[byte]int) map[byte]int {
	res := make(map[byte]int)
	for k, v := range freqMap {
		res[k] = v
	}
	return res
}

func freq(word string, c byte) int {
	cnt := 0
	for i := range word {
		if word[i] == c {
			cnt++
		}
	}
	return cnt
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}