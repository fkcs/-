package sort

func insertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		// 查找数据在有序组内位置
		for ; j >= 0; j-- {
			if arr[j] > value {
				// 数据移动
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		// 插入数据
		arr[j+1] = value
	}
}
