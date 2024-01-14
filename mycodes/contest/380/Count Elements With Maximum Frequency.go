func maxFrequencyElements(nums []int) int {
	numsMap := make(map[int]int)
	freqMap := make(map[int]int)
	maxFreq := 0
	for _, num := range nums {
		numsMap[num]++
	}
	for _, freq := range numsMap {
		freqMap[freq]++
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return freqMap[maxFreq] * maxFreq
}