package leetcode

// reference: https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/solution/tong-ji-zui-gao-fen-de-jie-dian-shu-mu-b-n810/
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	if n == 0 {
		return 0
	}

	children := make([][]int, n)
	for i := 1; i < n; i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}

	maxScore := 0
	res := 0

	var dfs func(int) int
	dfs = func(node int) int {
		score, size := 1, n-1
		for _, v := range children[node] {
			sz := dfs(v)
			score *= sz
			size -= sz
		}
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			res++
		} else if score > maxScore {
			maxScore = score
			res = 1
		}
		return n - size
	}

	dfs(0)
	return res
}
