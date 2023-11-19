
// WA
/*
[6,2,3,4,8,5,2,7,9]
[[0,1],[1,2],[2,3],[0,3],[2,4],[3,4],[2,2],[5,6]]
[5,3,8,2,6,1,4,6]
[[0,7],[3,5],[5,2],[3,0],[1,6]]
*/

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	cache := make([]int, n)
	i, j, k, is := 0, 1, 0, make([]int, 0, n)
	for ; j < n; j++ {
		if heights[j] > heights[j-1] {
			for k = j - 1; k >= i && heights[k] < heights[j]; k-- {
				cache[k] = j
			}
			for i2 := len(is) - 1; i2 >= 0 && heights[is[i2]] < heights[j]; i2-- {
				cache[is[i2]] = j
			}
			for ; i <= k; i++ {
				if cache[i] == 0 {
					is = append(is, i)
				}
			}
			i = j
		}
	}
	ans := make([]int, 0, len(queries))
	for _, q := range queries {
		if q[0] == q[1] || heights[q[0]] < heights[q[1]] {
			ans = append(ans, q[1])
			continue
		} else if cache[q[0]] == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, q[0])
		}
	}
	return ans
}