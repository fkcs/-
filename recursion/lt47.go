package recursion

import "sort"

// 题目描述：给定一个可包含重复数字地序列，返回所有不重复地全排列
func dfs1(nums []int, path *[]int, rst *[][]int, visted []bool) {
	// base case，如果path地长度等于nums，则认为该路径是其中一种排列结果
	if len(*path) == len(nums) {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*rst = append(*rst, tmp)
		return
	}
	// recurse rule：N叉树
	for i := 0; i < len(nums); i++ {
		// 当前节点被访问过
		// 或者当前节点与前一个节点相等进行去重，保证一定时从左往右逐个填入
		if visted[i] || (i > 0 && nums[i] == nums[i-1] && !visted[i-1]) {
			continue
		}
		visted[i] = true
		*path = append(*path, nums[i])
		dfs1(nums, path, rst, visted)
		// 进行回溯，当返回低层时恢复现场
		visted[i] = false
		*path = (*path)[:len(*path)-1]
	}
	return
}
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	// 排序目的：1）进行升序排列；2）相同数字地值聚合在一起
	sort.Ints(nums)
	visted := make([]bool, len(nums))
	path := make([]int, 0)
	rst := make([][]int, 0)
	dfs1(nums, &path, &rst, visted)
	return rst
}
