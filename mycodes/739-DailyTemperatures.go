package main

// Input: temperatures = [73,74,75,71,69,72,76,73]
// Output: [1,1,4,2,1,1,0,0]

func dailyTemperatures(temperatures []int) []int {
  awnser := make([]int, len(temperatures))
  stack := make([][]int, 0, len(temperatures))
  for i, t := range temperatures {
    n := len(stack)
    j := n - 1
    for j >= 0 && stack[j][0] < t {
      idx := stack[j][1]
      awnser[idx] = i - idx
      j--
    }
    stack = stack[:j+1]
    stack = append(stack, []int{t, i})
  }
  return awnser
}
