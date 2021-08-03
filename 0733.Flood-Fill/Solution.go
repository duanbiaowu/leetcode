package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, x, y, originColor, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) ||
		image[x][y] != originColor || originColor == newColor {
		return
	}

	image[x][y] = newColor

	dfs(image, x-1, y, originColor, newColor)
	dfs(image, x+1, y, originColor, newColor)
	dfs(image, x, y-1, originColor, newColor)
	dfs(image, x, y+1, originColor, newColor)
}

func bfs(image [][]int, x, y, originColor, newColor int) [][]int {
	if originColor == newColor {
		return image
	}

	dx := []int{0, 1, 0, -1} // x: 上右下左
	dy := []int{-1, 0, 1, 0} // y: 上右下左

	var queue [][]int
	queue = append(queue, []int{x, y})
	image[x][y] = newColor

	for len(queue) > 0 {
		for i := 0; i < 4; i++ {
			xx, yy := queue[0][0]+dx[i], queue[0][1]+dy[i]
			if xx >= 0 && xx < len(image) && yy >= 0 && yy < len(image[0]) &&
				image[xx][yy] == originColor && originColor != newColor {
				queue = append(queue, []int{xx, yy})
				image[xx][yy] = newColor
			}
		}
		queue = queue[1:]
	}

	return image
}
