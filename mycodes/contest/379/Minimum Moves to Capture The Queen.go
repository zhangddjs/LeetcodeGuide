func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
	if a == e { // same col rook and queen
		if a == c && ((b < d && f > d) || (b > d && f < d)) { // bishop in middle
			return 2
		}
		return 1
	}
	if b == f { // same col rook and queen
		if b == d && ((a < c && e > c) || (a > c && e < c)) { // bishop in middle
			return 2
		}
		return 1
	}
	if abs(c, e) == abs(d, f) { // same diag bishop and queen
		if abs(a, c) == abs(b, d) && abs(a, e) == abs(b, f) && ((c < a && a < e) || (c > a && a > e)) { // rook in middle
			return 2
		}
		return 1
	}
	return 2
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}