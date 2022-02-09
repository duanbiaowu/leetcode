package leetcode

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	// 图中可能存在自环和平行边
	g := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make(map[int]bool)
	}
	for _, edge := range graph {
		if edge[0] != edge[1] {
			g[edge[0]][edge[1]] = true
		}
	}

	// BFS
	var queue []int
	queue = append(queue, start)
	visit := make([]bool, n)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if !visit[cur] {
			visit[cur] = true
			if cur == target {
				return true
			}
			for i, _ := range g[cur] {
				if !visit[i] {
					queue = append(queue, i)
				}
			}
		}
	}
	return false
}
