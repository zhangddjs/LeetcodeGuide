package sortalgo

import "testing"

type sortFunc func([]int)

var testCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{"empty array", []int{}, []int{}},
	{"single element", []int{42}, []int{42}},
	{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	{"duplicates", []int{3, 3, 3, 1, 1, 2, 2}, []int{1, 1, 2, 2, 3, 3, 3}},
	{"negative numbers", []int{-1, -5, 3, 0, -2}, []int{-5, -2, -1, 0, 3}},
	{"large sorted", generateSorted(1000), generateSorted(1000)},
	{"large reverse", generateReverse(1000), generateSorted(1000)},
	{"large duplicates", generateDuplicates(500), generateDuplicatesSorted(500)},
}

func generateSorted(n int) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = i + 1
	}
	return arr
}

func generateReverse(n int) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = n - i
	}
	return arr
}

func generateDuplicates(n int) []int {
	arr := make([]int, n)
	for i := range n {
		arr[i] = (i % 10) + 1
	}
	return arr
}

func generateDuplicatesSorted(n int) []int {
	arr := generateDuplicates(n)
	expected := make([]int, n)
	copy(expected, arr)
	for i := 0; i < len(expected)-1; i++ {
		for j := 0; j < len(expected)-1-i; j++ {
			if expected[j] > expected[j+1] {
				expected[j], expected[j+1] = expected[j+1], expected[j]
			}
		}
	}
	return expected
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testSort(t *testing.T, sortAlgo sortFunc, algoName string) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arr := make([]int, len(tc.input))
			copy(arr, tc.input)
			sortAlgo(arr)
			if !equal(arr, tc.expected) {
				t.Errorf("%s failed for %s: expected %v, got %v", algoName, tc.name, tc.expected, arr)
			}
		})
	}
}

func TestHeapsort(t *testing.T) {
	testSort(t, heapsort, "HeapSort")
}

// func TestQuickSort(t *testing.T) {
// 	testSort(t, quicksort, "QuickSort")
// }
//
// func TestMergeSort(t *testing.T) {
// 	testSort(t, mergesort, "MergeSort")
// }
