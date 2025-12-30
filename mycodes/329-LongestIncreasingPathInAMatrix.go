package main

func longestIncreasingPath(matrix [][]int) int {
  dp := make([][]int, len(matrix))
  visited := make([][]int, len(matrix))
  res := 0
  for i := range matrix {
    dp[i] = make([]int, len(matrix[i]))
    visited[i] = make([]int, len(matrix[i]))
  }
  for i := range dp {
    for j := range dp[i]{
      res = max(res, dfs(i, j, matrix, dp, visited))
    }
  }
  return res
}

func dfs(i, j int, matrix, dp, visited [][]int) int {
  if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[i]) {
    return 0
  }
  if dp[i][j] != 0 {
    return dp[i][j]
  }
  visited[i][j] = 1
  res := 1
  if i > 0 && matrix[i-1][j] > matrix[i][j] && visited[i-1][j] != 1 {
    res = max(res, 1+dfs(i-1,j,matrix,dp,visited))
  }
  if j > 0 && matrix[i][j-1] > matrix[i][j] && visited[i][j-1] != 1 {
    res = max(res, 1+dfs(i,j-1,matrix,dp,visited))
  }
  if i < len(matrix)-1 && matrix[i+1][j] > matrix[i][j] && visited[i+1][j] != 1 {
    res = max(res, 1+dfs(i+1,j,matrix,dp,visited))
  }
  if j < len(matrix[i])-1 && matrix[i][j+1] > matrix[i][j] && visited[i][j+1] != 1 {
    res = max(res, 1+dfs(i,j+1,matrix,dp,visited))
  }
  visited[i][j] = 0
  dp[i][j] = res
  return res
}
