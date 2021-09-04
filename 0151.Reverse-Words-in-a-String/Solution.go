package leetcode

func reverseWords(s string) string {
	left, right := 0, len(s)-1

	// 去掉首端空格
	for ; left <= right && s[left] == ' '; left++ {
	}
	// 去掉末端空格
	for ; right >= left && s[right] == ' '; right-- {
	}
	if left > right {
		return ""
	}

	res := make([]byte, right-left+1)
	index := 0
	// 反转字符串 && 去掉多余空格
	for right >= left {
		// s[right + 1] != ' ' => 保留单词间至多一个空格
		if s[right] != ' ' || s[right+1] != ' ' {
			res[index] = s[right]
			index++
		}
		right--
	}

	// 反转每个单词
	for left, right = 0, 0; left < index; {
		for ; right < index && res[right] != ' '; right++ {
		}

		for i, j := left, right-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}

		right++
		left = right
	}

	return string(res[:index])
}
