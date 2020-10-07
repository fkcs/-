package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{5, 4, 2, 1, 7, 3, 4}
	t.Log(arr)
	mergeSort(arr)
	t.Log(arr)
}
