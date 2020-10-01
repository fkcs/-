// 组合求和
func recurseCombine(arr []int, path *[]int, rst *[][]int, sum, target, index int) {
	// base case：索引越界，或者求和大于目标值
	if index >= len(arr) || sum > target {
		return
	}
	// base case : 求和等于目标值
	if sum == target {
		tmp := make([]int,len(*path))
		copy(tmp,*path)
		*rst = append(*rst,tmp)
		return
	}
	// 从以上二叉树变成N叉树，只是降低调用栈深度，节省内存空间，不会降低时间复杂度
	for i:=index; i<len(arr);i++ {
		if target < arr[i] {
			continue
		}
		*path = append(*path,arr[i])
		recurseCombine(arr,path,rst,sum+arr[i],target,i)
		// 回溯：由于存在共享数据量path，因此回溯到低层需要恢复现场
		*path = (*path)[:len(*path)-1]
	}
	return
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int,0)
	path := make([]int,0)
	recurseCombine(candidates,&path,&ret,0,target,0)
	return ret
}
