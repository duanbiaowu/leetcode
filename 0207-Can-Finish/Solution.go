package leetcode

// reference: https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
func canFinish(numCourses int, prerequisites [][]int) bool {
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

	// BFS 搜索
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]

		// 当前课程的前置课程数量为 0, 可以直接学习
		// 将需要学习的课程总数量 - 1
		numCourses--

		// 从当前课程开始，进行 BFS 搜索将其作为前缀的课程的 后缀课程
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

	// 如果已经学习完课程，那么此时课程总数量应该等于 0
	// 因为每学习一门前置课程数量为 0 的课程后
	// 都要将学习的课程总数量 - 1
	return numCourses == 0
}

func canFinish2(numCourses int, prerequisites [][]int) bool {
	flags := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjacency[i] = []int{}
	}

	// Get the inDegree and adjacency of every course.
	for _, v := range prerequisites {
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(adjacency, flags, i) {
			return false
		}
	}
	return true
}

// 未被 DFS 访问：i == 0；
// 已被其他节点启动的 DFS 访问：i == -1；
// 已被当前节点启动的 DFS 访问：i == 1。
func dfs(adjacency [][]int, flags []int, index int) bool {
	// 节点在本轮中被第二次访问，说明图存在环
	if flags[index] == 1 {
		return false
	}
	if flags[index] == -1 {
		return true
	}
	flags[index] = 1
	for _, cur := range adjacency[index] {
		if !dfs(adjacency, flags, cur) {
			return false
		}
	}
	flags[index] = -1
	return true
}
