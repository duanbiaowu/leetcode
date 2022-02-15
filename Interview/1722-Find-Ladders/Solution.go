package leetcode

func findLadders(beginWord string, endWord string, wordList []string) []string {
	var path []string
	var visited = make(map[string]struct{})
	var res []string
	backtrack(beginWord, endWord, wordList, path, visited, &res)
	return res
}

// 回溯的时候不用回溯 visited 数组
// 因为之前走这个点走不通，那么无论你怎么走只要到了这个点，肯定还是走不通
func backtrack(begin, end string, wordList, path []string, visit map[string]struct{}, res *[]string) bool {
	path = append(path, begin)
	visit[begin] = struct{}{}

	if begin == end {
		*res = append(*res, path...)
		return true
	}

	for _, word := range wordList {
		if _, ok := visit[word]; ok {
			continue
		}
		if checkTranslate(begin, word) && backtrack(word, end, wordList, path, visit, res) {
			return true
		}
		// 此处区别于一般回溯：visited 无需再次回溯
	}

	path = path[:len(path)-1]
	return false
}

// 两个字符串是否恰好 1 位不同
func checkTranslate(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	cnt := 0
	for i := range a {
		if a[i] != b[i] {
			cnt++
		}
	}
	return cnt == 1
}
