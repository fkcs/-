package sort

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{5, 4, 2, 1, 7, 3}
	quickSort(arr)
	t.Log(arr)
}
