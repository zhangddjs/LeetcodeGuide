func addDigits(num int) int {
  res := num
  for res / 10 >= 1 {
    cur := 0
    for res > 0 {
      cur += res % 10
      res /= 10
    }
    res = cur
  }
  return res
}
