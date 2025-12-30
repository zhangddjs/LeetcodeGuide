// tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
// ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
// = ((10 * (6 / (12 * -11))) + 17) + 5
// = ((10 * (6 / -132)) + 17) + 5
// = ((10 * 0) + 17) + 5
// = (0 + 17) + 5
// = 17 + 5
// = 22
//
//

func evalRPN(tokens []string) int {
  stack := make([]int, 0, len(tokens))
  for _, t := range tokens {
    n := len(stack)
    res := 0
    if t == "+" {
      res = stack[n-1]+stack[n-2]
      stack = stack[:n-2]
    } else if t == "-" {
      res = stack[n-2]-stack[n-1]
      stack = stack[:n-2]
    } else if t == "*" {
      res = stack[n-1]*stack[n-2]
      stack = stack[:n-2]
    } else if t == "/" {
      res = stack[n-2]/stack[n-1]
      stack = stack[:n-2]
    } else {
      res,_ = strconv.Atoi(t)
    }
    stack = append(stack, res)
  }
  return stack[0]
}
