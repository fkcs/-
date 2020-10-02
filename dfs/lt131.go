package dfs

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。

示例:
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
*/

func partition(s string) [][]string {
	path := make([]string, 0)
	rst := make([][]string, 0)
	helper(s, 0, &path, &rst)
	return rst
}

func helper(s string, index int, path *[]string, rst *[][]string) {
	// base case : 当遍历完所有的字符串
	if index == len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*rst = append(*rst, tmp)
		return
	}
	for i := index; i < len(s); i++ {
		sub := s[index : i+1]
		if !isTrue(sub) {
			continue
		}
		*path = append(*path, sub)
		helper(s, i+1, path, rst)
		// 回溯，因为共享同一个内存变量
		*path = (*path)[:len(*path)-1]
	}
	return
}

func isTrue(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
