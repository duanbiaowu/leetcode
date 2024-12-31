package leetcode

// reference: https://leetcode-cn.com/problems/eight-queens-lcci/comments/1392929
// N 皇后在 N × N 棋盘上的各种摆法，
// 其中每个皇后都不同行、不同列，也不在对角线上。
func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}

	// 使用二维数组表示棋盘
	// 初始化为没有棋子的空位，使用字符 `.` 表示
	grid := make([][]byte, n)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}

	// 第 k 列有没有被占用
	// 皇后可以攻击同一列的棋子
	cols := make([]bool, n)

	// 主对角线方向有没有被占用（左上到右下）
	// 该方向的 x=row-col 是固定的，范围为 -n+1~n-1, 共 2n-1 个数
	// n-x 之后范围是 1~2n-1，用 2n 的数组就可以容纳
	dgs := make([]bool, n<<1)

	// 副对角线方向有没有被占用（右上到左下）
	// 该方向的 x=row+col 是固定的，范围为 0~2n-2, 共 2n-1 个数
	// 用 2n 的数组也可以表示 2n-1 条对角方向
	udgs := make([]bool, n<<1)

	var res [][]string

	backtrack(n, 0, grid, cols, dgs, udgs, &res)

	return res
}

// 回溯算法核心:
// 为每个皇后
// n 表示皇后的个数
// start 表示当前摆放的 行 对应的索引
// grid 表示棋盘
// cols, dgs, udgs 表示坐标的使用情况，用于检测攻击范围
func backtrack(n, start int, grid [][]byte, cols, dgs, udgs []bool, res *[][]string) {
	// 剪枝条件: 当前棋盘上的皇后数量 等于 参数 n
	if start == n {
		var row []string
		for i := range grid {
			row = append(row, string(grid[i]))
		}

		*res = append(*res, row)
		return
	}

	for i := 0; i < n; i++ {
		// 皇后可以攻击与之处在 同一行，或同一列，或同一斜线上的棋子
		// 必须确保当前棋子落脚坐标不在攻击范围内
		if !cols[i] && !dgs[n-start+i] && !udgs[start+i] {
			grid[start][i] = 'Q'
			cols[i] = true
			dgs[n-start+i] = true
			udgs[start+i] = true

			// 回溯摆放下一个皇后
			backtrack(n, start+1, grid, cols, dgs, udgs, res)

			grid[start][i] = '.'
			cols[i] = false
			dgs[n-start+i] = false
			udgs[start+i] = false
		}
	}
}
