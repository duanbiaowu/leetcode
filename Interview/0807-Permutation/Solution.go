package leetcode

func permutation(S string) []string {
	var res []string
	backtrack([]byte(S), 0, &res)
	return res
}

func backtrack(S []byte, begin int, res *[]string) {
	if begin == len(S) {
		*res = append(*res, string(S))
		return
	}
	for i := begin; i < len(S); i++ {
		S[i], S[begin] = S[begin], S[i]
		backtrack(S, begin+1, res)
		S[i], S[begin] = S[begin], S[i]
	}
}

func permutation2(S string) []string {
	var visited = make([]bool, len(S))
	var path []byte
	var res []string
	backtrack2(S, path, &visited, &res)
	return res
}

func backtrack2(S string, path []byte, visited *[]bool, res *[]string) {
	if len(S) == len(path) {
		*res = append(*res, string(path))
		return
	}
	for i := 0; i < len(S); i++ {
		if !(*visited)[i] {
			path = append(path, S[i])
			(*visited)[i] = true
			backtrack2(S, path, visited, res)
			(*visited)[i] = false
			path = path[:len(path)-1]
		}
	}
}
