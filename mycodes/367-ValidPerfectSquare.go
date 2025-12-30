func isPerfectSquare(num int) bool {
  l, h := 1, num
  for l <= h {
    m := l + (h-l)/2
    if m*m == num {
      return true
    } else if m*m > num {
      h = m-1
    } else {
      l = m+1
    }
  }
  return false
}
