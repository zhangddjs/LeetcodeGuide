// Generated on 2026-01-14 08:18:26
// Daily practice file: sortalgo_20260114.go

package sortalgo

// --------------------------------------------------------------

func heapsort(arr []int) {
	buildheap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, 0, i)
	}
}

func buildheap(arr []int) {
	for i := parent(len(arr) - 1); i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
}

func heapify(arr []int, i, j int) {
	for true {
		l, r := children(i)
		maxIdx := i
		if l < j && arr[l] > arr[maxIdx] {
			maxIdx = l
		}
		if r < j && arr[r] > arr[maxIdx] {
			maxIdx = r
		}
		if i == maxIdx {
			break
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
		i = maxIdx
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func children(i int) (int, int) {
	return i*2 + 1, i*2 + 2
}

// --------------------------------------------------------------

// The general idea is that you pick a "pivot" element from the array, then partition the array into two sub-arrays
// one with elements less than the pivot and one with elements greater than the pivot.
// Then, you recursively apply the same process to those sub-arrays
// Finally, you combine them all back together, and you end up with a sorted array.
func quicksort(arr []int) {
	quicksortarr(arr, 0, len(arr)-1)
}

func quicksortarr(arr []int, i, j int) {
	if i < j {
		p := partition(arr, i, j)
		quicksortarr(arr, i, p-1)
		quicksortarr(arr, p+1, j)
	}
}

func partition(arr []int, l, r int) int {
	x := arr[r]
	i, j := l, l
	for j < r {
		if arr[j] < x {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		j++
	}
	arr[r], arr[i] = arr[i], arr[r]
	return i
}

// --------------------------------------------------------------

// The basic idea is to split the array into two halves, recursively sort each half, and then merge the two sorted halves back together.
// Divide: If the array has more than one element, split it into two halves.
// Conquer: Recursively apply merge sort to each half until each sub-array has only one element (which is considered sorted).
// Merge: Combine the two sorted halves by comparing their elements and merging them in order, resulting in a fully sorted array.
func mergesort(arr []int) {
	mergesortarr(arr, 0, len(arr)-1)
}

func mergesortarr(arr []int, l, r int) {
	if l < r {
		m := (l + r) >> 1
		mergesortarr(arr, l, m)
		mergesortarr(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	nl, nr := m-l+1, r-m
	l1, l2 := make([]int, nl), make([]int, nr)
	for i := range nl {
		l1[i] = arr[l+i]
	}
	for i := range nr {
		l2[i] = arr[m+1+i]
	}
	i, j, k := 0, 0, l
	for i < nl && j < nr {
		if l1[i] < l2[j] {
			arr[k] = l1[i]
			i++
		} else {
			arr[k] = l2[j]
			j++
		}
		k++
	}
	for ; i < nl; i, k = i+1, k+1 {
		arr[k] = l1[i]
	}
	for ; j < nr; j, k = j+1, k+1 {
		arr[k] = l2[j]
	}
}

// --------------------------------------------------------------

// It starts from the second element and compares it to the elements before it, inserting it into its correct position in the sorted part. It continues this process until the entire array is sorted.
func insertsort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > x; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = x
	}
}

// --------------------------------------------------------------

// It works by repeatedly stepping through the list, comparing adjacent elements, and swapping them if they’re in the wrong order. This process is repeated until the list is sorted. Typically, with each pass, the largest unsorted element "bubbles up" to its correct position.
func bubblesort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
