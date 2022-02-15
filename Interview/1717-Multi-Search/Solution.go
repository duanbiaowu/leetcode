package leetcode

func multiSearch(big string, smalls []string) [][]int {
	var res [][]int
	// 构造一个字符字典 保存所有字符在字符串中出现的位置
	dict := [26][]int{}
	for i := range big {
		dict[big[i]-'a'] = append(dict[big[i]-'a'], i)
	}

	for i := range smalls {
		if len(smalls[i]) == 0 {
			res = append(res, []int{})
			continue
		}

		// 判定所有字符都在 big 中出现过
		flag := false
		for j := range smalls[i] {
			if dict[smalls[i][j]-'a'] == nil {
				flag = true
				break
			}
		}
		if flag {
			res = append(res, []int{})
			continue
		}

		// 判定是否为子字符串
		var row []int
		for _, k := range dict[smalls[i][0]-'a'] {
			if k+len(smalls[i]) <= len(big) && big[k:k+len(smalls[i])] == smalls[i] {
				row = append(row, k)
			}
		}
		res = append(res, row)
	}

	return res
}
