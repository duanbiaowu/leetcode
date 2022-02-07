package leetcode

func canPermutePalindrome(s string) bool {
	var counter [128]int
	for _, v := range s {
		if counter[v-' '] == 0 {
			counter[v-' ']++
		} else {
			counter[v-' ']--
		}
	}

	flag := true
	for _, v := range counter {
		if v > 0 {
			if flag {
				flag = false
			} else {
				return false
			}
		}
	}

	return true
}
