package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 使用 数组 表示每个课程作为其他课程的前置数量总和
	// 例如 课程 2, 3, 4 的前置课程都是课程 1, 表示为 inDegrees[1] = 3
	inDegrees := make([]int, numCourses)
	// 使用 邻接表 表示课程之间的前置拓扑关系矩阵
	adjacency := make([][]int, numCourses)

	// 初始化邻接表内部的的二维数组
	for i := 0; i < numCourses; i++ {
		adjacency[i] = []int{}
	}

	// 计算课程的前置数量和邻接表关系矩阵
	for _, v := range prerequisites {
		inDegrees[v[0]]++
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	// 如果课程的前置课程数量为 0
	// 说明该课程可以直接学习，加入 BFS 队列
	var queue []int
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 记录课程学习顺序
	res := make([]int, numCourses)
	// 记录课程学习顺序的位置索引
	pos := 0

	// BFS 搜索
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]

		// 更新课程学习顺序
		res[pos] = pre
		// 更新索引
		pos++

		for _, cur := range adjacency[pre] {
			// 学完了当前课程，后缀课程数量 - 1
			inDegrees[cur]--
			// 如果后缀课程的前置课程数量为 0
			// 说明该课程可以直接学习，加入 BFS 队列
			if inDegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	// 可以学习完所有课程
	if pos == numCourses {
		return res
	}

	return []int{}
}
