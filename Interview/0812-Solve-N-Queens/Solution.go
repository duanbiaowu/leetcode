package leetcode

// 概述
// N 皇后问题 研究的是如何将 N 个皇后放置在 N×N 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 皇后的走法是：可以横直斜走，格数不限。因此要求皇后彼此之间不能相互攻击，
// 等价于要求任何两个皇后都不能在同一行、同一列以及同一条斜线上。
// 显而易见，暴力解法的思路就是枚举所有坐标，然后判断是否符合条件。
// 具体来说，先讲第一个皇后放在第一行第一列，然后放第二个皇后，判断是否符合条件，不符合就回溯。
// 暴力算法核心: 将 N 个皇后放置在 N×N 的棋盘上的所有可能的情况，
//	  并对每一种情况判断是否满足皇后彼此之间不相互攻击。
// 暴力枚举的时间复杂度是非常高的，因此必须利用限制条件加以优化。

// 题解 1
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

// 题解 2
// Reference: https://leetcode.cn/problems/n-queens/solutions/398929/nhuang-hou-by-leetcode-solution
// 因为每个皇后必须位于不同的行和列
// 所以当 N 个皇后放在棋盘上面时，每个皇后必须满足下面 3 个条件约束:
//
//	每一行只有 1 个皇后
//	每一列只有 1 个皇后
//	每一条对角线只有 1 个皇后 (包括左上到右下的对角线，和右上到左下的对角线)
//
// 为了满足这些约束条件，可以使用回溯算法依次尝试每一个坐标，最后得出每个皇后的位置
// 单个回溯过程的终止条件 == 当前行数等于 N 时，说明所有皇后都已经放置完毕

// 但同时，回溯 (暴力枚举) 的时间复杂度是非常高的，因此必须利用限制条件加以优化。
// 剪枝策略 == 当前放置的皇后位置，不能在同一列、同一条对角线上
// 为了降低总的时间复杂度，理想情况下，可以使用 O(1) 的时间复杂度来判断当前位置是否可以放置一个皇后
// 而为了实现这个时间复杂度，可以使用 3 个 Set 集合存储已经放置的皇后的 列、主对角线、副对角线

// 下面是一个 4 * 4 的棋盘示例图:
//    0   1   2   3
// 0  .   .   .   .
// 1  .   .   .   .
// 2  .   .   .   .
// 3  .   .   .   .

func solveNQueens2(n int) [][]string {
	if n < 1 {
		return nil
	}

	// 表示每个皇后在每一行的具体位置
	queens := make([]int, n)
	for i := range queens {
		queens[i] = -1
	}

	// 3 个 Set 集合如下:
	// 表示每一列是否存在皇后
	// 例如下面的棋盘中，每一列只有一个皇后
	//    0   1   2   3
	// 0  Q   .   .   .
	// 1  .   Q   .   .
	// 2  .   .   Q   .
	// 3  .   .   .   Q

	columns := make(map[int]bool)

	// 表示主对角线 (左上到右下) 是否存在皇后
	// 如何表示主对角线的坐标呢？
	// 每一条主对角斜线的 行下标和列下标的差 是相等的

	// 例如下面的棋盘中，主对角线的坐标差值等于 0
	// 4 个坐标 (0, 0), (1, 1), (2, 2), (3, 3)

	//    0   1   2   3
	// 0  Q   .   .   .
	// 1  .   Q   .   .
	// 2  .   .   Q   .
	// 3  .   .   .   Q
	dgs := make(map[int]bool)

	// 表示副对角线 (右上到左下) 是否存在皇后
	// 如何表示副对角线的坐标呢？
	// 每一条主对角斜线的 行下标和列下标的和 是相等的

	// 例如下面的棋盘中，副对角线的坐标和等于 3
	// 4 个坐标 (0, 3), (1, 2), (2, 1), (3, 0)

	//    0   1   2   3
	// 0  .   .   .   Q
	// 1  .   .   Q   .
	// 2  .   Q   .   .
	// 3  Q   .   .   .
	dgs2 := make(map[int]bool)

	var res [][]string

	backtrack2(n, 0, queens, columns, dgs, dgs2, &res)

	return res
}

func backtrack2(n, row int, queens []int, cols, dgs, dgs2 map[int]bool, res *[][]string) {
	// 回溯终止条件: 当前行数等于参数 N (皇后的目标个数)
	// 将当前各个皇后的位置渲染为棋盘字符串格式，然后加入到结果集合中
	if row == n {
		*res = append(*res, generateBoard(queens, n))
		return
	}

	// 尝试在当前行的每一列放置一个皇后
	// 由于每个皇后必须位于不同的列，所以:
	//    第一个皇后有 N 列可以选择
	//    第二个皇后有 N-1 列可以选择
	//    ... 以此类推
	//    最后所有可能的情况不会超过 N！种，因此时间复杂度是 O(N!)
	for i := 0; i < n; i++ {
		// 剪枝策略
		// 确保新放置的皇后不会和已经放置的皇后冲突
		if cols[i] || dgs[row-i] || dgs2[row+i] {
			continue
		}

		// 放置一个皇后
		// 更新当前皇后所在行对应的列，以及对应的主/副对角线信息
		queens[row] = i
		cols[i], dgs[row-i], dgs2[row+i] = true, true, true

		// 继续放置下一个皇后，row+1 表示下一行
		backtrack2(n, row+1, queens, cols, dgs, dgs2, res)

		// 撤销当前皇后的位置，回溯到上一个状态
		queens[row] = -1
		cols[i], dgs[row-i], dgs2[row+i] = false, false, false
	}
}

// 根据皇后的位置生成对应的棋盘字符串形式
func generateBoard(queens []int, n int) []string {
	board := []string{}

	for i := 0; i < n; i++ {
		// 每一行的固定默认样式
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}

		// 将当前行的皇后位置设置为 'Q'
		row[queens[i]] = 'Q'

		board = append(board, string(row))
	}

	return board
}
