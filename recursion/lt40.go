// 组合求和
func recurseSum2(candidates []int, ret *[][]int, path *[]int, target, index int) {
	if target != 0 && index == len(candidates) {
		return
	}
	if target == 0 {
		tmp := make([]int,len(*path))
		copy(tmp,*path)
		*ret = append(*ret,tmp)
		return
	}

	for i:=index; i<len(candidates); i++ {
		if target-candidates[i] < 0 {
			continue
		}
		*path = append(*path,candidates[i])
		recurseSum2(candidates,ret,path,target-candidates[i],i+1)
		*path = (*path)[:len(*path)-1]
		for i<len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int,0)
	path := make([]int,0)
	recurseSum2(candidates,&ret,&path,target,0)
	return ret
}
