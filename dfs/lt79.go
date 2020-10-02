package dfs

// leetcode 79 单词搜索
/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
提示：
board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/
func exist(bord [][]byte, word string) bool {
	visted := make([][]bool, len(bord))
	for i := 0; i < len(bord); i++ {
		visted[i] = make([]bool, len(bord[0]))
	}
	for i := 0; i < len(bord); i++ {
		for j := 0; j < len(bord[0]); j++ {
			if findWords(bord, i, j, 0, word, visted) {
				return true
			}
		}
	}
	return false
}

func findWords(bord [][]byte, x, y int, index int, word string, visted [][]bool) bool {
	if index == len(word) {
		return true
	}
	if x < 0 || x >= len(bord) || y < 0 || y >= len(bord[0]) || visted[x][y] ||
		bord[x][y] != word[index] {
		return false
	}
	xp := []int{-1, +1, 0, 0}
	yp := []int{0, 0, -1, +1}
	visted[x][y] = true
	for i := 0; i < 4; i++ {
		if findWords(bord, x+xp[i], y+yp[i], index+1, word, visted) {
			return true
		}
	}
	visted[x][y] = false
	return false
}
