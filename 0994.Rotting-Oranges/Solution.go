package leetcode

var (
	dx = []int{0, 1, 0, -1} // x: 上右下左
	dy = []int{-1, 0, 1, 0} // y: 上右下左
)

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

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				code := i*cols + j
				queue = append(queue, code)
				depth[code] = 0
			}
		}
	}

	cnt := 0

	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]

		for k := 0; k < 4; k++ {
			x := code/cols + dx[k]
			y := code%cols + dy[k]
			if x >= 0 && x < rows && y >= 0 && y < cols && grid[x][y] == 1 {
				grid[x][y] = 2
				newCode := x*cols + y

				queue = append(queue, newCode)
				depth[newCode] = depth[code] + 1
				cnt = depth[newCode]
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

	return cnt
}
