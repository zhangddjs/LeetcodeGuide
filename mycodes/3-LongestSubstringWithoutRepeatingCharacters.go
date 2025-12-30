func lengthOfLongestSubstring(s string) int {
  seen := make(map[byte]bool)
  res, j := 0, 0
  for i := range s {
    if seen[s[i]] {
      for s[j] != s[i] {
        seen[s[j]] = false
        j++
      }
      j++
    }
    seen[s[i]] = true
    res = max(res, i-j+1)
  }
  return res
}
