package leetcode

// reference: https://leetcode-cn.com/problems/eight-queens-lcci/comments/1392929
// N 皇后在 N × N 棋盘上的各种摆法，
// 其中每个皇后都不同行、不同列，也不在对角线上。
func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}

	var res [][]string
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}

	// 第 k 列有没有被占用
	col := make([]bool, n)
	// 主对角线方向有没有被占用（左上到右下）
	// 该方向的x=row-col是固定的，范围为-n+1~n-1共2n-1个数，
	// n-x之后范围是1~2n-1，用2n的数组就可以容纳
	dg := make([]bool, n<<1)
	// 副对角线方向有没有被占用（右上到左下）
	// 该方向的x=row+col是固定的，范围为0~2n-2共2n-1个数，
	// 用2n的数组也可以表示2n-1条对角方向
	udg := make([]bool, n<<1)

	backtrack(0, n, grid, col, dg, udg, &res)

	return res
}

func backtrack(begin, n int, grid [][]byte, col, dg, udg []bool, res *[][]string) {
	if begin == n {
		var row []string
		for i := range grid {
			row = append(row, string(grid[i]))
		}
		*res = append(*res, row)
		return
	}

	for j := 0; j < n; j++ {
		if !col[j] && !dg[n-begin+j] && !udg[begin+j] {
			grid[begin][j] = 'Q'
			col[j] = true
			dg[n-begin+j] = true
			udg[begin+j] = true
			backtrack(begin+1, n, grid, col, dg, udg, res)
			grid[begin][j] = '.'
			col[j] = false
			dg[n-begin+j] = false
			udg[begin+j] = false
		}
	}
}
