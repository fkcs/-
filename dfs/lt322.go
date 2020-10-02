package dfs

import (
	"math"
	"sort"
)

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
你可以认为每种硬币的数量是无限的。

示例 1：

输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1
*/

func coinChange(coins []int, amount int) int {
	// 进行排序，以便后面兑换先从大的面值开始
	sort.Ints(coins)
	min := math.MaxInt32
	dfs(coins, amount, amount, &min, len(coins)-1, 0)
	if min == math.MaxInt32 {
		return -1
	}
	return min
}

func dfs(coins []int, amount int, left int, min *int, index, cnt int) {
	// base case1 : 满足条件情况
	if left == 0 {
		if *min > cnt {
			*min = cnt
		}
		return
	}
	// base case2 : 越界，由于此条件会过滤掉base case1某些场景，因此放在后面
	if index < 0 || left < 0 {
		return
	}

	// recurse rule
	for i := left / coins[index]; i >= 0 && (i+cnt < *min); i-- {
		dfs(coins, amount, left-coins[index]*i, min, index-1, cnt+i)
	}
	return
}
