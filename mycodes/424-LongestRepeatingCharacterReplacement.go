func characterReplacement(s string, k int) int {
  res := k
  for c := 'A'; c <= 'Z'; c++ {
    res = max(res, longestByChar(s, byte(c), k))
  }
  return res
}

func longestByChar(s string, c byte, k int) int {
  res, j := k, 0
  for i := range s {
    if s[i] != c {
      k--
    }
    if k < 0 {
      if s[j] != c {
        k++
      }
      j++
    }
    res = max(res, i-j+1)
  }
  return res
}
