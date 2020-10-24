package hashmap

/*
给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组（的长度）。

示例 1:

输入: [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
示例 2:

输入: [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
*/

// 哈希表实现，key保存差值，value对应索引
func findMaxLength(nums []int) int {
	diff := make([]int, len(nums)+1)
	hm := make(map[int]int)
	res := 0
	hm[0] = 0
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			diff[i] = diff[i-1] - 1
		} else {
			diff[i] = diff[i-1] + 1
		}
		if v, ok := hm[diff[i]]; ok {
			res = Max(res, i-v)
		} else {
			hm[diff[i]] = i
		}
	}
	return 0
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
