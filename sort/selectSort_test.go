package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{2, 4, 5, 9, 3, 1, 2}
	selectSort(arr)
	fmt.Println(arr)
}
