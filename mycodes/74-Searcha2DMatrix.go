func searchMatrix(matrix [][]int, target int) bool {
  l, h := 0, len(matrix)*len(matrix[0])-1
  for l <= h {
    m := l + (h-l)/2
    i, j := toIdx(m, len(matrix[0]))
    if matrix[i][j] == target {
      return true
    } else if matrix[i][j] > target {
      h = m-1
    } else {
      l = m+1
    }
  }
  return false
}

func toIdx(x, col int) (int, int) {
  return x/col, x%col
}
