func minOperations(k int) int {
	a := int(math.Sqrt(float64(k)))
	if a*a == k {
		return a - 1 + a - 1
	}
	res := a - 1
	for b, i := a, 1; b*i < k; i++ {
		res++
	}
	return res
}