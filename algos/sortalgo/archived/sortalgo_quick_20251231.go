package sortalgo

func quicksort(arr []int) {
	quicksortarr(arr, 0, len(arr)-1)
}

func quicksortarr(arr []int, l, r int) {
	if l < r {
		i := partition(arr, l, r)
		quicksortarr(arr, l, i-1)
		quicksortarr(arr, i+1, r)
	}
}

func partition(arr []int, l, r int) int {
	x := arr[r]
	i := l
	for j := l; j < r; j++ {
		if arr[j] < x {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
