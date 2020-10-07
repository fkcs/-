package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{2, 4, 5, 9, 3, 1, 2}
	bubbleSort(arr)
	fmt.Println(arr)
}
