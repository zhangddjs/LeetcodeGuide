package sortalgo

// --------------------------------------------------------------

func heapsort(arr []int) {
	buildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func buildHeap(arr []int) {
	n := len(arr)
	for j := parent(n - 1); j >= 0; j-- {
		heapify(arr, j, n)
	}
}

func heapify(arr []int, i, j int) {
	l, r := children(i)
	for true {
		maxIdx := i
		if l < j && arr[l] > arr[maxIdx] {
			maxIdx = l
		}
		if r < j && arr[r] > arr[maxIdx] {
			maxIdx = r
		}
		if maxIdx == i {
			break
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		i = maxIdx
		l, r = children(i)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

// --------------------------------------------------------------

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
	i, j := l, l
	x := arr[r]
	for j < r {
		if arr[j] < x {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
		j++
	}
	arr[j], arr[i] = arr[i], arr[j]
	return i
}

// --------------------------------------------------------------

// arr [2,1,5,3]

func mergesort(arr []int) {
	mergesortarr(arr, 0, len(arr)-1)
}

func mergesortarr(arr []int, l, r int) {
	if l >= r { // l=0, r=3
		return
	}
	mid := l + (r-l)/2          // 1
	mergesortarr(arr, l, mid)   // sort (0,1)
	mergesortarr(arr, mid+1, r) // sort (2,3)
	merge(arr, l, mid, r)       // merge (arr 0, 2, 3)
}

func merge(arr []int, l, m, r int) {
	left, right := make([]int, m-l+1), make([]int, r-m)
	for i := 0; i < m-l+1; i++ {
		left[i] = arr[l+i]
	}
	for j := 0; j < r-m; j++ {
		right[j] = arr[m+1+j]
	}
	i, j, k := 0, 0, l
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// --------------------------------------------------------------

func insertsort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i - 1
		for j = i - 1; j >= 0 && arr[j] > x; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = x
	}
}

// --------------------------------------------------------------

func bubblesort(arr []int) {
	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}

// --------------------------------------------------------------

func toposort(arr []int) {

}
