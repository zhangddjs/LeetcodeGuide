func topKFrequent(nums []int, k int) []int {
  freqMap := make(map[int]int, len(nums))
  for _, num := range nums {
    freqMap[num]++
  }
  freq := make([]Elem, 0, len(freqMap))
  for k, v := range freqMap {
    freq = append(freq, Elem{k,v})
  }
  sort.Slice(freq, func(a, b int) bool {
    return freq[a].Freq > freq[b].Freq
  })
  res := make([]int, 0, len(freq))
  for i := 0; i < k; i++ {
    res = append(res, freq[i].Val)
  }
  return res
}

type Elem struct {
  Val int
  Freq int
}
