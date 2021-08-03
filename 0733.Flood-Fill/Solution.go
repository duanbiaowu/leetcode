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
