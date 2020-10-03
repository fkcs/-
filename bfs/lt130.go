package bfs

/*
给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。

示例:
X X X X
X O O X
X X O X
X O X X

运行你的函数后，矩阵变为：
X X X X
X X X X
X X X X
X O X X
*/

func helper(bord [][]byte, x, y int) {
	if bord[x][y] == 'X' {
		return
	}

	xp := []int{-1, 1, 0, 0}
	yp := []int{0, 0, 1, -1}
	pos := x*len(bord[0]) + y
	queue := make([]int, 0)
	bord[x][y] = 'Y'
	queue = append(queue, pos)

	for len(queue) > 0 {
		pos = queue[0]
		queue = queue[1:]
		xc, yc := pos/len(bord[0]), pos%len(bord[0])
		for i := 0; i < 4; i++ {
			// xt/yt不能是xc/yc
			xt, yt := xp[i]+xc, yp[i]+yc
			if xt >= 0 && xt < len(bord) && yt >= 0 && yt < len(bord[0]) && bord[xt][yt] == 'O' {
				bord[xt][yt] = 'Y'
				pos = xt*len(bord[0]) + yt
				queue = append(queue, pos)
			}
		}
	}
}

func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	for i := 0; i < len(board); i++ {
		helper(board, i, 0)
		helper(board, i, len(board[0])-1)
	}
	for j := 0; j < len(board[0]); j++ {
		helper(board, 0, j)
		helper(board, len(board)-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}
