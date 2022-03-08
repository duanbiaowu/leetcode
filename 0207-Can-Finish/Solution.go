package leetcode

// reference: https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
func canFinish(numCourses int, prerequisites [][]int) bool {
	inDegrees := make([]int, numCourses)
	adjacency := make([][]int, numCourses)
	var queue []int

	for i := 0; i < numCourses; i++ {
		adjacency[i] = []int{}
	}
	// Get the inDegree and adjacency of every course.
	for _, v := range prerequisites {
		inDegrees[v[0]]++
		adjacency[v[1]] = append(adjacency[v[1]], v[0])
	}

	// Get all the courses with the inDegree of 0.
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	// BFS TopSort
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		numCourses--
		for _, cur := range adjacency[pre] {
			inDegrees[cur]--
			if inDegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

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
