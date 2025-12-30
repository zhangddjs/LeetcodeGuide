func minEatingSpeed(piles []int, h int) int {
  l, hi := 1, getMaxSpeed(piles)
  for l < hi {
    m := l + (hi-l)/2
    hour := totalHour(piles, m)
    if hour <= h {
      hi = m
    } else {
      l = m + 1
    }
  }
  return hi
}

func totalHour(piles []int, s int) int {
  res := 0
  for _, p := range piles {
    h := p / s
    if p % s > 0 {
      h++
    }
    res += h
  }
  return res
}

func getMaxSpeed(piles []int) int {
  res := 0
  for _, p := range piles {
    res = max(res, p)
  }
  return res
}
