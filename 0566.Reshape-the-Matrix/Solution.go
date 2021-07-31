package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	if m == 0 {
		return mat
	}
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	for i := 0; i < m*n; i++ {
		res[i/c][i%c] = mat[i/n][i%n]
	}

	return res
}
