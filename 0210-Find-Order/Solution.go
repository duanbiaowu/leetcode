package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
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

	res := make([]int, numCourses)
	pos := 0
	// BFS TopSort
	for len(queue) > 0 {
		pre := queue[0]
		queue = queue[1:]
		res[pos] = pre
		pos++

		for _, cur := range adjacency[pre] {
			inDegrees[cur]--
			if inDegrees[cur] == 0 {
				queue = append(queue, cur)
			}
		}
	}

	if pos == numCourses {
		return res
	}
	return []int{}
}
