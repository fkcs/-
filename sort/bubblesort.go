package sort

func bubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	// 提前冒泡结束的标志
	var isFinish bool
	for i := 0; i < len(arr); i++ {
		isFinish = false
		for j := 0; j < len(arr)-i-1; j++ {
			// 此处是决定升序还是降序排列，当前是升序
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isFinish = true
			}
		}
		// 没有数据交换提前退出
		if !isFinish {
			break
		}
	}
	return
}
