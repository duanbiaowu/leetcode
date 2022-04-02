package leetcode

func reverseWords(s string) string {
	left, right := 0, len(s)-1

	// 1. 去除左边空格
	for left < right && s[left] == ' ' {
		left++
	}

	// 2. 去除右边空格
	for right >= left && s[right] == ' ' {
		right--
	}
	if left > right {
		return ""
	}

	res := make([]byte, right-left+1)
	index := 0

	// 3. 反转字符串 && 去除多余空格
	for right >= left {
		for s[right] == ' ' && s[right-1] == ' ' {
			right--
		}
		res[index] = s[right]
		index++
		right--
	}

	// 4. 反转每个单词
	left, right = 0, 0
	for left < index {
		for right < index && res[right] != ' ' {
			right++
		}

		i, j := left, right-1
		for i < j {
			res[i], res[j] = res[j], res[i]
			i++
			j--
		}

		right++
		left = right
	}

	return string(res[:index])
}
