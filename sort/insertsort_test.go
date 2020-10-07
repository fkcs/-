package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{2, 4, 5, 9, 3, 1, 2}
	insertSort(arr)
	fmt.Println(arr)
}
