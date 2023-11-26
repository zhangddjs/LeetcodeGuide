func beautifulSubstrings(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	cnt := 0
	byt := []byte(s)
	vows := make([]int, len(byt))
	if vowels[byt[0]] {
		vows[0]++
	}
	for i := 1; i < len(byt); i++ {
		vows[i] = vows[i-1]
		if vowels[byt[i]] {
			vows[i]++
		}
		consi := i + 1 - vows[i]
		if vows[i] == consi && vows[i]*consi%k == 0 {
			cnt++
		}
	}

	for i := 1; i < len(byt); i++ {
		for j := i + 1; j < len(byt); j += 2 {
			vowsij := vows[j] - vows[i-1]
			consij := j - i + 1 - vowsij
			if vowsij == consij && vowsij*consij%k == 0 {
				cnt++
			}
		}
	}

	return cnt
}