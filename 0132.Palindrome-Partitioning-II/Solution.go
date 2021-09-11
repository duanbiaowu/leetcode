package leetcode

import "math"

// https://leetcode-cn.com/problems/palindrome-partitioning-ii/solution/xiang-jie-liang-bian-dong-tai-gui-hua-ji-s5xr/
func minCut(s string) int {
	n := len(s)
	if n < 2 {
		return 0
	}
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			g[i][j] = true
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			g[i][j] = s[i] == s[j] && g[i+1][j-1]
		}
	}

	f := make([]int, n)
	for i := 0; i < n; i++ {
		if g[0][i] {
			continue
		}
		f[i] = math.MaxInt64
		for j := 0; j < i; j++ {
			if g[j+1][i] && f[j]+1 < f[i] {
				f[i] = f[j] + 1
			}
		}
	}

	return f[n-1]
}
