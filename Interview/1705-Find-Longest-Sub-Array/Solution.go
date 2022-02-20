package leetcode

func findLongestSubarray(array []string) []string {
	if len(array) == 0 {
		return nil
	}

	start, end := 0, 0
	preSum := 0
	// Map 记录上一个 preSum 位置
	// 如果出现相同的 preSum，说明上一个位置到这个位置和为 0
	leftIndexMap := make(map[int]int)
	leftIndexMap[0] = -1

	for i := range array {
		if isDigit(array[i][0]) {
			preSum++
		} else {
			preSum--
		}
		if j, ok := leftIndexMap[preSum]; ok {
			if i-j > end-start {
				start = j
				end = i
			}
		} else {
			leftIndexMap[preSum] = i
		}
	}
	return array[start+1 : end+1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
