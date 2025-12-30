func checkInclusion(s1 string, s2 string) bool {
  if len(s2) < len(s1) {
    return false
  }
  s1m := buildAlpCntMap(s1)
  s2m := buildAlpCntMap(s2[0:len(s1)])
  if judgeMapEqual(s1m, s2m) {
    return true
  }
  for i := 1; i <= len(s2)-len(s1); i++ {
    s2m[s2[i-1]]--
    if s2m[s2[i-1]] == 0 {
      delete(s2m, s2[i-1])
    }
    s2m[s2[i+len(s1)-1]]++
    if judgeMapEqual(s1m, s2m) {
      return true
    }
  }
  return false
}

func buildAlpCntMap(s string) map[byte]int {
  res := make(map[byte]int, 27)
  for i := range s {
    res[s[i]]++
  }
  return res
}

func judgeMapEqual(m1, m2 map[byte]int) bool {
  if len(m1) != len(m2) {
    return false
  }
  for k, v := range m1 {
    if m2[k] != v {
      return false
    }
  }
  return true
}
