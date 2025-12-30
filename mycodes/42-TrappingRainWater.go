package main

func trap(height []int) int {
  wall := make([]int, 0, len(height))
  res := 0
  maxHeight := 0
  for i := range height {
    h := min(maxHeight, height[i])
    for j := len(wall)-1; j >= 0; j-- {
      if wall[j] < h {
        res += h - wall[j]
        wall[j] = h
      } else {
        break
      }
    }
    wall = append(wall, height[i])
    maxHeight = max(maxHeight, height[i])
  }
  return res
}
