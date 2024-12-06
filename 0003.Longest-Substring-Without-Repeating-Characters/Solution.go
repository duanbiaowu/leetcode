package leetcode

// 暴力解法(超时)
//func lengthOfLongestSubstring(s string) int {
//	res := 0
//
//	for i := 0; i < len(s); i++ {
//		for j := i + 1; j <= len(s); j++ {
//			if !hasDuplicateChar(s, i, j) {
//				res = max(res, j-i)
//			}
//		}
//	}
//	return res
//}

// 暴力解法2 (不超时)
// func lengthOfLongestSubstringSlow(s string) int {
// 	longest := 0
// 	n := len(s)

// 	for i := 0; i < n; i++ {
// 		set := make(map[byte]struct{})

// 		for j := i; j < n; j++ {
// 			if _, ok := set[s[j]]; ok {
// 				break
// 			}

// 			set[s[j]] = struct{}{}
// 		}

// 		longest = max(longest, len(set))
// 	}

// 	return longest
// }

// 滑动窗口-1
//func lengthOfLongestSubstring(s string) int {
//	var charIndex [256]int
//	maxLen, substringLen, lastRepeatIndex := 0, 0, 0
//
//	for i := 0; i < len(s); i++ {
//		if index := charIndex[s[i]]; index > 0 {
//			// 更新阶段性最长无重复字符串长度
//			maxLen = max(maxLen, substringLen)
//			lastRepeatIndex = max(lastRepeatIndex, index)
//			substringLen = i + 1 - lastRepeatIndex
//		} else {
//			substringLen++
//		}
//		charIndex[s[i]] = i + 1
//	}
//
//	return max(maxLen, substringLen)
//}

// 滑动窗口-2
func lengthOfLongestSubstring(s string) int {
	var win [256]int
	res, n := 0, len(s)

	for left, right := 0, 0; right < n; right++ {
		c := s[right]
		win[c]++

		// 遇到重复的字符时，开始收缩窗口大小
		for win[c] > 1 {
			win[s[left]]--
			left++
		}

		// 更新已知的最大窗口
		res = max(res, right-left+1)
	}

	return res
}

func hasDuplicateChar(s string, start, end int) bool {
	sMap := make(map[uint8]int)

	for i := start; i < end; i++ {
		if _, ok := sMap[s[i]]; ok {
			return true
		}
		sMap[s[i]] = 1
	}

	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
