package main

func spiralOrder(matrix [][]int) []int {
  res := make([]int, 0, len(matrix)*len(matrix[0]))
  halfwidth := min((len(matrix)+1)/2, (len(matrix[0])+1)/2)
  for i := 0; i < halfwidth; i++ {
    res = append(res, getOrderOfBoarder(i,i,matrix)...)
  }
  return res
}

func getOrderOfBoarder(starti, startj int, matrix [][]int) []int {
  endj := len(matrix[starti])-startj
  endi := len(matrix)-starti
  res := make([]int, 0)
  for j := startj; j < endj; j++ {
    res = append(res, matrix[starti][j])  
  }
  for i := starti+1; i < endi-1; i++ {
    res = append(res, matrix[i][endj-1])  
  }
  for j := endj-1; j >= startj && endi-1 != starti; j-- {
    res = append(res, matrix[endi-1][j])  
  }
  for i := endi-2; i > starti && endj-1 != startj; i-- {
    res = append(res, matrix[i][startj])  
  }
  return res
}
