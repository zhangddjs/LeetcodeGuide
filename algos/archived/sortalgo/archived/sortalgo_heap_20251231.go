package sortalgo

func heapsort(arr []int) {
	buildheap(arr)
	n := len(arr)
	for j := n - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		heapify(arr, 0, j)
	}
}

func buildheap(arr []int) {
	n := len(arr)
	for i := parent(n - 1); i >= 0; i-- {
		heapify(arr, i, n)
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
		arr[maxIdx], arr[i] = arr[i], arr[maxIdx]
		i = maxIdx
		l, r = children(i)
	}
}

func parent(n int) int {
	return (n - 1) / 2
}

func children(n int) (int, int) {
	return 2*n + 1, 2*n + 2
}
