func areSimilar(mat [][]int, k int) bool {
	rows, cols := len(mat), len(mat[0])
	k = k % cols
	if k == 0 {
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] != mat[i][(j+k)%cols] {
				return false
			}
		}
	}
	return true
}