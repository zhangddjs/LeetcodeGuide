package sortalgo

func heapsort(arr []int) {
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		swiftDown(arr, 0, i)
	}
}

// 不要用swiftUp，因为复杂度更大，而且必须从叶子节点开始
func heapify(arr []int) {
	n := len(arr)
	p := parent(n - 1)
	for p >= 0 {
		swiftDown(arr, p, n)
		p--
	}
}

func swiftUp(arr []int, i int) {
	p := parent(i)
	for arr[p] < arr[i] {
		arr[i], arr[p] = arr[p], arr[i]
		i, p = p, parent(p)
	}
}

func swiftDown(arr []int, cur, n int) {
	l, r := children(cur)
	for cur < n {
		maxIdx, maxVal := cur, arr[cur]
		if l < n && arr[l] > maxVal {
			maxIdx, maxVal = l, arr[l]
		}
		if r < n && arr[r] > maxVal {
			maxIdx, maxVal = r, arr[r]
		}
		if maxVal == arr[cur] {
			break
		}
		arr[cur], arr[maxIdx] = arr[maxIdx], arr[cur]
		cur = maxIdx
		l, r = children(cur)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}
