func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiag := float64(0)
	areaMap := make(map[float64]int)
	for _, d := range dimensions {
		diag := math.Sqrt(float64(d[0]*d[0] + d[1]*d[1]))
		area := d[0] * d[1]
		if maxDiag <= diag {
			maxDiag = diag
			if areaMap[diag] < area {
				areaMap[diag] = area
			}
		}
	}
	return areaMap[maxDiag]
}