package recursion

// 组合求和III
func dfsx(k, target int, index int, sum int, path *[]int, rst *[][]int, visted []bool) {
	// base case1: 把该放置在初始判断
	if sum == target && len(*path) == k {
		tmp := make([]int, k)
		copy(tmp, *path)
		*rst = append(*rst, tmp)
		return
	}
	// base case2：不满足的条件
	if index > 9 || sum > target || len(*path) > k {
		return
	}

	// recurse rule:
	for i := index; i <= 9; i++ {
		if visted[i] || i > target {
			continue
		}
		visted[i] = true
		*path = append(*path, i)
		dfsx(k, target, i+1, sum+i, path, rst, visted)
		visted[i] = false
		*path = (*path)[:len(*path)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	visted := make([]bool, 10)
	path := make([]int, 0)
	rst := make([][]int, 0)
	dfsx(k, n, 1, 0, &path, &rst, visted)
	return rst
}
