// 全排列实现
func helper(arr []int, path *[]int, rst *[][]int, visted []bool) {
	// base case
	if len(*path) == len(arr) {
		tmp := make([]int,len(*path))
		copy(tmp,*path)
		*rst = append(*rst,tmp)
		return
	}
	// recurse rule
	for i:=0; i<len(arr); i++ {
		if visted[i] {
			continue
		}
		visted[i] = true
		*path = append(*path,arr[i])
		helper(arr,path,rst,visted)
		visted[i] = false
		*path = (*path)[:len(*path)-1]
	}
	return
}
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	visted := make([]bool,len(nums))
	rst := make([][]int,0)
	path := make([]int,0)
	helper(nums,&path,&rst,visted)
	return rst
}
