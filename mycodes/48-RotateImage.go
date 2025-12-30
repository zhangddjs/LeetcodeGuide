package main

func rotate(matrix [][]int)  {
  halfi := (min(len(matrix), len(matrix[0]))+1)/2
  for i := 0; i < halfi; i++ {
    rotateSubMatrix(i, i, matrix)
  }
}

func rotateSubMatrix(starti, startj int, matrix [][]int) {
  endi, endj := len(matrix)-starti, len(matrix[0])-startj
  for j := startj; j < endj-1; j++ {
    distancej := j - startj
    tmp := matrix[starti][j]
    matrix[starti][j] = matrix[endi-1-distancej][startj]
    matrix[endi-1-distancej][startj] = matrix[endi-1][endj-1-distancej]
    matrix[endi-1][endj-1-distancej] = matrix[starti+distancej][endj-1]
    matrix[starti+distancej][endj-1] = tmp
  }
}
