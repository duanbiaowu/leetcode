package leetcode

var (
	dx = []int{0, 1, 0, -1} // x: 上右下左
	dy = []int{-1, 0, 1, 0} // y: 上右下左
)

// 题目抽象后: 矩阵的渲染/着色问题
// 直接使用 BFS 算法即可，每一分钟渲染一部分橘子 (改变一部分橘子的状态)
// 值 0 代表空单元格
// 值 1 代表新鲜橘子
// 值 2 代表腐烂的橘子
// 所以 BFS 遍历时只需要判定值 == 1 的橘子即可，当值变为 2 时，直接加入队列
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return -1
	}
	cols := len(grid[0])

	// 使用矩阵数组索引作为队列值和 Map-key
	// 当然，也可以使用二维数组 var queue [][2]int
	var queue []int
	var depth = make(map[int]int)

	// 先将已经腐烂的橘子加入队列
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				code := i*cols + j
				queue = append(queue, code)
				depth[code] = 0
			}
		}
	}

	// 初始化最终的分钟数为 0
	minutes := 0

	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]

		for k := 0; k < 4; k++ {
			x := code/cols + dx[k]
			y := code%cols + dy[k]

			// 获取当前已经腐烂的橘子最邻近的橘子
			if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
				grid[x][y] = 2
				// 计算邻近橘子的坐标
				newCode := x*cols + y
				// 将邻近的橘子加入队列
				queue = append(queue, newCode)
				// 更新邻近橘子的腐烂分钟数
				// 题目定义: 每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。
				// 例如当一个 (初始化时) 2 分钟腐烂的橘子，其邻近的橘子腐烂时间为 3 分钟
				depth[newCode] = depth[code] + 1
				// 更新最新的分钟数
				minutes = depth[newCode]
			}
		}
	}

	// 检测是否还有橘子未腐烂
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
