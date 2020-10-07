package sort

func quickSort(arr []int) {
	sepQSort(arr, 0, len(arr)-1)
}

func sepQSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	// 获取分区点位置
	mid := partition(arr, start, end)
	// 左区域
	sepQSort(arr, start, mid-1)
	// 右区域
	sepQSort(arr, mid+1, end)
}

func partition(arr []int, start, end int) int {
	privot := arr[end]
	i := start
	for j := start; j < end; j++ {
		// j位置左侧是已处理区域，右侧是未处理区域
		// 如果当前值小于分区点，则交换到分区点左侧，加入到已处理区间尾部即Arr[i]位置
		// 此处决定是升序还是降序，当前是升序
		if arr[j] < privot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
