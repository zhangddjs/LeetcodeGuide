func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	biggestArea := int64(0)
	for i := range bottomLeft {
		for j := range bottomLeft {
			if i == j {
				continue
			}
			a := Rect{Point{bottomLeft[i][0], bottomLeft[i][1]}, Point{topRight[i][0], topRight[i][1]}}
			b := Rect{Point{bottomLeft[j][0], bottomLeft[j][1]}, Point{topRight[j][0], topRight[j][1]}}
			area := a.ComputeIntersectionSquare(b)
			biggestArea = max64(biggestArea, area)
		}
	}
	return biggestArea
}

type Point struct {
	x int
	y int
}

func (p Point) IsSmallerThan(q Point) bool {
	return p.x <= q.x && p.y <= q.y
}

type Rect struct {
	BottomLeft Point
	TopRight   Point
}

func (a Rect) MaxSquare() int64 {
	width := min(a.TopRight.x-a.BottomLeft.x, a.TopRight.y-a.BottomLeft.y)
	return int64(width) * int64(width)
}

func (a Rect) IsValid() bool {
	return a.BottomLeft.IsSmallerThan(a.TopRight)
}

func (a Rect) ComputeIntersectionSquare(b Rect) int64 {
	intersRect := Rect{
		BottomLeft: Point{max(a.BottomLeft.x, b.BottomLeft.x), max(a.BottomLeft.y, b.BottomLeft.y)},
		TopRight:   Point{min(a.TopRight.x, b.TopRight.x), min(a.TopRight.y, b.TopRight.y)},
	}
	if !intersRect.IsValid() {
		return 0
	}
	return intersRect.MaxSquare()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}