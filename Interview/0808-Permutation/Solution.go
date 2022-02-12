package leetcode

import "sort"

func permutation(S string) []string {
	// S 未排序
	s := []byte(S)
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	var res []string
	backtrack(string(s), []byte{}, make([]bool, len(S)), &res)
	return res
}

func backtrack(S string, path []byte, visited []bool, res *[]string) {
	if len(S) == len(path) {
		*res = append(*res, string(path))
		return
	}
	for i := 0; i < len(S); i++ {
		// 去重剪枝
		if visited[i] || i > 0 && S[i] == S[i-1] && !visited[i-1] {
			continue
		}
		path = append(path, S[i])
		visited[i] = true
		backtrack(S, path, visited, res)
		visited[i] = false
		path = path[:len(path)-1]
	}
}
