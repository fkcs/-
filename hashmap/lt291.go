package hashmap

/*
Given a pattern and a string str, find if str follows the same
pattern.
Here follow means a full match, such that there is a bijection
between a letter in pattern and a non-empty substring in str.
Examples:
1. pattern = "abab", str = "redblueredblue" should return true.
2. pattern = "aaaa", str = "asdasdasdasd" should return true.
3. pattern = "aabb", str = "xyzabcxzyabc" should return false.
*/
// 递归+hashmap
func wordPattern2(pattern string, s string) bool {
	hm := make(map[byte]string)
	hm_t := make(map[string]byte)
	return helper(pattern, s, 0, 0, hm, hm_t)
}

func helper(pattern string, s string, i, j int, hm map[byte]string, hm_t map[string]byte) bool {
	// base case1：同时达到边界
	if i == len(pattern) && j == len(s) {
		return true
	}
	// base case2 : 有且仅有一个提前达到边界
	if i == len(pattern) || j == len(s) {
		return false
	}
	// recurse rule 1: 当有重复的模式时
	if v, ok := hm[pattern[j]]; ok {
		// 长度超出范围
		if j+len(v) > len(s) {
			return false
		}
		// 判断hm_t是否匹配
		if hm_t[s[j:j+len(v)]] != pattern[j] {
			return false
		}
		return helper(pattern, s, i+1, j+len(v), hm, hm_t)
	}
	// recurse rule 2: 无重复模式时，需要切割后续字符串
	for index := j; index < len(s); index++ {
		str := s[j : index+1]
		// 为了保证完全匹配，跳过该位置
		if _, ok := hm_t[str]; ok {
			continue
		}
		hm[pattern[i]] = str
		hm_t[str] = pattern[i]
		if helper(pattern, s, i+1, index+1, hm, hm_t) {
			return true
		}
		// 回溯
		delete(hm, pattern[i])
		delete(hm_t, str)
	}
	return false
}
