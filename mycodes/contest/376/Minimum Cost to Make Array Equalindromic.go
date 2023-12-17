// WA use mean, need use median instead.
// https://stackoverflow.com/questions/35973950/find-a-number-for-minimum-sum-of-absolute-difference-in-an-array

func minimumCost(nums []int) int64 {
	sum := int64(0)
	for _, num := range nums {
		sum += int64(num)
	}
	avg := int64(math.Ceil(float64(sum) / float64(len(nums))))
	upPalin, downPalin := getRecentPalin(avg)
	costUp := computeCost(nums, upPalin)
	costDown := computeCost(nums, downPalin)
	if costUp < costDown {
		return costUp
	}
	return costDown
}

func getRecentPalin(avg int64) (int64, int64) {
	div := int64(10)
	dig := make([]int64, 0)
	for avg != 0 {
		dig = append(dig, avg%div)
		avg /= div
	}
	mid := (len(dig) + 1) / 2
	midIdx := mid - 1
	// up Palin
	upPalin := int64(0)
	for i, j := len(dig)-1, 0; i > j; i-- {
		dig[j] = dig[i]
		j++
	}
	for i := 0; i < len(dig); i++ {
		upPalin *= div
		upPalin += dig[i]
	}
	// down Palin
	if len(dig) == 2 && dig[len(dig)-1] == 1 {
		return upPalin, 9
	}
	for i := len(dig) - 1 - midIdx; i < len(dig); i++ {
		dig[i]--
		if dig[i] >= 0 {
			break
		}
		dig[i] = 9
	}
	downPalin := int64(0)
	for i, j := len(dig)-1, 0; i > j; i-- {
		dig[j] = dig[i]
		j++
	}
	for i := 0; i < len(dig); i++ {
		downPalin *= div
		downPalin += dig[i]
	}
	return upPalin, downPalin
}

func computeCost(nums []int, palin int64) int64 {
	cost := int64(0)
	for _, num := range nums {
		cost += abs(int64(num) - palin)
	}
	return cost
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}
