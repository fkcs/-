package sort

func selectSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		// 在未排序区域找最小值位置
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 交换最小值，放到有序区间末尾
		arr[i], arr[min] = arr[min], arr[i]
	}
	return
}
