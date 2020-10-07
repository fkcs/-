package sort

func mergeSort(arr []int) {
	recurse(arr, 0, len(arr)-1)
}

func recurse(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := start + (end-start)>>1
	recurse(arr, start, mid)
	recurse(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmp := make([]int, 0)
	i, j := start, mid+1
	// 合并两个有序区间
	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	tmp = append(tmp, arr[i:mid+1]...)
	tmp = append(tmp, arr[j:end+1]...)
	copy(arr[start:end+1], tmp)
}
