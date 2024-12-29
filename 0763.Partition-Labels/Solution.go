package leetcode

func partitionLabels(s string) []int {
	lastPos := [26]int{}

	// 计算字符串中每个字符最后出现位置
	for i, c := range s {
		lastPos[c-'a'] = i
	}

	res := []int{}

	// start 和 end 表示当前字符片段的开始和结束位置
	start, end := 0, 0

	// 遍历字符串
	// 维护当前字符片段的最后出现位置 (end 变量)
	for i, c := range s {
		// 如果当前字符最后出现位置 大于 已有的最后位置
		// 更新 已有的最后位置
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}

		// 如果当前字符片段可以分割为一个字符串
		if i == end {
			// 将当前字符片段的长度 (start 和 end 变量之间的距离) 加入到结果集合中
			res = append(res, end-start+1)
			// 更新 start 变量
			start = end + 1
		}
	}

	return res
}
