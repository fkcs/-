package hashmap

import "strings"

/*word pattern*/
/*
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
*/
// 采用双hash实现，一一对应
func wordPattern(pattern string, s string) bool {
	hm1 := make(map[byte]string)
	hm2 := make(map[string]byte)
	rst := strings.Split(s, " ")
	if len(pattern) != len(rst) {
		return false
	}
	for k, v := range rst {
		if v1, ok := hm1[pattern[k]]; ok {
			if v1 != v {
				return false
			}
		} else {
			hm1[pattern[k]] = v
		}

		if v2, ok := hm2[v]; ok {
			if v2 != pattern[k] {
				return false
			}
		} else {
			hm2[v] = pattern[k]
		}
	}
	return true
}
