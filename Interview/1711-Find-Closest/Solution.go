package leetcode

func findClosest(words []string, word1 string, word2 string) int {
	start, end, res := -10000, 10000, 10000
	for i := range words {
		if words[i] == word1 {
			start = i
		}
		if words[i] == word2 {
			end = i
		}
		res = min(res, abs(end-start))
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
